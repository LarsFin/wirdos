package managers

import (
	"sort"

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

	sort.Slice(c.faces, func(i, j int) bool {
		// TODO: sort by y position too
		return c.faces[i].Layer > c.faces[j].Layer
	})

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
