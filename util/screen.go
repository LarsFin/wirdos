package util

import (
	"github.com/gopxl/pixel/v2"
)

// different from a Face, a Screen is a composite of static 'tiles' that make up
// layers which never change in ordering, this allows for optimised rendering with
// fewer draw calls
type Screen struct {
	Layer int8

	batch *pixel.Batch
}

func (s *Screen) AddSprite(sprite *pixel.Sprite, pos pixel.Vec) {
	sprite.Draw(s.batch, pixel.IM.Moved(pos))
}

func (s *Screen) Draw(t pixel.Target) {
	s.batch.Draw(t)
}

func NewScreen(layer int8) *Screen {
	return &Screen{
		Layer: layer,
		batch: pixel.NewBatch(&pixel.TrianglesData{}, nil),
	}
}
