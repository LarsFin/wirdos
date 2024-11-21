package util

import (
	"math"
	"strings"
	"unicode"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/ext/text"
	"golang.org/x/image/font/basicfont"
)

// A helper for styling and formatting text given a defined container, for
// instance the ability to form newlines based on length of lines and words
type TextBox struct {
	writer *text.Text
	bounds pixel.Rect
	scale float64

	text string
	textAnimation *TextAnimation
}

type TextAnimation struct {
	// characters per second
	speed float64
	// number of characters written, including text sourced whitespace
	// newlines from box formatting are not included
	charactersWritten uint
	timeSinceLastWrite float64
}

type TextBoxOptions struct {
	Bounds pixel.Rect
	TextScale float64
	LineHeightScale float64
	TextAnimationOptions *TextAnimationOptions
}

type TextAnimationOptions struct {
	// characters per second
	Speed float64
}

func (tb *TextBox) Update() {
	// should no animation have been provided; return as there's nothing to update
	if tb.textAnimation == nil {
		return
	}

	// if the animation has been processed, return
	if tb.textAnimation.charactersWritten >= uint(len(tb.text)) {
		return
	}

	tb.textAnimation.timeSinceLastWrite += DeltaTime
	tick := 1. / tb.textAnimation.speed

	if tick > tb.textAnimation.timeSinceLastWrite {
		return
	}

	toWriteThisFrame := uint(math.Floor(tb.textAnimation.timeSinceLastWrite / tick))
	byFrameCharactersWritten := uint(math.Min(
		float64(tb.textAnimation.charactersWritten + toWriteThisFrame),
		float64(len(tb.text)),
	))
	tb.textAnimation.timeSinceLastWrite = math.Mod(tb.textAnimation.timeSinceLastWrite, tick)

	// TODO: what if a word is too long to fit on a line even on its own? This will be an infinite loop...
	for tb.textAnimation.charactersWritten < byFrameCharactersWritten {
		i := tb.textAnimation.charactersWritten

		// if this is the first character to a word, check word requires newline
		if i == 0 || !unicode.IsSpace(rune(tb.text[i])) && unicode.IsSpace(rune(tb.text[i -1])) {
			word := string(tb.text[i])
			for wli := i; wli < toWriteThisFrame && unicode.IsSpace(rune(tb.text[wli])); wli++ {
				word += string(tb.text[wli])
			}
			if tb.writer.BoundsOf(word).Max.X * tb.scale > tb.bounds.W() {
				tb.writer.WriteRune('\n')
			}
		}

		tb.writer.WriteByte(tb.text[i])
		tb.textAnimation.charactersWritten++
	}
}

func (tb *TextBox) writeAll() {
	// TODO: each word here is only split by a space, this should work against newline characters too...
	for i, word := range strings.Split(tb.text, " ") {
		if i > 0 {
			tb.writer.WriteRune(' ')
		}

		if tb.writer.BoundsOf(word).Max.X * tb.scale > tb.bounds.W() {
			tb.writer.WriteRune('\n')
		}

		tb.writer.WriteString(word)
	}
}

func (tb *TextBox) Draw(t pixel.Target, matrix pixel.Matrix) {
	tb.writer.Draw(t, matrix.Scaled(tb.writer.Orig, tb.scale))
}

func (tb *TextBox) SetText(s string) {
	tb.text = s
	tb.writer.Clear()

	if tb.textAnimation != nil {
		// TODO: probably should have a reset method on an animation for this
		tb.textAnimation.charactersWritten = 0
		tb.textAnimation.timeSinceLastWrite = 0
	} else {
		tb.writeAll()
	}
}

func NewTextBox(options TextBoxOptions) (*TextBox) {
	// TODO: this should be encapsulated in a separate util file, possibly
	// this should be bundled together in a sub package.
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	var textScale float64 = 1
	if options.TextScale > 0 {
		textScale = options.TextScale
	}

	lineHeight := atlas.LineHeight()
	if options.LineHeightScale > 0 {
		lineHeight *= options.LineHeightScale
	}

	bounds := options.Bounds

	writer := text.New(
		pixel.V(bounds.Min.X, bounds.Max.Y - lineHeight),
		atlas,
	)

	writer.Color = pixel.RGB(1, 1, 1)
	writer.LineHeight = lineHeight

	var textAnimation *TextAnimation
	if options.TextAnimationOptions != nil {
		textAnimation = &TextAnimation{
			speed: options.TextAnimationOptions.Speed,
		}
	}

	return &TextBox{
		writer: writer,
		bounds: bounds,
		scale: textScale,
		textAnimation: textAnimation,
	}
}