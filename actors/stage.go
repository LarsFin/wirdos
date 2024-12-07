package actors

// Is this the right place to put the stage, in 'actors'? Difficulty is that otherwise
// we design a cyclic package dependency as actors need to know about the stage and the
// stage needs to know about actors

import (
	"fmt"

	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/resources"
	"github.com/wirdos/util"
)

type Stage struct {
	Walls []pixel.Rect
	Boundary pixel.Rect
	// TODO: should be a list of dynamic faces not just one
	Character *Character
	Props []*Prop
	Boards []*util.Board
	spawnPoint pixel.Vec
}

func (s *Stage) SpawnPoint() pixel.Vec {
	return s.spawnPoint
}

// returns list of all drawable entities, note this list is unordered
func (s *Stage) GetDrawables() []util.Drawable {
	boardCount, propCount := len(s.Boards), len(s.Props)

	drawables := make([]util.Drawable, boardCount+propCount+1)

	for i, board := range s.Boards {
		drawables[i] = board
	}

	for i, prop := range s.Props {
		drawables[boardCount+i] = prop.Face()
	}

	drawables[boardCount+propCount] = s.Character.Face()

	return drawables
}

func LoadStage(path string, character *Character) (*Stage, error) {
	stageData, err := resources.LoadJSON[resources.StageData](fmt.Sprintf("stages/%s", path))

	if err != nil {
		return nil, err
	}

	props := make([]*Prop, len(stageData.Props))

	for i, propData := range stageData.Props {
		props[i], err = NewProp(propData.Key, propData.Position.ToPixelVec())
		if err != nil {
			return nil, err
		}
	}

	walls := make([]pixel.Rect, len(stageData.Walls))

	for i, wall := range stageData.Walls {
		walls[i] = wall.ToPixelRect()
	}

	painter, err := util.NewPainter(stageData.Palettes)

	if err != nil {
		return nil, err
	}

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
		Props: props,
		Boards: boards,
		spawnPoint: stageData.SpawnPoint.ToPixelVec(),
	}, nil
}
