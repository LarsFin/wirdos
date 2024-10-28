package actors

// Is this the right place to put the stage, in 'actors'? Difficulty is that otherwise
// we design a cyclic package dependency as actors need to know about the stage and the
// stage needs to know about actors
// TODO: ponder whether this is the right place for the stage

import (
	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources"
	"github.com/wirdos/util"
)

type Stage struct {
	Walls []pixel.Rect
	// TODO: should be a list of dynamic faces not just one
	Character *Character
	// TODO: could have multiple screens, i.e; foreground or parallax and
	// so on (probably not the best name)
	Board *util.Board
	spawnPoint pixel.Vec
}

func (s *Stage) SpawnPoint() pixel.Vec {
	return s.spawnPoint
}

func LoadStage(path string, character *Character) (*Stage, error) {
	stageData, err := resources.LoadJSON[resources.StageData](path)

	if err != nil {
		return nil, err
	}

	walls := make([]pixel.Rect, len(stageData.Walls))

	for i, wall := range stageData.Walls {
		walls[i] = wall.ToPixelRect()
	}

	painter, err := util.NewPainter(stageData.PaletteName)

	if err != nil {
		return nil, err
	}

	// TODO: work out screens against layers, maybe should just exist in data?
	board := painter.PaintBoard(0, stageData.Tiles)

	return &Stage{
		Walls: walls,
		Character: character,
		Board: board,
		spawnPoint: stageData.SpawnPoint.ToPixelVec(),
	}, nil
}
