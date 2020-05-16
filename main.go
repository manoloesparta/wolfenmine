package main

import (
	"wolfenmine/ray"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(gameLoop)
}

func gameLoop() {
	win := ray.Setup()
	area := ray.Map{Grid: ray.LoadMap(), Window: win}
	player := ray.NewPlayer(win)

	area.Draw()
	player.Draw()

	for !win.Closed() {
		if ray.KeyPressed(win) {

			if win.Pressed(pixelgl.KeyW) {
				player.Walk(1)
			} else if win.Pressed(pixelgl.KeyS) {
				player.Walk(-1)
			} else if win.Pressed(pixelgl.KeyD) {
				player.Turn(1)
			} else if win.Pressed(pixelgl.KeyA) {
				player.Turn(-1)
			}

			if win.JustReleased(pixelgl.KeyW) || win.JustReleased(pixelgl.KeyS) {
				player.Walk(0)
			} else if win.JustReleased(pixelgl.KeyD) || win.JustReleased(pixelgl.KeyA) {
				player.Turn(0)
			}

			area.Draw()
			player.Draw()
		}
		win.Update()
	}
}
