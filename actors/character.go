package actors

import (
	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources"
	"github.com/wirdos/util"
)

type Character struct {
	speed float64
	body  *util.Body
	face *util.Face
	velocity pixel.Vec

	// TODO: this should be fed through Stage/Map on Update instead
	walls []pixel.Rect
}

func (c *Character) Update() {
	if c.velocity.Len() > 0 {
		c.body.Move(c.velocity, c.walls)
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

func NewCharacter(pos pixel.Vec, speed float64, walls []pixel.Rect) (*Character, error) {
	pic, err := resources.LoadPNG("sprite")

	if err != nil {
		return nil, err
	}

	return &Character{
		speed: speed,
		body:  util.NewBody(pos, pixel.R(-4, -8, 4, 0)),
		face: util.NewFace(0, pixel.NewSprite(pic, pic.Bounds()), pos),
		velocity: pixel.ZV,
		// TODO: feed via map/stage on update
		walls: walls,
	}, nil
}
