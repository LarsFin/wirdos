package util

import (
	"github.com/gopxl/pixel/v2"
)

type Body struct {
	Position pixel.Vec
	collider Collider
}

func (b *Body) Move(velocity pixel.Vec, walls []pixel.Rect) {
	newPos := pixel.V(b.Position.X, b.Position.Y)

	if velocity.X != 0 {
		lateralShift := velocity.X * DeltaTime
		lateralPush := float64(0)

		for _, wall := range walls {
			lateralSolidPush := b.collider.CollidesRect(b.Position.Add(pixel.V(lateralShift, 0)), wall).W()

			if lateralSolidPush > lateralPush {
				lateralPush = lateralSolidPush
			}
		}

		if velocity.X < 0 {
			lateralPush = -lateralPush
		}

		newPos.X = (b.Position.X + lateralShift) - lateralPush
	}

	if velocity.Y != 0 {
		verticalShift := velocity.Y * DeltaTime
		verticalPush := float64(0)

		for _, wall := range walls {
			verticalSolidPush := b.collider.CollidesRect(b.Position.Add(pixel.V(0, verticalShift)), wall).H()

			if verticalSolidPush > verticalPush {
				verticalPush = verticalSolidPush
			}
		}

		if velocity.Y < 0 {
			verticalPush = -verticalPush
		}

		newPos.Y = (b.Position.Y + verticalShift) - verticalPush
	}

	// if the body is moving diagonally into a wall, to help optimise the movement we can add any
	// pushback towards the allowed motion
	if velocity.X != 0 && velocity.Y != 0 {
		if newPos.X != b.Position.X && newPos.Y == b.Position.Y {
			resolvedMotion := velocity.Len()
			if (velocity.X < 0) { resolvedMotion = -resolvedMotion }
			b.Move(pixel.V(resolvedMotion, 0), walls)
			return
		}
		
		if newPos.Y != b.Position.Y && newPos.X == b.Position.X {
			resolvedMotion := velocity.Len()
			if (velocity.Y < 0) { resolvedMotion = -resolvedMotion }
			b.Move(pixel.V(0, resolvedMotion), walls)
			return
		}
	}

	b.Position = newPos
}

func NewBody(pos pixel.Vec, relativeCollisionFrame pixel.Rect) *Body {
	return &Body{
		Position: pos,
		collider: NewRectCollider(relativeCollisionFrame),
	}
}
