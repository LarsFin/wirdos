package stage

import "github.com/gopxl/pixel/v2"

// a solid rectangle that cannot be moved through by actors
type Wall struct {
	r pixel.Rect
}

func (w *Wall) Collides(other pixel.Rect) pixel.Rect {
	return w.r.Intersect(other)
}

func NewWall(r pixel.Rect) *Wall {
	return &Wall{r: r}
}
