package actors

import "github.com/gopxl/pixel/v2"

type Solid struct {
	collision pixel.Rect
}

func (s *Solid) Collides(other pixel.Rect) pixel.Rect {
	return s.collision.Intersect(other)
}

func NewSolid(pos pixel.Vec, dimensions pixel.Vec) *Solid {
	return &Solid{
		collision: pixel.R(
			pos.X-dimensions.X/2,
			pos.Y-dimensions.Y/2,
			pos.X+dimensions.X/2,
			pos.Y+dimensions.Y/2,
		),
	}
}
