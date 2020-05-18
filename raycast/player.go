package raycast

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/faiface/pixel"
)

// Player is the camera
type Player struct {
	Position  pixel.Vec
	Direction pixel.Vec
	Plane     pixel.Vec
}

// NewPlayer is the constructor
func NewPlayer(posX, posY, dirX, dirY, planeX, planeY float64) Player {
	return Player{
		Position:  pixel.V(posX, posY),
		Direction: pixel.V(dirX, dirY),
		Plane:     pixel.V(planeX, planeY),
	}
}

// Turn around
func (p *Player) Turn(s float64) {
	p.Direction.Y = p.Direction.X*math.Sin(s) + p.Direction.Y*math.Cos(s)
	p.Direction.X = p.Direction.X*math.Cos(s) - p.Direction.Y*math.Sin(s)
	p.Plane.Y = p.Plane.X*math.Sin(s) + p.Plane.Y*math.Cos(s)
	p.Plane.X = p.Plane.X*math.Cos(s) - p.Plane.Y*math.Sin(s)
}

// MoveFront wherever you want
func (p *Player) MoveFront(s float64, grid [][]string) {
	if grid[int(p.Position.X+p.Direction.X*s)][int(p.Position.Y)] == "0" {
		p.Position.X += p.Direction.X * s
	}

	if grid[int(p.Position.X)][int(p.Position.Y+p.Direction.Y*s)] == "0" {
		p.Position.Y += p.Direction.Y * s
	}
}

// MoveBack wherever you want
func (p *Player) MoveBack(s float64, grid [][]string) {
	if grid[int(p.Position.X-p.Direction.X*s)][int(p.Position.Y)] == "0" {
		p.Position.X -= p.Direction.X * s
	}

	if grid[int(p.Position.X)][int(p.Position.Y-p.Direction.Y*s)] == "0" {
		p.Position.Y -= p.Direction.Y * s
	}
}

// MoveSideWays can move from left to right and viceversa
func (p *Player) MoveSideWays(s float64, grid [][]string) {
	if grid[int(p.Position.X-p.Plane.X*s)][int(p.Position.Y)] == "0" {
		p.Position.X -= p.Plane.X * s
	}

	if grid[int(p.Position.X)][int(p.Position.Y-p.Plane.Y*s)] == "0" {
		p.Position.Y -= p.Plane.Y * s
	}
}

// Cast makes the necesary calculations for casting
func (p *Player) Cast(grid [][]string) *pixel.Sprite {

	img := image.NewRGBA(image.Rect(0, 0, int(WindowWidth), int(WindowHeight)))

	ww := int(WindowWidth)
	wh := int(WindowHeight)

	draw.Draw(img, image.Rect(0, 0, ww, wh/2), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.Point{}, draw.Src)
	draw.Draw(img, image.Rect(0, wh/2, ww, wh), &image.Uniform{color.RGBA{64, 64, 64, 255}}, image.Point{}, draw.Src)

	for i := 0.0; i < WindowWidth; i++ {
		cameraX := (2 * i / WindowWidth) - 1

		stepX := 0.0
		sideDistX := 0.0
		mapX := p.Position.X
		rayDirX := p.Direction.X + p.Plane.X*cameraX
		deltaDistX := math.Abs(1 / rayDirX)

		if rayDirX < 0 {
			stepX = -1
			sideDistX = (p.Position.X - mapX) * deltaDistX
		} else {
			stepX = 1
			sideDistX = (mapX + 1 - p.Position.X) * deltaDistX
		}

		stepY := 0.0
		sideDistY := 0.0
		mapY := p.Position.Y
		rayDirY := p.Direction.Y + p.Plane.Y*cameraX
		deltaDistY := math.Abs(1 / rayDirY)

		if rayDirY < 0 {
			stepY = -1
			sideDistY = (p.Position.Y - mapY) * deltaDistY
		} else {
			stepY = 1
			sideDistY = (mapY + 1 - p.Position.Y) * deltaDistY
		}

		hit := 0.0
		side := 0.0

		for hit == 0 {
			if sideDistX < sideDistY {
				sideDistX += deltaDistX
				mapX += stepX
				side = 0
			} else {
				sideDistY += deltaDistY
				mapY += stepY
				side = 1
			}

			if grid[int(mapX)][int(mapY)] != "0" {
				hit = 1
			}
		}

		perpWallDist := 0.0
		if side == 0 {
			perpWallDist = (mapX - p.Position.X + (1-stepX)/2) / rayDirX
		} else {
			perpWallDist = (mapY - p.Position.Y + (1-stepY)/2) / rayDirY
		}

		lineHeight := WindowHeight / perpWallDist
		drawStart := -lineHeight/2 + WindowHeight/2
		if drawStart < 0 {
			drawStart = 0
		}

		drawEnd := lineHeight/2 + WindowHeight/2
		if drawEnd >= WindowHeight {
			drawEnd = WindowHeight - 1
		}

		color := Colors[grid[int(mapX)][int(mapY)]]
		if side == 1 {
			color.R = color.R / 2
			color.G = color.G / 2
			color.B = color.B / 2
		}

		for j := drawStart; j < drawEnd-1; j++ {
			img.Set(int(i), int(j), color)
		}
	}
	pic := pixel.PictureDataFromImage(img)
	return pixel.NewSprite(pic, pic.Bounds())
}
