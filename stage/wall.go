package stage

import "github.com/gopxl/pixel/v2"

// a solid rectangle that cannot be moved through by actors
type Wall struct {
	r pixel.Rect
}

func (w *Wall) CollisionArea() pixel.Rect {
	return w.r
}

func NewWall(r pixel.Rect) *Wall {
	return &Wall{r: r}
}
