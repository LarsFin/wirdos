package managers

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/wirdos/util"
)

type Camera struct {
	// TODO: faces should come through a stage/map
	faces []*util.Face
	pos pixel.Vec
	zoom float64
	window *opengl.Window
}

func (c *Camera) Render() {
	c.window.SetMatrix(pixel.IM.Scaled(c.pos, c.zoom))

	c.window.Clear(pixel.RGB(1, 1, 1))

	// TODO: determine ordering of faces to draw (e.g; layering and 'distance')
	for _, face := range c.faces {
		face.Draw(c.window)
	}

	c.window.Update()
}

func NewCamera(window *opengl.Window, pos pixel.Vec, zoom float64, faces []*util.Face) *Camera {
	return &Camera{
		faces: faces,
		pos: pos,
		zoom: zoom,
		window: window,
	}
}
