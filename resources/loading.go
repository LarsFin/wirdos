package resources

import (
	"embed"
	"encoding/json"
	"fmt"
	"image"
	"io"

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
	path = fmt.Sprintf("assets/%s.json", path)
	file, err := assets.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	var data K

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
