package managers

import (
	"sort"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/wirdos/actors"
	"github.com/wirdos/ui"
	"github.com/wirdos/util"
)

type Camera struct {
	stage *actors.Stage
	ui *ui.UI
	pos pixel.Vec
	zoom float64
	window *opengl.Window
}

func (c *Camera) Update() {
	worldView := c.worldView()

	// if the boundary is not defined for a stage then it'll be a zero rect, for that reason
	// we should simply lock camera to center of stage (consider height and width separately)
	if worldView.Size().Len() > c.stage.Boundary.Size().Len() {
		c.pos = c.stage.Boundary.Center()
		return
	}

	delta := c.stage.Character.Pos().Sub(c.pos)
	limit := 8.

	if delta.Len() > limit {
		over := delta.Sub(delta.Unit().Scaled(limit))
		c.pos = c.pos.Add(over)
	}

	// need to reload world view as the camera position has changed
	worldView = c.worldView()
	clamp := util.ContainmentTranslation(worldView, c.stage.Boundary)
	c.pos = c.pos.Add(clamp)
}

func (c *Camera) Render() {
	c.window.Clear(pixel.RGB(1, 1, 1))

	// we have to translate by the inverse to ensure components are rendered in their respected
	// position, this then helps to ensure the 'camera position' is correctly plained against
	// world coordinates
	mp := c.pos.Scaled(-c.zoom).Add(c.window.Bounds().Center())
	c.window.SetMatrix(pixel.IM.Moved(mp).Scaled(mp, c.zoom))

	drawables := c.stage.GetDrawables()
	c.draw(drawables)

	// UI is rendered last onto the camera as it's always foreground, we also need to apply an IM
	// to the window as we want UI components to be rendered staticially
	c.window.SetMatrix(pixel.IM)

	drawables = c.ui.GetDrawables()
	c.draw(drawables)

	c.window.Update()
}

// handles ordering of drawable entities before making draw call to GPU
func (c *Camera) draw(drawables []util.Drawable) {
	sort.Slice(drawables, func(i, j int) bool {
		return drawables[i].Layer() > drawables[j].Layer()
	})

	for _, drawable := range drawables {
		drawable.Draw(c.window)
	}
}

func (c *Camera) worldView() pixel.Rect {
	realDimensions := c.window.Bounds().Center().Scaled(1/c.zoom)
	return pixel.R(
		c.pos.X-realDimensions.X,
		c.pos.Y-realDimensions.Y,
		c.pos.X+realDimensions.X,
		c.pos.Y+realDimensions.Y,
	)
}

func NewCamera(window *opengl.Window, pos pixel.Vec, zoom float64, stage *actors.Stage, ui *ui.UI) *Camera {
	return &Camera{
		stage: stage,
		ui: ui,
		pos: pos,
		zoom: zoom,
		window: window,
	}
}
