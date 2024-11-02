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
	Boundary pixel.Rect
	// TODO: should be a list of dynamic faces not just one
	Character *Character
	Boards []*util.Board
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

	painter := util.NewPainter()

	boards := make([]*util.Board, len(stageData.Boards))
	for i, boardData := range stageData.Boards {
		boards[i], err = painter.PaintBoard(boardData)
		if err != nil {
			return nil, err
		}
	}

	return &Stage{
		Walls: walls,
		Boundary: stageData.Boundary.ToPixelRect(),
		Character: character,
		Boards: boards,
		spawnPoint: stageData.SpawnPoint.ToPixelVec(),
	}, nil
}
