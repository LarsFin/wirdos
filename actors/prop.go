package actors

import (
	"fmt"

	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/events"
	"github.com/wirdos/resources"
	"github.com/wirdos/util"
)

// TODO: consider another name for 'actors' directory if it's going to contain more than just actors

type Prop struct {
	face *util.Face
	interactiveArea pixel.Rect
	interactionEvent *events.Event
}

func (p *Prop) Face() *util.Face {
	return p.face
}

func (p *Prop) Interaction(point pixel.Vec) *events.Event {
	if p.interactiveArea.Contains(point) {
		return p.interactionEvent
	}

	return nil
}

func NewProp(name string, pos pixel.Vec) (*Prop, error) {
	data, err := resources.LoadJSON[resources.PropData](fmt.Sprintf("props/%s", name))

	if err != nil {
		return nil, err
	}

	palette, err := util.NewPalette(data.Palette.Name)

	if err != nil {
		return nil, err
	}

	face := util.NewFace(data.Layer, palette, data.Palette.InitialKey, pos)

	return &Prop{
		face: face,
		interactiveArea: pixel.R(
			pos.X-data.InteractiveDimensions.X/2,
			pos.Y-data.InteractiveDimensions.Y/2,
			pos.X+data.InteractiveDimensions.X/2,
			pos.Y+data.InteractiveDimensions.Y/2,
		),
		interactionEvent: events.FromData(data.InteractionEvent),
	}, nil
}
