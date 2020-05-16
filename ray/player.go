package ray

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"

	"github.com/faiface/pixel/pixelgl"
)

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
		moveSpeed:     3,
		rotationSpeed: 3 * (math.Pi / 180),
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

// Draw puts ourself in the map
func (p *Player) Draw() {
	p.rotationAngle += p.turnDirection * p.rotationSpeed

	step := p.walkDirection * p.moveSpeed
	p.x += math.Cos(p.rotationAngle) * step
	p.y += math.Sin(p.rotationAngle) * step

	dir := line(pixel.RGB(1, 0, 0), p.x, p.y, p.x+math.Cos(p.rotationAngle)*50, p.y+math.Sin(p.rotationAngle)*50)
	cir := circle(pixel.RGB(1, 0, 0), p.x, p.y, p.radius)

	dir.Draw(p.window)
	cir.Draw(p.window)
}

func circle(color pixel.RGBA, x float64, y float64, radius float64) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = color
	imd.Push(pixel.V(x, y))
	imd.Circle(radius, 0)
	return imd
}

func line(color pixel.RGBA, xStart float64, yStart float64, xEnd float64, yEnd float64) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = color
	imd.Push(pixel.V(xStart, yStart), pixel.V(xEnd, yEnd))
	imd.Line(2)
	return imd
}
