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
	textbox *TextBox
	window *opengl.Window
}

func (ui *UI) Render() {
	ui.window.SetMatrix(pixel.IM)
	ui.textbox.Draw(ui.window)
}

func NewUI(window *opengl.Window) (*UI, error) {
	palette, err := util.NewPalette("ui-elements")

	if err != nil {
		return nil, err
	}

	textBox := NewTextBox(palette)
	textBox.WriteText("This is static hardcoded text\nEventually to be dynamically generated and organised")

	return &UI{
		textbox: textBox,
		window: window,
	}, nil
}
