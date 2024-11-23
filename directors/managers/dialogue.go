package managers

import (
	"github.com/wirdos/directors/input"
	"github.com/wirdos/ui"
)

// Manages the sequence of a dialogue event, directing the UI interface
type Dialogue struct {
	ui *ui.UI
	dialogueBox *ui.DialogueBox
}

func (d *Dialogue) FeedInput(input *input.Input) {
	if input.Interact {
		if d.dialogueBox != nil {
			if d.dialogueBox.CurrentlyAnimating() {
				d.dialogueBox.SkipTextAnimation()
			} else {
				d.dialogueBox.Destroy()
				d.dialogueBox = nil
			}
		} else {
			d.BeginScript()
		}
	}
}

// TODO: should take script
func (d *Dialogue) BeginScript() {
	dialogueBox := ui.NewDialogueBox(d.ui.Theme())
	dialogueBox.WriteText(
		"This is a very long piece of text which is printed on multiple lines by code and not designed with newlines as part of design... at least I hope so, it's designed so to split on word but there is a question of the text size which possibly overlaps no?",
	)

	d.dialogueBox = dialogueBox
	d.ui.AddComponent(dialogueBox)
}

func NewDialogue(ui *ui.UI) *Dialogue {
	return &Dialogue{
		ui: ui,
	}
}
