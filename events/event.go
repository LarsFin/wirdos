package events

import (
	"fmt"

	"github.com/wirdos/resources"
)

type EventType int

const (
	StartDialogue EventType = iota
	EndDialogue EventType = iota
)

type Event struct {
	Type EventType
	ResourceName string
}

func FromData(eventData resources.Event) *Event {
	switch eventData.Type {
	case "start_dialogue":
		return NewEvent(StartDialogue, eventData.ResourceName)
	default:
		// TODO: consider better logger down the road
		fmt.Printf("warning: could not map event from data event-type '%s'", eventData.Type)
		return nil
	}
}

func NewEvent(eventType EventType, resourceName string) *Event {
	return &Event{
		Type: eventType,
		ResourceName: resourceName,
	}
}
