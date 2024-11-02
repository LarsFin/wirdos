package util

import (
	"github.com/gopxl/pixel/v2"
)

// different from a Face, a Board is a composite of static 'tiles' that make up
// layers which never change in ordering, this allows for optimised rendering with
// fewer draw calls using pixel's batch
type Board struct {
	layer int8
	batch *pixel.Batch
}

func (b *Board) Draw(t pixel.Target) {
	b.batch.Draw(t)
}

func (b *Board) Layer() int8 {
	return b.layer
}
