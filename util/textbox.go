package util

import (
	"fmt"
	"math"
	"unicode"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/ext/text"
	"github.com/wirdos/logger"
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
	skip bool
	cancelled bool
}

type TextBoxOptions struct {
	Bounds pixel.Rect
	TextAtlas *text.Atlas
	TextScale float64
	LineHeightScale float64
	TextAnimationOptions *TextAnimationOptions
}

type TextAnimationOptions struct {
	// characters per second
	Speed float64
}

func (tb *TextBox) Update() {
	if !tb.InAnimation() {
		return
	}

	tb.textAnimation.timeSinceLastWrite += DeltaTime
	tick := 1. / tb.textAnimation.speed

	if tick > tb.textAnimation.timeSinceLastWrite {
		return
	}

	var byFrameCharactersWritten uint

	if tb.textAnimation.skip {
		byFrameCharactersWritten = uint(len(tb.text))
		tb.textAnimation.skip = false
	} else {
		toWriteThisFrame := uint(math.Floor(tb.textAnimation.timeSinceLastWrite / tick))
		byFrameCharactersWritten = uint(math.Min(
			float64(tb.textAnimation.charactersWritten + toWriteThisFrame),
			float64(len(tb.text)),
		))
	}

	tb.textAnimation.timeSinceLastWrite = math.Mod(tb.textAnimation.timeSinceLastWrite, tick)

	err := tb.writeText(tb.textAnimation.charactersWritten, byFrameCharactersWritten)

	if err != nil {
		logger.Warn(fmt.Sprintf("failed to write text: (%s) to container: %v", tb.text, tb.bounds))
		tb.textAnimation.cancelled = true
		return
	}

	tb.textAnimation.charactersWritten = byFrameCharactersWritten
}

func (tb *TextBox) outOfBounds(s string) bool {
	return tb.writer.BoundsOf(s).Max.X * tb.scale > tb.bounds.W()
}

// determines whether word from index is potentially written is beyond container
func (tb *TextBox) peekWord(i uint) bool {
	word := ""
	for uint(len(tb.text)) > i && !unicode.IsSpace(rune(tb.text[i])) {
		word += string(tb.text[i])
		i++
	}
	return tb.outOfBounds(word)
}

// write text between 'from' and 'to', including 'from' but excluding 'to'.
func (tb *TextBox) writeText(from, to uint) error {
	for i := from; i < to; i++ {
		isSpace := unicode.IsSpace(rune(tb.text[i]))
		newLine := i == 0 || tb.text[i - 1] == '\n'
		newWord := !isSpace && newLine || unicode.IsSpace(rune(tb.text[i - 1]))
		
		if isSpace && tb.outOfBounds(string(tb.text[i])) {
			// if the character is whitespace that would put the text out of bounds on the horizontal
			// of the textbox; skip character as it's not valuable to write
			continue
		} else if newWord && tb.peekWord(i) {
			// if it's a new line and the word doesn't fit the container in width, this means the text
			// can never be rendered *currently*, potentially it could return a scalar reduction required
			// to fit the text in width of the container.
			if newLine {
				return fmt.Errorf("failed to write text: (%s) to container: %v", tb.text, tb.bounds)
			}

			tb.writer.WriteRune('\n')
		}

		tb.writer.WriteByte(tb.text[i])
	}

	return nil
}

func (tb *TextBox) InAnimation() bool {
	return tb.textAnimation != nil &&
		!tb.textAnimation.cancelled &&
		tb.textAnimation.charactersWritten < uint(len(tb.text))
}

func (tb *TextBox) SkipAnimation() {
	tb.textAnimation.skip = true
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
		err := tb.writeText(0, uint(len(s)))
		if err != nil {
			logger.Warn(fmt.Sprintf("failed to write text: (%s) to container: %v", tb.text, tb.bounds))
		}
	}
}

func NewTextBox(options TextBoxOptions) (*TextBox) {
	atlas := options.TextAtlas

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
			cancelled: false,
		}
	}

	return &TextBox{
		writer: writer,
		bounds: bounds,
		scale: textScale,
		textAnimation: textAnimation,
	}
}
