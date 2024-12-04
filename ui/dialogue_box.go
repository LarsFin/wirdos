package ui

import (
	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/util"
)

// TODO: this should really be DialogueBox, then a separate definition could exist
// for DialogueBox around text behaviour and helpers (potentially existing solution?)
type DialogueBox struct {
	characterName *util.TextBox
	textBox *util.TextBox
	boxFace *util.Face
	isDestroyed bool
}

func (tb *DialogueBox) Update() {
	tb.textBox.Update()
}

func (tb *DialogueBox) WriteName(name string) {
	tb.characterName.SetText(name)
}

func (tb *DialogueBox) WriteText(text string) {
	tb.textBox.SetText(text)
}

func (tb *DialogueBox) Draw(t pixel.Target) {
	tb.boxFace.Draw(t)
	tb.textBox.Draw(t, pixel.IM)
	tb.characterName.Draw(t, pixel.IM)
}

func (tb *DialogueBox) Destroy() {
	tb.isDestroyed = true
}

func (tb *DialogueBox) IsDestroyed() bool {
	return tb.isDestroyed
}

func (tb *DialogueBox) SkipTextAnimation() {
	tb.textBox.SkipAnimation()
}

func (tb *DialogueBox) CurrentlyAnimating() bool {
	return tb.textBox.InAnimation()
}

func characterNameAreaBox(dialogueArea pixel.Rect) pixel.Rect {
	return pixel.R(
		dialogueArea.Min.X + 32,
		dialogueArea.Max.Y - 64,
		dialogueArea.Max.X + 32,
		dialogueArea.Max.Y - 32,
	)
}

func textAreaBox(dialogueArea pixel.Rect) pixel.Rect {
	return pixel.R(
		dialogueArea.Min.X + 32,
		dialogueArea.Min.Y + 32,
		dialogueArea.Max.X + 32,
		dialogueArea.Max.Y - 72,
	)
}

func NewDialogueBox(theme *Theme) *DialogueBox {
	face := util.NewFace(0, theme.Palette, "only", theme.Palette.Pic.Bounds().Center())

	characterNameBox := util.NewTextBox(util.TextBoxOptions{
		Bounds: characterNameAreaBox(theme.Palette.Pic.Bounds()),
		TextAtlas: theme.TextAtlas,
		TextScale: 3.,
	})

	textBox := util.NewTextBox(util.TextBoxOptions{
		Bounds: textAreaBox(theme.Palette.Pic.Bounds()),
		TextAtlas: theme.TextAtlas,
		TextScale: 2.,				
		LineHeightScale: 1.5,
		TextAnimationOptions: &util.TextAnimationOptions{
			Speed: 60,
		},
	})

	return &DialogueBox{
		characterName: characterNameBox,
		textBox: textBox,
		boxFace: face,
		isDestroyed: false,
	}
}
