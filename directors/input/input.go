package input

import (
	"github.com/gopxl/pixel/v2"
)

type Input interface {
	Update()

	Direction() pixel.Vec
	Interact() bool
	Exit() bool
}