package raycast

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const wallWidth = 4
const numRays = windowWitdth / wallWidth

// Ray is for detecting walls
type Ray struct {
	Angle float64
	Win   *pixelgl.Window
	Main  *Player
	Grid  *Map

	WallHitX float64
	WallHitY float64
	Distance float64

	FacingDown  bool
	FacingRight bool
}

// NewRay creates a new ray the right way
func NewRay(angle float64, win *pixelgl.Window, pla *Player, gri *Map) Ray {
	angle = math.Mod(angle, 2*math.Pi)
	if angle < 0 {
		angle = (2 * math.Pi) + angle
	}

	return Ray{
		Angle: angle,
		Win:   win,
		Main:  pla,
		Grid:  gri,

		WallHitX: 0,
		WallHitY: 0,
		Distance: 0,

		FacingDown:  angle > 0 && angle < math.Pi,
		FacingRight: angle < 0.5*math.Pi || angle > 1.5*math.Pi,
	}
}

// Cast sends the ray to the wall
func (r *Ray) Cast(column float64) {
	yInter := math.Floor(r.Main.y/tileSize) * tileSize
	if r.FacingDown {
		yInter += 32
	}

	xInter := r.Main.x + (yInter-r.Main.y)/math.Tan(r.Angle)

	yStep := tileSize
	if !r.FacingDown {
		yStep *= -1
	}

	xStep := tileSize / math.Tan(r.Angle)
	if !r.FacingRight && xStep > 0 {
		xStep *= -1
	} else if r.FacingRight && xStep < 0 {
		xStep *= -1
	}

	xNext := xInter
	yNext := yInter
	if !r.FacingDown {
		yNext--
	}

	// horzWall := false
	xWallHit := 0.0
	yWallHit := 0.0

	for xNext >= 0 && xNext <= windowWitdth && yNext >= 0 && yNext <= windowHeight {
		if r.Grid.HasWallAt(xNext, yNext) {
			// horzWall = true
			xWallHit = xNext
			yWallHit = yNext

			line(pixel.RGB(1, 0, 0), r.Main.x, r.Main.y, xWallHit, yWallHit).Draw(r.Win)
			break
		} else {
			xNext += xStep
			yNext += yStep
		}
	}

}

// Draw a ray in the screen
func (r *Ray) Draw() {
	x, y := r.Main.Position()

	xAlmost := x + math.Cos(r.Angle)*30
	yAlmost := y + math.Sin(r.Angle)*30

	almost := line(pixel.RGB(1, 0, 0), x, y, xAlmost, yAlmost)
	almost.Draw(r.Win)
}
