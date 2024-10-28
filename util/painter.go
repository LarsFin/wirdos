package util

import (
	"fmt"

	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources"
)

type Painter struct {
	pic pixel.Picture
	palette map[string]*pixel.Sprite
}

func (p *Painter) PaintBoard(layer int8, tiles []resources.TileData) *Board {
	batch := pixel.NewBatch(&pixel.TrianglesData{}, p.pic)
	
	// TODO: what happens if the key doesn't exist?
	for _, tile := range tiles {
		sprite := p.palette[tile.Key]
		sprite.Draw(batch, pixel.IM.Moved(tile.Position.ToPixelVec()))
	}

	return &Board{
		Layer: layer,
		batch: batch,
	}
}

// TODO: should a painter be bound to one palette? Not sure this makes sense
func NewPainter(paletteAtlasName string) (*Painter, error) {
	paletteAtlas, err := resources.LoadJSON[resources.PaletteData](fmt.Sprintf("atlases/%s", paletteAtlasName))

	if err != nil {
		return nil, err
	}

	pic, err := resources.LoadPNG(paletteAtlas.ImgSrc)

	if err != nil {
		return nil, err
	}

	palette := make(map[string]*pixel.Sprite)

	for _, texture := range paletteAtlas.Textures {
		palette[texture.Key] = pixel.NewSprite(pic, texture.Frame.ToPixelRect())
	}

	return &Painter{
		pic: pic,
		palette: palette,
	}, nil
}
