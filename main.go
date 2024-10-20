package main

import (
	_ "image/png"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/wirdos/actors"
	"github.com/wirdos/input"
	"github.com/wirdos/resources"
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

	ipos := window.Bounds().Center()

	sprite.Draw(window, pixel.IM.Moved(ipos))
	character := actors.NewCharacter(ipos, 64)
	input := input.NewKeyboardMouse(window)

	for !window.Closed() {
		util.UpdateDeltaTime()

		input.Update()

		if input.Exit() {
			window.SetClosed(true)
		}

		window.Clear(pixel.RGB(1, 1, 1))

		character.Update(input.Direction())

		sprite.Draw(window, pixel.IM.Moved(character.Pos()))

		window.Update()
	}
}

func main() {
	opengl.Run(run)
}
