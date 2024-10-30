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
	// TODO: work out how to do this...
	c.pos = pixel.IM.Project(c.stage.Character.Pos()).Sub(c.window.Bounds().Center())
}

func (c *Camera) Render() {
	c.window.SetMatrix(pixel.IM.Moved(c.pos).Scaled(c.pos, c.zoom))

	c.window.Clear(pixel.RGB(1, 1, 1))

	// TODO: layering through sorting + y position
	for _, board := range c.stage.Boards {
		board.Draw(c.window)
	}

	c.stage.Character.Face().Draw(c.window)

	c.window.Update()
}

func NewCamera(window *opengl.Window, pos pixel.Vec, zoom float64, stage *actors.Stage) *Camera {
	return &Camera{
		stage: stage,
		pos: pos,
		zoom: zoom,
		window: window,
	}
}
