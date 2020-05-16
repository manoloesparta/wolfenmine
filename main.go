package main

import (
	"time"
	"wolfenmine/raycast"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(gameLoop)
}

func gameLoop() {
	win := raycast.Setup()
	area := raycast.Map{Grid: raycast.LoadMap(), Window: win}
	player := raycast.NewPlayer(win)

	area.Draw()
	player.Draw(&area)

	for !win.Closed() {
		if raycast.KeyPressed(win) {
			if win.Pressed(pixelgl.KeyW) {
				player.Walk(1)
			}
			if win.Pressed(pixelgl.KeyS) {
				player.Walk(-1)
			}
			if win.Pressed(pixelgl.KeyA) {
				player.Turn(1)
			}
			if win.Pressed(pixelgl.KeyD) {
				player.Turn(-1)
			}

			area.Draw()
			player.Draw(&area)
		}

		player.Walk(0)
		player.Turn(0)

		time.Sleep(time.Millisecond * 10)
		win.Update()
	}
}
