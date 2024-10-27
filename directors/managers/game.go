package managers

import (
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/wirdos/actors"
	"github.com/wirdos/directors/input"
	"github.com/wirdos/structure"
	"github.com/wirdos/util"
)

type Game struct {
	player *Player
	camera *Camera

	input input.Input
	character *actors.Character

	stage *structure.Stage

	window *opengl.Window
}

func (g *Game) Update() {
	// update delta time
	util.UpdateDeltaTime()

	// update input first
	g.input.Update()

	g.player.Update()

	// check whether game is requested to close
	if g.player.RequestsExit() {
		g.window.SetClosed(true)
		return
	}

	g.character.Update(g.stage)

	g.camera.Render()
}

func NewGame(window *opengl.Window) (*Game, error) {
	// TODO: this setup should obviously not be hardcoded here
	center := window.Bounds().Center()

	player := NewPlayer()
	input := input.NewKeyboardMouse(window)
	character, err := actors.NewCharacter(center, 72)

	if err != nil {
		return nil, err
	}

	player.SetInput(input)
	player.SetPuppet(character)

	stage := structure.LoadStage("test", character)

	camera := NewCamera(window, center, 4, []*util.Face{
		character.Face(),
	})

	return &Game{
		player: player,
		camera: camera,
		input: input,
		character: character,
		window: window,
	}, nil
}
