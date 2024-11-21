package managers

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/wirdos/actors"
	"github.com/wirdos/directors/input"
	"github.com/wirdos/ui"
	"github.com/wirdos/util"
)

type Game struct {
	player *Player
	camera *Camera

	input input.Input
	character *actors.Character

	stage *actors.Stage

	ui *ui.UI

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

	g.character.Update()

	g.camera.Update()
	g.ui.Update()

	// TODO: workout render pipeline flow here, or at least improve structuring
	g.window.Clear(pixel.RGB(1, 1, 1))

	g.camera.Render()
	g.ui.Render()

	g.window.Update()
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

	stage, err := actors.LoadStage("stage", character)

	if err != nil {
		return nil, err
	}

	character.PlaceOnStage(stage)

	camera := NewCamera(window, center, 4, stage)
	ui, err := ui.NewUI(window)

	if err != nil {
		return nil, err
	}

	return &Game{
		player: player,
		camera: camera,
		input: input,
		character: character,
		window: window,
		ui: ui,
		stage: stage,
	}, nil
}
