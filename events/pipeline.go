package events

import "slices"

// a pipeline is shared across managers and actors, events are pushed into
// a pipeline and then a resolving object must take the event off and handle
// it appropriately
type Pipeline struct {
	events []Event
}

func (p *Pipeline) PushEvent(newEvent Event) {
	// if the event is idempotent and events already has an unhandled event
	// with the given id; don't add the event
	if newEvent.Idempotent() {
		for _, e := range p.events {
			if e.Id() == newEvent.Id() {
				return
			}
		}
	}

	p.events = append(p.events, newEvent)
}

func PopEventOfType[E Event](p *Pipeline) *E {
	eventIndex := -1

	for i, e := range p.events {
		if _, ok := e.(E); ok {
			eventIndex = i
			break
		}
	}

	if eventIndex < 0 {
		return nil
	}

	eventOfType := p.events[eventIndex].(E)
	p.events = slices.Delete(p.events, eventIndex, 1)
	return &eventOfType
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		events: make([]Event, 0),
	}
}
