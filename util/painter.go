package util

import (
	"fmt"

	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources"
)

type Painter struct {
	palette *Palette
}

func (p *Painter) PaintBoard(boardData resources.BoardData) (*Board, error) {
	batch := pixel.NewBatch(&pixel.TrianglesData{}, p.palette.Pic)

	for _, tile := range boardData.Tiles {
		if _, ok := p.palette.Textures[tile.Key]; !ok {
			return nil, fmt.Errorf("texture with key %s not found in palette", tile.Key)
		}

		sprite := p.palette.Textures[tile.Key]
		sprite.Draw(batch, pixel.IM.Moved(tile.Position.ToPixelVec()))
	}

	return &Board{
		layer: boardData.Layer,
		batch: batch,
	}, nil
}

func NewPainter(paletteNames []string) (*Painter, error) {
	if len(paletteNames) == 0 {
		return nil, fmt.Errorf("no palette names provided")
	}

	palettes := make([]*Palette, len(paletteNames))
	var err error

	for i, paletteName := range paletteNames {
		palettes[i], err = NewPalette(paletteName)

		if err != nil {
			return nil, err
		}
	}

	var palette *Palette

	if len(palettes) == 1 {
		palette = palettes[0]
	} else {
		palette, err = CombinePalettes(palettes)

		if err != nil {
			return nil, err
		}
	}

	return &Painter{
		palette: palette,
	}, nil
}
