package ui

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

// UI isn't rendered through the camera, this is to avoid it being transformed
// based on the camera's matrix, it should be simply directly rendered to the
// view based on fixed coordinates
type UI struct {
	dialogueBox *DialogueBox
	window *opengl.Window
}

func (ui *UI) Update() {
	if ui.dialogueBox != nil {
		ui.dialogueBox.Update()
	}
}

func (ui *UI) Render() {
	ui.window.SetMatrix(pixel.IM)

	if ui.dialogueBox != nil {
		ui.dialogueBox.Draw(ui.window)
	}
}

// TODO: obviously should be more abstract than this, should have a map of drawable
// components which are drawable
func (ui *UI) AddDialogueBox(db *DialogueBox) {
	ui.dialogueBox = db
}

// TODO: see above TODO
func (ui *UI) DeleteDialogueBox() {
	ui.dialogueBox = nil
}

func NewUI(window *opengl.Window) (*UI, error) {
	return &UI{
		window: window,
	}, nil
}
