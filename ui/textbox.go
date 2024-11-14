package ui

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/ext/text"
	"github.com/wirdos/util"
	"golang.org/x/image/font/basicfont"
)

// TODO: this should really be DialogueBox, then a separate definition could exist
// for TextBox around text behaviour and helpers (potentially existing solution?)
type TextBox struct {
	text *text.Text
	face *util.Face
}

func (tb *TextBox) Clear() {
	tb.text.Clear()
}

func (tb *TextBox) WriteText(text string) {
	tb.text.WriteString(text)
}

func (tb *TextBox) Draw(t pixel.Target) {
	tb.face.Draw(t)
	// TODO: probably don't want to scale this, ideally use font size I think?
	// otherwise we'll have to determine textbox bounds based on scalar... also
	// this is quite inefficient based on the text quantity
	tb.text.Draw(t, pixel.IM.Scaled(tb.text.Orig, 3))
}

func NewTextBox(palette *util.Palette) *TextBox {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	text := text.New(pixel.V(64, 172), atlas)
	text.Color = pixel.RGB(1, 1, 1)

	face := util.NewFace(0, palette, "only", palette.Pic.Bounds().Center())

	return &TextBox{
		text: text,
		face: face,
	}
}
