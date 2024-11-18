package ui

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/ext/text"
	"github.com/wirdos/util"
	"golang.org/x/image/font/basicfont"
)

// TODO: this should really be DialogueBox, then a separate definition could exist
// for DialogueBox around text behaviour and helpers (potentially existing solution?)
type DialogueBox struct {
	textBox *util.TextBox
	boxFace *util.Face
}

func (tb *DialogueBox) WriteText(text string) {
	tb.textBox.SetText(text)
}

func (tb *DialogueBox) Draw(t pixel.Target) {
	tb.boxFace.Draw(t)
	tb.textBox.Draw(t, pixel.IM)
}

func textAreaBox(dialogueArea pixel.Rect) pixel.Rect {
	return pixel.R(
		dialogueArea.Min.X + 32,
		dialogueArea.Min.Y + 32,
		dialogueArea.Max.X + 32,
		dialogueArea.Max.Y - 64,
	)
}

func NewDialogueBox(palette *util.Palette) *DialogueBox {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	text := text.New(pixel.V(64, 172), atlas)
	text.Color = pixel.RGB(1, 1, 1)

	face := util.NewFace(0, palette, "only", palette.Pic.Bounds().Center())
	textBox := util.NewTextBox(textAreaBox(palette.Pic.Bounds()), 2., 1.5)

	return &DialogueBox{
		textBox: textBox,
		boxFace: face,
	}
}
