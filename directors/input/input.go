package input

import (
	"github.com/gopxl/pixel/v2"
)

type InputController interface {
	Update()
	Input() *Input
}

type Input struct {
	Direction pixel.Vec
	Interact bool
	Exit bool
}
