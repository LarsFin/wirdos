package managers

import (
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

		startDialogueEvent := g.eventPipeline.PullEventOfType(events.StartDialogue)

		if startDialogueEvent != nil {
			g.dialogue.BeginScript(startDialogueEvent.ResourceName)
			g.setState(InDialogue)
		}
	case InDialogue:
		g.ui.Update()

		endDialogueEvent := g.eventPipeline.PullEventOfType(events.EndDialogue)

		if endDialogueEvent != nil {
			g.setState(InPlay)
		}
	}

	g.camera.Update()
	g.camera.Render()
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

	ui, err := ui.NewUI(window)

	if err != nil {
		return nil, err
	}

	camera := NewCamera(window, center, 4, stage, ui)

	dialogue := NewDialogue(ui, eventPipeline)

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
