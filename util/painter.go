package util

import (
	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources"
)

type Painter struct {
	palettes map[string]*Palette
}

func (p *Painter) PaintBoard(boardData resources.BoardData) (*Board, error) {
	palette, err := p.loadPalette(boardData.PaletteName)

	if err != nil {
		return nil, err
	}

	batch := pixel.NewBatch(&pixel.TrianglesData{}, palette.pic)

	// TODO: what happens if the key doesn't exist?
	for _, tile := range boardData.Tiles {
		sprite := palette.textures[tile.Key]
		sprite.Draw(batch, pixel.IM.Moved(tile.Position.ToPixelVec()))
	}

	return &Board{
		Layer: boardData.Layer,
		batch: batch,
	}, nil
}

func (p *Painter) loadPalette(paletteName string) (*Palette, error) {
	if palette := p.palettes[paletteName]; palette != nil {
		return palette, nil
	}

	palette, err := NewPalette(paletteName)

	if err != nil {
		return nil, err
	}

	p.palettes[paletteName] = palette
	return palette, nil
}

func NewPainter() *Painter {
	return &Painter{
		palettes: make(map[string]*Palette),
	}
}
