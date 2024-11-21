package ui

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/wirdos/util"
)

// UI isn't rendered through the camera, this is to avoid it being transformed
// based on the camera's matrix, it should be simply directly rendered to the
// view based on fixed coordinates
type UI struct {
	dialogueBox *DialogueBox
	window *opengl.Window
}

func (ui *UI) Update() {
	ui.dialogueBox.Update()
}

func (ui *UI) Render() {
	ui.window.SetMatrix(pixel.IM)
	ui.dialogueBox.Draw(ui.window)
}

func NewUI(window *opengl.Window) (*UI, error) {
	palette, err := util.NewPalette("ui-elements")

	if err != nil {
		return nil, err
	}

	dialogueBox := NewDialogueBox(palette)
	dialogueBox.WriteText(
		"This is a very long piece of text which is printed on multiple lines by code and not designed with newlines as part of design... at least I hope so, it's designed so to split on word but there is a question of the text size which possibly overlaps no?",
	)

	return &UI{
		dialogueBox: dialogueBox,
		window: window,
	}, nil
}
