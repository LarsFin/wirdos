package ui

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/text"
	"github.com/wirdos/util"
	"golang.org/x/image/font/basicfont"
)

// UI isn't rendered through the camera, this is to avoid it being transformed
// based on the camera's matrix, it should be simply directly rendered to the
// view based on fixed coordinates
type UI struct {
	dialogueBox *DialogueBox
	window *opengl.Window

	theme *Theme
}

func (ui *UI) Update() {
	if ui.dialogueBox != nil {
		ui.dialogueBox.Update()
	}
}

func (ui *UI) Render() {
	ui.window.SetMatrix(pixel.IM)

	// TODO: look into whether it's possible to batch here, it would be easy if the
	// components only used sprites from the ui-elements files but they're also going
	// to commonly write text. The problem there is that the text can't be written to
	// the batch as it's not the same picture, if it's not part of the same batch it's
	// likely there could be layering issues
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

// TODO: theme is probably the wrong word right now but I like the idea of choosing
// different ui sprites based on a theme selected by the player
func (ui *UI) Theme() *Theme {
	return ui.theme
}

func NewUI(window *opengl.Window) (*UI, error) {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	palette, err := util.NewPalette("ui-elements")

	if err != nil {
		return nil, err
	}

	return &UI{
		window: window,
		theme: &Theme{
			Palette: palette,
			TextAtlas: atlas,
		},
	}, nil
}
