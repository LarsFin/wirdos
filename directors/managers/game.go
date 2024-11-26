package managers

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/wirdos/actors"
	"github.com/wirdos/directors/input"
	"github.com/wirdos/events"
	"github.com/wirdos/ui"
	"github.com/wirdos/util"
)

type GameState int

const (
	InPlay GameState = iota
	InDialogue GameState= iota
)

type Game struct {
	player *Player
	camera *Camera

	input input.InputController
	character *actors.Character

	stage *actors.Stage

	ui *ui.UI

	dialogue *Dialogue

	state GameState
	window *opengl.Window

	eventPipeline *events.Pipeline
}

func (g *Game) Update() {
	// always update delta time and input first
	util.UpdateDeltaTime()
	g.input.Update()

	g.player.Update()
	if g.player.RequestsExit() {
		g.window.SetClosed(true)
		return
	}

	switch g.state {
	case InPlay:
		g.character.Update()

		dialogueEvent := events.PopEventOfType[*events.DialogueEvent](g.eventPipeline)

		if dialogueEvent != nil {
			// TODO: use dialogue event script name...
			g.dialogue.BeginScript()
			g.setState(InDialogue)
		}
	case InDialogue:
		g.ui.Update()
	}

	g.camera.Update()

	// TODO: work out render pipeline flow here, or at least improve structuring
	g.window.Clear(pixel.RGB(1, 1, 1))

	g.camera.Render()
	g.ui.Render()

	g.window.Update()
}

func (g *Game) setState(state GameState) {
	g.state = state

	switch state {
	case InDialogue:
		g.player.SetPuppet(g.dialogue)
		return;
	case InPlay:
		g.player.SetPuppet(g.character)
		return;
	}
}

func NewGame(window *opengl.Window) (*Game, error) {
	// TODO: this setup should obviously not be hardcoded here
	center := window.Bounds().Center()

	eventPipeline := events.NewPipeline()

	player := NewPlayer()
	input := input.NewKeyboardMouse(window)
	character, err := actors.NewCharacter(center, 72, eventPipeline)

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

	// TODO: this really makes me feel like dialogue shouldn't be a manager but I can't
	// quite work out why or what it should be instead...
	dialogue := NewDialogue(ui)

	return &Game{
		player: player,
		camera: camera,
		input: input,
		character: character,
		window: window,
		ui: ui,
		stage: stage,
		dialogue: dialogue,
		state: InPlay,
		eventPipeline: eventPipeline,
	}, nil
}
