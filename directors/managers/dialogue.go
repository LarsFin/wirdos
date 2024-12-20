package managers

import (
	"fmt"

	"github.com/wirdos/directors/input"
	"github.com/wirdos/events"
	"github.com/wirdos/resources"
	"github.com/wirdos/ui"
)

// Manages the sequence of a dialogue event, directing the UI interface
type Dialogue struct {
	ui *ui.UI
	dialogueBox *ui.DialogueBox
	eventPipeline *events.Pipeline

	currentScript *resources.ScriptData
	lineNum int
}

func (d *Dialogue) FeedInput(input *input.Input) {
	if input.Interact && d.currentScript != nil {
		if d.dialogueBox.CurrentlyAnimating() {
			d.dialogueBox.SkipTextAnimation()
		} else {
			d.NextLine()
		}
	}
}

func (d *Dialogue) BeginScript(scriptName string) error {
	d.lineNum = 0
	script, err := resources.LoadJSON[resources.ScriptData](
		fmt.Sprintf("scripts/%s/%s", resources.GameOptions.GetLanguage(), scriptName),
	)

	if err != nil {
		return err
	}

	d.currentScript = script

	dialogueBox := ui.NewDialogueBox(d.ui.Theme())
	d.dialogueBox = dialogueBox
	d.ui.AddComponent(dialogueBox)

	d.NextLine()

	return nil
}

func (d *Dialogue) NextLine() {
	if d.currentScript == nil {
		return
	}

	if d.lineNum >= len(d.currentScript.Lines) {
		d.scriptFinished()
		return
	}

	line := d.currentScript.Lines[d.lineNum]
	d.dialogueBox.WriteLine(line)
	d.lineNum++
}

func (d *Dialogue) scriptFinished() {
	d.dialogueBox.Destroy()
	d.dialogueBox = nil
	d.currentScript = nil
	d.lineNum = 0

	d.eventPipeline.PushEvent(events.NewSimpleEvent(events.EndDialogue))
}

func NewDialogue(ui *ui.UI, eventPipeline *events.Pipeline) *Dialogue {
	return &Dialogue{
		ui: ui,
		eventPipeline: eventPipeline,
	}
}
