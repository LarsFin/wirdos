package actors

import (
	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/stage"
	"github.com/wirdos/util"
)

type Character struct {
	speed float64
	body  *util.Body
}

func (c *Character) Update(direction pixel.Vec, walls []*stage.Wall) {
	if direction.Len() > 0 {
		c.body.Move(direction.Unit().Scaled(c.speed), walls)
	}
}

func (c *Character) Pos() pixel.Vec {
	return c.body.Position
}

func NewCharacter(pos pixel.Vec, speed float64, collision_dimensions pixel.Vec) *Character {
	return &Character{
		speed: speed,
		body:  util.NewBody(pos, pixel.R(-4, -8, 4, 0)),
	}
}
