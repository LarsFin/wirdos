package actors

import (
	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/stage"
	"github.com/wirdos/util"
)

type Character struct {
	speed float64
	body  *util.Body
	velocity pixel.Vec

	// TODO: this should be fed through Map on Update instead
	walls []*stage.Wall
}

func (c *Character) Update() {
	if c.velocity.Len() > 0 {
		c.body.Move(c.velocity, c.walls)
	}
}

func (c *Character) FeedDirection(direction pixel.Vec) {
	if direction.Len() > 0 {
		c.velocity = direction.Unit().Scaled(c.speed)
	} else {
		c.velocity = pixel.ZV
	}
}

func (c *Character) Pos() pixel.Vec {
	return c.body.Position
}

func NewCharacter(pos pixel.Vec, speed float64, walls []*stage.Wall) *Character {
	return &Character{
		speed: speed,
		body:  util.NewBody(pos, pixel.R(-4, -8, 4, 0)),
		velocity: pixel.ZV,
		// TODO: feed via map on update
		walls: walls,
	}
}
