package util

import (
	"fmt"

	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources"
)

type Palette struct {
	pic pixel.Picture
	textures map[string]*pixel.Sprite
}

func NewPalette(paletteName string) (*Palette, error) {
	paletteAtlas, err := resources.LoadJSON[resources.PaletteData](fmt.Sprintf("atlases/%s", paletteName))

	if err != nil {
		return nil, err
	}

	pic, err := resources.LoadPNG(paletteAtlas.ImgSrc)

	if err != nil {
		return nil, err
	}

	textures := make(map[string]*pixel.Sprite)

	for _, texture := range paletteAtlas.Textures {
		textures[texture.Key] = pixel.NewSprite(pic, texture.Frame.ToPixelRect())
	}

	return &Palette{
		pic: pic,
		textures: textures,
	}, nil
}
