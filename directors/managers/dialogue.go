package managers

import (
	"github.com/wirdos/directors/input"
	"github.com/wirdos/ui"
	"github.com/wirdos/util"
)

// Manages the sequence of a dialogue event, directing the UI interface
type Dialogue struct {
	ui *ui.UI
	dialogueBox *ui.DialogueBox
}

func (d *Dialogue) FeedInput(input *input.Input) {
	if input.Interact {
		if d.dialogueBox != nil {
			d.dialogueBox = nil
			d.ui.DeleteDialogueBox()
		} else {
			// TODO: very messy interim...
			err := d.BeginScript()
			if err != nil {
				panic(err)
			}
		}
	}
}

// TODO: should take script
func (d *Dialogue) BeginScript() error {
	// TODO: the knowledge of the palette for the dialogue box definitely seems like something either
	// the dialogue box or the UI should know, not the dialogue manager
	palette, err := util.NewPalette("ui-elements")

	if err != nil {
		return err
	}

	dialogueBox := ui.NewDialogueBox(palette)
	dialogueBox.WriteText(
		"This is a very long piece of text which is printed on multiple lines by code and not designed with newlines as part of design... at least I hope so, it's designed so to split on word but there is a question of the text size which possibly overlaps no?",
	)

	d.dialogueBox = dialogueBox
	d.ui.AddDialogueBox(dialogueBox)

	return nil
}

func NewDialogue(ui *ui.UI) *Dialogue {
	return &Dialogue{
		ui: ui,
	}
}
