package util

import (
	"fmt"

	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources"
)

type Palette struct {
	Pic pixel.Picture
	Textures map[string]*pixel.Sprite
}

// A palette is a collection of textures with a key that determines the location of the texture within a picture.
// This can be useful for straightforward static sprites that don't need to be animated.
func NewPalette(paletteName string) (*Palette, error) {
	paletteAtlas, err := resources.LoadJSON[resources.PaletteData](fmt.Sprintf("atlases/%s", paletteName))

	if err != nil {
		return nil, err
	}

	picture, err := resources.LoadPNG(paletteAtlas.ImgSrc)

	if err != nil {
		return nil, err
	}

	data := pixel.PictureDataFromPicture(picture)
	textures := make(map[string]*pixel.Sprite)

	for _, texture := range paletteAtlas.Textures {
		textures[texture.Key] = pixel.NewSprite(data, texture.Frame.ToPixelRect())
	}

	return &Palette{
		Pic: data,
		Textures: textures,
	}, nil
}

// The pixel library can only draw sprites from the same picture onto a batch. This function takes a slice of palettes
// which have different pictures, then combines them into one picture and creates a new map of textures with updated
// frame information. This allows for optimised drawing of sprites from different texture files onto the same batch.
func CombinePalettes(palettes []*Palette) (*Palette, error) {
	areas := make([]*pixel.Rect, len(palettes))
	dimensions := pixel.ZV

	for i, palette := range palettes {
		picDimensions := palette.Pic.Bounds().Size()
		areas[i] = &pixel.Rect{
			Min: pixel.V(dimensions.X, 0),
			Max: pixel.V(dimensions.X + picDimensions.X, picDimensions.Y),
		}

		height := dimensions.Y
		if picDimensions.Y > height {
			height = picDimensions.Y
		}

		dimensions = pixel.V(
			dimensions.X + picDimensions.X,
			height,
		)
	}

	data := pixel.MakePictureData(pixel.R(0, 0, dimensions.X, dimensions.Y))

	for x := 0; x < int(dimensions.X); x++ {
		for y := 0; y < int(dimensions.Y); y++ {
			// find the picture that contains this pixel
			for i, area := range areas {
				point := pixel.V(float64(x), float64(y))
				// can't use contains as it considers max edges as containing but this doesn't translate to pixel coordinates
				if area.Min.X <= point.X && point.X < area.Max.X && area.Min.Y <= point.Y && point.Y < area.Max.Y {
					pic := pixel.PictureDataFromPicture(palettes[i].Pic)
					data.Pix[data.Index(point)] = pic.Pix[pic.Index(point.Sub(area.Min))]
				}
			}
		}
	}

	textures := make(map[string]*pixel.Sprite)

	for i, palette := range palettes {
		for key, texture := range palette.Textures {
			if _, ok := textures[key]; ok {
				return nil, fmt.Errorf("duplicate texture key when combining palettes: %s", key)
			}

			textures[key] = pixel.NewSprite(data, texture.Frame().Moved(areas[i].Min))
		}
	}

	return &Palette{
		Pic: data,
		Textures: textures,
	}, nil
}
