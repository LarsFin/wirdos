package util

import (
	"github.com/gopxl/pixel/v2"
)

// Face is the render information of a component within a scene

type Face struct {
	Layer int8

	spriteMap map[string]*pixel.Sprite
	activeSpriteKey string

	// TODO: obviously wrong to hold the position in two places if the body is the most relevant
	// but let's explore for now
	pos pixel.Vec
}

func (f *Face) Update(pos pixel.Vec) {
	f.pos = pos
}

func (f *Face) Position() pixel.Vec {
	return f.pos
}

func (f *Face) Draw(t pixel.Target) {
	f.spriteMap[f.activeSpriteKey].Draw(t, pixel.IM.Moved(f.pos))
}

func (f *Face) SetSpriteKey(key string) {
	f.activeSpriteKey = key
}

func NewFace(layer int8, spriteMap map[string]*pixel.Sprite, spriteKey string, pos pixel.Vec) *Face {
	return &Face{
		Layer: layer,
		spriteMap: spriteMap,
		activeSpriteKey: spriteKey,
		pos: pos,
	}
}
