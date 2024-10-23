package main

import (
	_ "image/png"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/wirdos/actors"
	"github.com/wirdos/directors/input"
	"github.com/wirdos/resources"
	"github.com/wirdos/stage"
	"github.com/wirdos/util"
)

func run() {
	wcfg := opengl.WindowConfig{
		Title: "wirdos",
		Bounds: pixel.R(0, 0, 1240, 720),
	}

	window, err := opengl.NewWindow(wcfg)

	if err != nil {
		panic(err)
	}

	pic, err := resources.LoadPNG("sprite")

	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	pic, err = resources.LoadPNG("wall")

	if err != nil {
		panic(err)
	}

	wall := pixel.NewSprite(pic, pic.Bounds())

	ipos := window.Bounds().Center()
	window.SetMatrix(pixel.IM.Scaled(ipos, 4))

	character := actors.NewCharacter(ipos, 72, pixel.V(8, 16))
	input := input.NewKeyboardMouse(window)

	walls := make([]*stage.Wall, 0)

	walls = append(walls, stage.NewWall(pixel.R(ipos.X+32, ipos.Y-16, ipos.X+64, ipos.Y+48)))
	walls = append(walls, stage.NewWall(pixel.R(ipos.X, ipos.Y+48, ipos.X+32, ipos.Y+80)))

	for !window.Closed() {
		util.UpdateDeltaTime()

		input.Update()

		if input.Exit() {
			window.SetClosed(true)
		}

		window.Clear(pixel.RGB(1, 1, 1))

		character.Update(input.Direction(), walls)

		wall.Draw(window, pixel.IM.Moved(ipos.Add(pixel.V(48, 0))))
		wall.Draw(window, pixel.IM.Moved(ipos.Add(pixel.V(48, 32))))
		wall.Draw(window, pixel.IM.Moved(ipos.Add(pixel.V(16, 64))))

		sprite.Draw(window, pixel.IM.Moved(character.Pos()))

		window.Update()
	}
}

func main() {
	opengl.Run(run)
}
