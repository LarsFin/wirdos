package managers

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/wirdos/actors"
)

type Camera struct {
	stage *actors.Stage
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
	limit := float64(8)

	if delta.Len() > limit {
		over := delta.Sub(delta.Unit().Scaled(limit))
		c.pos = c.pos.Add(over)
	}

	// TODO: if the view is outside stage boundary, it should be clamped within boundary instead
}

func (c *Camera) Render() {
	// we have to translate by the inverse to ensure components are rendered in their respected
	// position, this then helps to ensure the 'camera position' is correctly plained against
	// world coordinates
	mp := c.pos.Scaled(-c.zoom).Add(c.window.Bounds().Center())
	c.window.SetMatrix(pixel.IM.Moved(mp).Scaled(mp, c.zoom))

	c.window.Clear(pixel.RGB(1, 1, 1))

	// TODO: layering through sorting + y position
	for _, board := range c.stage.Boards {
		board.Draw(c.window)
	}

	c.stage.Character.Face().Draw(c.window)

	c.window.Update()
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

func NewCamera(window *opengl.Window, pos pixel.Vec, zoom float64, stage *actors.Stage) *Camera {
	return &Camera{
		stage: stage,
		pos: pos,
		zoom: zoom,
		window: window,
	}
}
