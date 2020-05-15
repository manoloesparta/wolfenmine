package main

import (
	"wolfenmine/ray"

	"github.com/faiface/pixel/pixelgl"
)

var space ray.Map = ray.Map{Grid: ray.LoadMap()}

func main() {
	pixelgl.Run(gameLoop)
}

func gameLoop() {
	win := ray.Setup()
	space.Draw(win)
	for !win.Closed() {
		win.Update()
	}
}
