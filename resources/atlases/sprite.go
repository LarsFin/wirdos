package atlases

import (
	"fmt"

	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources"
)

type SpriteSheetData struct {
	ImgSrc string `json:"sheetName"`
	Sprites []SpriteData `json:"sprites"`
}

type SpriteData struct {
	Key string `json:"key"`
	Frame Rect `json:"frame"`
}

func GenerateSpriteMap(path string) (map[string]*pixel.Sprite, error) {
	data, err := resources.LoadJSON[SpriteSheetData](fmt.Sprintf("atlases/%s", path))

	if err != nil {
		return nil, err
	}

	pic, err := resources.LoadPNG(data.ImgSrc)

	if err != nil {
		return nil, err
	}

	sprites := make(map[string]*pixel.Sprite)

	for _, spriteData := range data.Sprites {
		sprites[spriteData.Key] = pixel.NewSprite(pic, spriteData.Frame.ToPixelRect())
	}

	return sprites, nil
}
