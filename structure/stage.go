package structure

import (
	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources/atlases"
	"github.com/wirdos/util"
)

type Stage struct {
	Walls []pixel.Rect
	// TODO: should be a list of dynamic faces not just one
	Face *util.Face
	// TODO: could have multiple screens, i.e; foreground or parallax and
	// so on (probably not the best name)
	Screen *util.Screen
}

func LoadStage(path string) (*Stage, error) {
	stageData, err := atlases.LoadStageData(path)

	if err != nil {
		return nil, err
	}

	walls := make([]pixel.Rect, len(stageData.Walls))

	for i, wall := range stageData.Walls {
		walls[i] = wall.ToPixelRect()
	}

	palette, err := atlases.LoadPalette(stageData.PaletteName)

	if err != nil {
		return nil, err
	}

	// TODO: work out screens against layers, maybe should just exist in data?
	screen := util.NewScreen(0)

	for _, tile := range stageData.Tiles {
		screen.AddSprite(palette[tile.Key], tile.Position.ToPixelVec())
	}

	return &Stage{
		Walls: walls,
		Character: character,
		Screen: screen,
	}, nil
}
