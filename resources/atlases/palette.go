package atlases

import (
	"fmt"

	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources"
)

type PaletteData struct {
	ImgSrc string `json:"sheetName"`
	Textures []TextureData `json:"textures"`
}

type TextureData struct {
	Key string `json:"key"`
	Frame Rect `json:"frame"`
}

func LoadPalette(path string) (map[string]*pixel.Sprite, error) {
	data, err := resources.LoadJSON[PaletteData](fmt.Sprintf("atlases/%s", path))

	if err != nil {
		return nil, err
	}

	pic, err := resources.LoadPNG(data.ImgSrc)

	if err != nil {
		return nil, err
	}

	palette := make(map[string]*pixel.Sprite)

	for _, textureData := range data.Textures {
		palette[textureData.Key] = pixel.NewSprite(pic, textureData.Frame.ToPixelRect())
	}

	return palette, nil
}
