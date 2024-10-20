package actors

import (
	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/util"
)

type Character struct {
	pos pixel.Vec
	speed float64
}

func (c *Character) Update(direction pixel.Vec) {
	if direction.Len() > 0 {
		mv := direction.Unit().Scaled(c.speed * util.DeltaTime)
		c.pos = c.pos.Add(mv)
	}
}

func (c *Character) Pos() pixel.Vec {
	return c.pos
}

func NewCharacter(pos pixel.Vec, speed float64) *Character {
	return &Character{
		pos: pos,
		speed: speed,
	}
}
