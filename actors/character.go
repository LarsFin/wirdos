package actors

import (
	"fmt"

	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources/atlases"
	"github.com/wirdos/util"
)

type Character struct {
	speed float64
	body  *util.Body
	face *util.Face
	velocity pixel.Vec
	facingDirection pixel.Vec
	stage *Stage
}

func (c *Character) Update() {
	if c.velocity.Len() > 0 {
		c.facingDirection = c.velocity.Unit()
		c.face.SetSpriteKey(util.Direction(c.velocity))
		c.body.Move(c.velocity, c.stage.Walls)
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

func (c *Character) FeedInteract(interact bool) {
	if interact {
		interactPoint := c.facingDirection.Scaled(16).Add(c.body.Position)
		for _, prop := range c.stage.Props {
			if prop.Interacting(interactPoint) {
				fmt.Print("Interacting with prop.\n")
			}
		}
	}
}

func (c *Character) Pos() pixel.Vec {
	return c.body.Position
}

func (c *Character) Face() *util.Face {
	return c.face
}

func (c *Character) PlaceOnStage(stage *Stage) {
	c.stage = stage
	c.body.Position = stage.SpawnPoint()
}

func NewCharacter(pos pixel.Vec, speed float64) (*Character, error) {
	// TODO: sprite map isn't different to texture map, should update so a palette is used here
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
