package resources

import (
	"embed"
	"fmt"
	"image"

	_ "image/png"

	"github.com/gopxl/pixel/v2"
)

//go:embed assets/*
var assets embed.FS

func LoadPNG(path string) (pixel.Picture, error) {
	path = fmt.Sprintf("assets/imgs/%s.png", path)

	file, err := assets.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

func LoadJSON[K any](path string) (*K, error) {
	path = fmt.Sprintf("assets/data/%s.json", path)
	data, err := assets.ReadFile(path)

	if err != nil {
		return nil, err
	}

	result, err := Deserialise[K](data, JSON)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func LoadToml[K any](path string) (*K, error) {
	path = fmt.Sprintf("assets/data/%s.toml", path)
	data, err := assets.ReadFile(path)

	if err != nil {
		return nil, err
	}

	result, err := Deserialise[K](data, TOML)

	if err != nil {
		return nil, err
	}

	return result, nil
}
