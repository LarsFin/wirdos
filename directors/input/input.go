package input

import (
	"github.com/gopxl/pixel/v2"
)

type Input interface {
	Update()

	Direction() pixel.Vec
	Exit() bool
}