package input

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

type KeyboardMouse struct {
	window *opengl.Window
	
	direction pixel.Vec
	exit bool
}

func (kbm *KeyboardMouse) Update() {
	kbm.direction = pixel.ZV

	if kbm.window.Pressed(pixel.KeyA) {
		kbm.direction.X--
	}

	if kbm.window.Pressed(pixel.KeyD) {
		kbm.direction.X++
	}

	if kbm.window.Pressed(pixel.KeyW) {
		kbm.direction.Y++
	}

	if kbm.window.Pressed(pixel.KeyS) {
		kbm.direction.Y--
	}

	kbm.exit = kbm.window.JustPressed(pixel.KeyEscape)
}

func (kbm *KeyboardMouse) Direction() pixel.Vec {
	return kbm.direction
}

func (kbm *KeyboardMouse) Exit() bool {
	return kbm.exit
}

func NewKeyboardMouse(window *opengl.Window) *KeyboardMouse {
	return &KeyboardMouse{window: window}
}
