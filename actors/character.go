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

func (c * Character) Update(direction pixel.Vec, solids []*Solid) {
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
			lateralPush := float64(0)

			for _, solid := range solids {
				lateralSolidPush := solid.Collides(currentColliderFrame.Moved(pixel.V(lateralShift, 0))).W()
				
				if lateralSolidPush > lateralPush {
					lateralPush = lateralSolidPush
				}
			}

			newPos.X = (c.pos.X + lateralShift) - (lateralPush * direction.X)
		}

		if normal.Y != 0 {
			verticalShift := normal.Y * c.speed * util.DeltaTime
			verticalPush := float64(0)

			for _, solid := range solids {
				verticalSolidPush := solid.Collides(currentColliderFrame.Moved(pixel.V(0, verticalShift))).H()

				if verticalSolidPush > verticalPush {
					verticalPush = verticalSolidPush
				}
			}

			newPos.Y = (c.pos.Y + verticalShift) - verticalPush * direction.Y
		}

		// If a character is moving diagonally against a wall, resolve the direction so they move straight, thus
		// faster so the player doesn't have to ensure they're directing in a straight line
		if normal.X != 0 && normal.Y != 0 {
			if newPos.X != c.pos.X && newPos.Y == c.pos.Y {
				c.Update(pixel.V(direction.X, 0), solids)
				return
			}

			if newPos.Y != c.pos.Y && newPos.X == c.pos.X {
				c.Update(pixel.V(0, direction.Y), solids)
				return
			}
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
