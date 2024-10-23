package util

import "github.com/gopxl/pixel/v2"

type Collider interface {
	CollidesRect(p pixel.Vec, r pixel.Rect) pixel.Rect
}

type RectCollider struct {
	// relative to the owning body's position
	r pixel.Rect
}

func (rc *RectCollider) CollidesRect(p pixel.Vec, r pixel.Rect) pixel.Rect {
	return pixel.R(p.X+rc.r.Min.X, p.Y+rc.r.Min.Y, p.X+rc.r.Max.X, p.Y+rc.r.Max.Y).Intersect(r)
}

func NewRectCollider(r pixel.Rect) *RectCollider {
	return &RectCollider{r: r}
}
