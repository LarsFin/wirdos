package util

import (
	"strings"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/ext/text"
	"golang.org/x/image/font/basicfont"
)

// A helper for styling and formatting text given a defined container, for
// instance the ability to form newlines based on length of lines and words
type TextBox struct {
	writer *text.Text
	bounds pixel.Rect
}

func (tb *TextBox) Draw(t pixel.Target, matrix pixel.Matrix) {
	tb.writer.Draw(t, matrix)
}

func (tb *TextBox) SetText(s string) {
	words := strings.SplitAfter(s, " ")
	for _, word := range words {
		if tb.writer.BoundsOf(word).Max.X > tb.bounds.W() {
			tb.writer.WriteRune('\n')
		}

		tb.writer.WriteString(word)
	}
}

func NewTextBox(box pixel.Rect) (*TextBox) {
	// TODO: this should be encapsulated in a separate util file, possibly
	// this should be bundled together in a sub package.
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	writer := text.New(
		pixel.V(box.Min.X, box.Max.Y - atlas.LineHeight()),
		atlas,
	)

	writer.Color = pixel.RGB(1, 1, 1)

	return &TextBox{
		writer: writer,
		bounds: box,
	}
}
