package atlases

import (
	"fmt"

	"github.com/wirdos/resources"
)

type StageData struct {
	PaletteName string `json:"palette"`
	SpawnPoint Vec `json:"spawnPoint"`
	Walls []Rect `json:"walls"`
	Tiles []TileData `json:"tiles"`
}

type TileData struct {
	Key string `json:"key"`
	Position Vec `json:"position"`
}

func LoadStageData(path string) (*StageData, error) {
	data, err := resources.LoadJSON[StageData](fmt.Sprintf("atlases/%s", path))

	if err != nil {
		return nil, err
	}

	return data, nil
}
