package actors

import (
	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources/atlases"
	"github.com/wirdos/structure"
	"github.com/wirdos/util"
)

type Character struct {
	speed float64
	body  *util.Body
	face *util.Face
	velocity pixel.Vec
}

func (c *Character) Update(stage *structure.Stage) {
	if c.velocity.Len() > 0 {
		c.face.SetSpriteKey(util.Direction(c.velocity))
		c.body.Move(c.velocity, stage.Walls)
	}

	c.face.Update(c.body.Position)
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

func (c *Character) Face() *util.Face {
	return c.face
}

func NewCharacter(pos pixel.Vec, speed float64) (*Character, error) {
	spriteMap, err := atlases.GenerateSpriteMap("character")

	if err != nil {
		return nil, err
	}

	return &Character{
		speed: speed,
		body:  util.NewBody(pos, pixel.R(-4, -8, 4, 0)),
		face: util.NewFace(0, spriteMap, "right", pos),
		velocity: pixel.ZV,
	}, nil
}
