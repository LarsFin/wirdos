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

// A palette is constructed from a slice of 'palette names', these point to pictures and have atlas information
// on where certain textures are located in the picture. The palette merges these sourced pictures into one
// picture and merges the atlas information into a map of textures. This allows for drawing sprites from different
// texture files onto the same batch as they have the same picture in memory.
// TODO: this could do with a refactor, its doing too much... also makes no sense for a new palette to receive
// a list of palette names
func NewPalette(paletteNames []string) (*Palette, error) {
	paletteAtlases := make([]*resources.PaletteData, len(paletteNames))
	pictures := make([]pixel.Picture, len(paletteNames))
	areas := make([]*pixel.Rect, len(paletteNames))
	dimensions := pixel.ZV

	for i, paletteName := range paletteNames {
		paletteAtlas, err := resources.LoadJSON[resources.PaletteData](fmt.Sprintf("atlases/%s", paletteName))

		if err != nil {
			return nil, err
		}

		paletteAtlases[i] = paletteAtlas
		pictures[i], err = resources.LoadPNG(paletteAtlas.ImgSrc)

		if err != nil {
			return nil, err
		}

		picDimensions := pictures[i].Bounds().Size()
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
					pic := pixel.PictureDataFromPicture(pictures[i])
					data.Pix[data.Index(point)] = pic.Pix[pic.Index(point.Sub(area.Min))]
				}
			}
		}
	}

	textures := make(map[string]*pixel.Sprite)

	for i, paletteAtlas := range paletteAtlases {
		for _, texture := range paletteAtlas.Textures {
			frame := texture.Frame.ToPixelRect()
			textures[texture.Key] = pixel.NewSprite(data, frame.Moved(areas[i].Min))
		}
	}

	return &Palette{
		Pic: data,
		Textures: textures,
	}, nil
}
