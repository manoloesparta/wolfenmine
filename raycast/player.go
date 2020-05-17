package raycast

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const fov = 60 * (math.Pi / 180)

// Player is our point of view
type Player struct {
	x             float64
	y             float64
	window        *pixelgl.Window
	radius        float64
	turnDirection float64
	walkDirection float64
	rotationAngle float64
	moveSpeed     float64
	rotationSpeed float64
}

// NewPlayer creates a new player with default values
func NewPlayer(win *pixelgl.Window) Player {
	return Player{
		x:             windowWitdth / 2,
		y:             windowHeight / 2,
		window:        win,
		radius:        3,
		turnDirection: 0,
		walkDirection: 0,
		rotationAngle: math.Pi / 2,
		moveSpeed:     8,
		rotationSpeed: 8 * (math.Pi / 180),
	}
}

// Walk changes your walkdirection
func (p *Player) Walk(w float64) {
	p.walkDirection = w
}

// Turn changes your turn direction
func (p *Player) Turn(t float64) {
	p.turnDirection = t
}

// Position returns the current position
func (p *Player) Position() (float64, float64) {
	return p.x, p.y
}

// Draw puts ourself in the map
func (p *Player) Draw(m *Map) {
	p.rotationAngle += p.turnDirection * p.rotationSpeed

	step := p.walkDirection * p.moveSpeed
	xNew := p.x + math.Cos(p.rotationAngle)*step
	yNew := p.y + math.Sin(p.rotationAngle)*step

	if !m.HasWallAt(xNew, yNew) {
		p.x = xNew
		p.y = yNew
	}

	dirX := p.x + math.Cos(p.rotationAngle)*30
	dirY := p.y + math.Sin(p.rotationAngle)*30
	dir := line(pixel.RGB(1, 0, 0), p.x, p.y, dirX, dirY)
	dir.Draw(p.window)

	cir := circle(pixel.RGB(1, 0, 0), p.x, p.y, p.radius)
	cir.Draw(p.window)
}
