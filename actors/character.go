package actors

import (
	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/util"
)

type Character struct {
	pos pixel.Vec
	speed float64
	collisionDimensions pixel.Vec
}

func (c *Character) Update(direction pixel.Vec, solid *Solid) {
	if direction.Len() > 0 {
		normal := direction.Unit()
		currentColliderFrame := pixel.R(
			c.pos.X-c.collisionDimensions.X/2,
			c.pos.Y-c.collisionDimensions.Y/2,
			c.pos.X+c.collisionDimensions.X/2,
			c.pos.Y+c.collisionDimensions.Y/2,
		)

		newPos := pixel.V(c.pos.X, c.pos.Y)

		if normal.X != 0 {
			lateralShift := normal.X * c.speed * util.DeltaTime
			lateralCollision := solid.Collides(currentColliderFrame.Moved(pixel.V(lateralShift, 0)))
			newPos.X = (c.pos.X + lateralShift) - (lateralCollision.Max.X - lateralCollision.Min.X) * direction.X
		}

		if normal.Y != 0 {
			verticalShift := normal.Y * c.speed * util.DeltaTime
			verticalCollision := solid.Collides(currentColliderFrame.Moved(pixel.V(0, verticalShift)))
			newPos.Y = (c.pos.Y + verticalShift) - (verticalCollision.Max.Y - verticalCollision.Min.Y) * direction.Y
		}

		c.pos = newPos
	}
}

func (c *Character) Pos() pixel.Vec {
	return c.pos
}

func NewCharacter(pos pixel.Vec, speed float64, collision_dimensions pixel.Vec) *Character {
	return &Character{
		pos: pos,
		speed: speed,
		collisionDimensions: collision_dimensions,
	}
}
