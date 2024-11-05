package util

import (
	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources"
)

type Painter struct {
	palette *Palette
}

func (p *Painter) PaintBoard(boardData resources.BoardData) (*Board, error) {
	batch := pixel.NewBatch(&pixel.TrianglesData{}, p.palette.Pic)

	// TODO: what happens if the key doesn't exist?
	for _, tile := range boardData.Tiles {
		sprite := p.palette.Textures[tile.Key]
		sprite.Draw(batch, pixel.IM.Moved(tile.Position.ToPixelVec()))
	}

	return &Board{
		layer: boardData.Layer,
		batch: batch,
	}, nil
}

func NewPainter(paletteNames []string) (*Painter, error) {
	palette, err := NewPalette(paletteNames)

	if err != nil {
		return nil, err
	}

	return &Painter{
		palette: palette,
	}, nil
}
