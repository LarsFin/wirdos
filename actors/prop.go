package actors

import (
	"fmt"

	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources"
)

// TODO: consider another name for 'actors' directory if it's going to contain more than just actors

type Prop struct {
	interactiveArea pixel.Rect
}

// TODO: a prop should return some data when interacted with so other party knows what behaviour to take
func (p *Prop) Interacting(point pixel.Vec) bool {
	fmt.Printf("Checking if %v is in %v\n", point, p.interactiveArea)
	return p.interactiveArea.Contains(point)
}

func NewProp(name string, pos pixel.Vec) (*Prop, error) {
	data, err := resources.LoadJSON[resources.PropData](fmt.Sprintf("props/%s", name))

	if err != nil {
		return nil, err
	}

	return &Prop{
		interactiveArea: pixel.R(
			pos.X-data.InteractiveDimensions.X/2,
			pos.Y-data.InteractiveDimensions.Y/2,
			pos.X+data.InteractiveDimensions.X/2,
			pos.Y+data.InteractiveDimensions.Y/2,
		),
	}, nil
}
