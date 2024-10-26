package managers

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/wirdos/actors"
	"github.com/wirdos/directors/input"
	"github.com/wirdos/resources"
	"github.com/wirdos/util"
)

type Game struct {
	player *Player
	camera *Camera

	input input.Input
	character *actors.Character

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

	g.camera.Render()
}

func NewGame(window *opengl.Window) (*Game, error) {
	// TODO: this setup should obviously not be hardcoded here
	center := window.Bounds().Center()

	walls := make([]pixel.Rect, 0)
	walls = append(walls, pixel.R(center.X+32, center.Y-16, center.X+64, center.Y+48))
	walls = append(walls, pixel.R(center.X, center.Y+48, center.X+32, center.Y+80))

	player := NewPlayer()
	input := input.NewKeyboardMouse(window)
	character, err := actors.NewCharacter(center, 72, walls)

	if err != nil {
		return nil, err
	}

	player.SetInput(input)
	player.SetPuppet(character)

	// TODO: this should come from the map/stage
	pic, err := resources.LoadPNG("wall")

	if err != nil {
		return nil, err
	}

	wallSprite := pixel.NewSprite(pic, pic.Bounds())

	camera := NewCamera(window, center, 4, []*util.Face{
		character.Face(),
		util.NewFace(1, wallSprite, pixel.V(center.X+48, center.Y)),
		util.NewFace(1, wallSprite, pixel.V(center.X+48, center.Y+32)),
		util.NewFace(1, wallSprite, pixel.V(center.X+16, center.Y+64)),
	})

	return &Game{
		player: player,
		camera: camera,
		input: input,
		character: character,
		window: window,
	}, nil
}
