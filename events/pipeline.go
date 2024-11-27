package events

import "slices"

// a pipeline is shared across managers and actors, events are pushed into
// a pipeline and then a resolving object must take the event off and handle
// it appropriately
type Pipeline struct {
	events []*Event
}

func (p *Pipeline) PushEvent(newEvent *Event) {
	// if an event for this type and resource name already exists, don't add it
	for _, e := range p.events {
		if e.Type == newEvent.Type && e.ResourceName == newEvent.ResourceName {
			return
		}
	}

	p.events = append(p.events, newEvent)
}

func (p *Pipeline) PullEventOfType(eventType EventType) *Event {
	eventIndex := -1

	for i, e := range p.events {
		if e.Type == eventType {
			eventIndex = i
			break
		}
	}

	if eventIndex < 0 {
		return nil
	}

	eventOfType := p.events[eventIndex]
	p.events = slices.Delete(p.events, eventIndex, 1)
	return eventOfType
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		events: make([]*Event, 0),
	}
}
