package resources

import (
	"fmt"

	"github.com/gopxl/pixel/v2"
)

type SpriteSheetData struct {
	ImgSrc string `json:"sheetName"`
	Sprites []SpriteData `json:"sprites"`
}

type SpriteData struct {
	Key string `json:"key"`
	Frame SpriteFrameData `json:"frame"`
}

type SpriteFrameData struct {
	MinX float64 `json:"minX"`
	MinY float64 `json:"minY"`
	MaxX float64 `json:"maxX"`
	MaxY float64 `json:"maxY"`
}

func (sfd *SpriteFrameData) Rect() pixel.Rect {
	return pixel.R(sfd.MinX, sfd.MinY, sfd.MaxX, sfd.MaxY)
}

func GenerateSpriteMap(path string) (map[string]*pixel.Sprite, error) {
	data, err := LoadJSON[SpriteSheetData](fmt.Sprintf("atlases/%s", path))

	if err != nil {
		return nil, err
	}

	pic, err := LoadPNG(data.ImgSrc)

	if err != nil {
		return nil, err
	}

	sprites := make(map[string]*pixel.Sprite)

	for _, spriteData := range data.Sprites {
		sprites[spriteData.Key] = pixel.NewSprite(pic, spriteData.Frame.Rect())
	}

	return sprites, nil
}
