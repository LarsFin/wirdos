package input

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

type KeyboardMouse struct {
	window *opengl.Window
	input *Input
}

func (kbm *KeyboardMouse) Update() {
	kbm.input = &Input{
		Interact: kbm.interact(),
		Direction: kbm.direction(),
		Exit: kbm.exit(),
	}
}

func (kbm *KeyboardMouse) Input() *Input {
	return kbm.input
}

func (kbm *KeyboardMouse) direction() pixel.Vec {
	direction := pixel.ZV

	if kbm.window.Pressed(pixel.KeyA) {
		direction.X--
	}

	if kbm.window.Pressed(pixel.KeyD) {
		direction.X++
	}

	if kbm.window.Pressed(pixel.KeyW) {
		direction.Y++
	}

	if kbm.window.Pressed(pixel.KeyS) {
		direction.Y--
	}

	return direction
}

// TODO: buffer?
func (kbm *KeyboardMouse) interact() bool {
	return kbm.window.JustPressed(pixel.KeyM)
}

func (kbm *KeyboardMouse) exit() bool {
	return kbm.window.JustPressed(pixel.KeyEscape)
}

func NewKeyboardMouse(window *opengl.Window) *KeyboardMouse {
	return &KeyboardMouse{window: window}
}

