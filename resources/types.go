package resources

import "github.com/gopxl/pixel/v2"

// marshable struct which can be deserialised and converted to a pixel.Rect
type Rect struct {
	MinX float64 `json:"minX"`
	MinY float64 `json:"minY"`
	MaxX float64 `json:"maxX"`
	MaxY float64 `json:"maxY"`
}

func (r *Rect) ToPixelRect() pixel.Rect {
	return pixel.R(r.MinX, r.MinY, r.MaxX, r.MaxY)
}

// marshable struct which can be deserialised and converted to a pixel.Vec
type Vec struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (v *Vec) ToPixelVec() pixel.Vec {
	return pixel.V(v.X, v.Y)
}

type StageData struct {
	SpawnPoint Vec `json:"spawnPoint"`
	Walls []Rect `json:"walls"`
	Boards []BoardData `json:"boards"`
}

type BoardData struct {
	Layer int8 `json:"layer"`
	PaletteName string `json:"palette"`
	Tiles []TileData `json:"tiles"`
}

type TileData struct {
	Key string `json:"key"`
	Position Vec `json:"position"`
}

type PaletteData struct {
	ImgSrc string `json:"sheetName"`
	Textures []TextureData `json:"textures"`
}

type TextureData struct {
	Key string `json:"key"`
	Frame Rect `json:"frame"`
}
