package util

import "github.com/gopxl/pixel/v2"

type Drawable interface {
	Draw(pixel.Target)
	Layer() int8
}
