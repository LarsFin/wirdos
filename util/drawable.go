package util

import "github.com/gopxl/pixel/v2"

type Drawable interface {
	Draw(t pixel.Target)
	Layer() int8
}
