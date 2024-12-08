package main

import (
	_ "image/png"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/wirdos/directors/managers"
	"github.com/wirdos/logger"
)

func run() {
	wcfg := opengl.WindowConfig{
		Title:  "wirdos",
		Bounds: pixel.R(0, 0, 1240, 720),
	}

	window, err := opengl.NewWindow(wcfg)

	if err != nil {
		panic(err)
	}

	game, err := managers.NewGame(window)

	if err != nil {
		panic(err)
	}

	for !window.Closed() {
		game.Update()
	}
}

func main() {
	// initialise logger
	// TODO: determine config here to set level + output
	logger.InitLogger()

	opengl.Run(run)
}
