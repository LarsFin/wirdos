package events

type Event interface {
	Id() string
	Idempotent() bool
}

type DialogueEvent struct {
	id string
	scriptName string
}

func (de *DialogueEvent) Id() string {
	return de.id
}

func (de *DialogueEvent) Type() string {
	return "dialogue"
}

func (de *DialogueEvent) Idempotent() bool {
	return true
}

func (de *DialogueEvent) ScriptName() string {
	return de.scriptName
}

func NewDialogueEvent(id, scriptName string) *DialogueEvent {
	return &DialogueEvent{
		id: id,
		scriptName: scriptName,
	}
}
