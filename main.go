package main

import (
	"wolfenmine/raycast"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(gameLoop)
}

func gameLoop() {
	win := raycast.Setup()
	center := win.Bounds().Center()

	player := raycast.NewPlayer(22, 10, -1, 0, 0, 0.66)
	grid := raycast.LoadMap()

	sprite := player.Cast(grid)
	sprite.Draw(win, pixel.IM.Moved(center).Scaled(center, 3))

	for !win.Closed() {
		if raycast.KeyPressed(win) {

			if win.Pressed(pixelgl.KeyD) {
				player.Turn(-0.01)
			}

			if win.Pressed(pixelgl.KeyA) {
				player.Turn(0.01)
			}

			if win.Pressed(pixelgl.KeyW) {
				player.MoveFront(0.1, grid)
			}

			if win.Pressed(pixelgl.KeyS) {
				player.MoveBack(0.1, grid)
			}

			sprite := player.Cast(grid)
			sprite.Draw(win, pixel.IM.Moved(center).Scaled(center, 3))
		}
		win.Update()
	}
}
