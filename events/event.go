package events

type EventType int

const (
	StartDialogue EventType = iota
	EndDialogue EventType = iota
)

type Event struct {
	Type EventType
	ResourceName string
}

func NewEvent(eventType EventType, resourceName string) *Event {
	return &Event{
		Type: eventType,
		ResourceName: resourceName,
	}
}
