package main

import (
	_ "image/png"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/wirdos/v2/resources"
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

	r1 := pixel.R(0, 0, 100, 100)
	r2 := pixel.R(50, 50, 150, 150)

	r1.Intersect(r2)

	window.Clear(pixel.RGB(0, 0, 0))
	sprite.Draw(window, pixel.IM.Moved(window.Bounds().Center()))

	for !window.Closed() {
		window.Update()
	}
}

func main() {
	opengl.Run(run)
}
