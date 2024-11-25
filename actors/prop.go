package actors

import (
	"fmt"

	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources"
	"github.com/wirdos/util"
)

// TODO: consider another name for 'actors' directory if it's going to contain more than just actors

type Prop struct {
	face *util.Face
	interactiveArea pixel.Rect
}

func (p *Prop) Face() *util.Face {
	return p.face
}

// TODO: a prop should return some data when interacted with so other party knows what behaviour to take
func (p *Prop) Interacting(point pixel.Vec) bool {
	return p.interactiveArea.Contains(point)
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
	}, nil
}
