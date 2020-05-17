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
	WasHitVert  bool
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
		WasHitVert:  false,
	}
}

// Cast sends the ray to the wall
func (r *Ray) Cast(column float64) {

	// HORIZONTAL

	horzWall := false
	xWallHit := 0.0
	yWallHit := 0.0

	yInter := math.Floor(r.Main.y/tileSize) * tileSize

	if r.FacingDown {
		yInter += tileSize
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

	for xNext >= 0 && xNext <= windowWitdth && yNext >= 0 && yNext <= windowHeight {
		if r.Grid.HasWallAt(xNext, yNext) {
			horzWall = true
			xWallHit = xNext
			yWallHit = yNext
			break
		} else {
			xNext += xStep
			yNext += yStep
		}
	}

	// VERTICAL

	vertWall := false
	xWallHit2 := 0.0
	yWallHit2 := 0.0

	xInter2 := math.Floor(r.Main.x/tileSize) * tileSize

	if r.FacingRight {
		xInter2 += tileSize
	}

	yInter2 := r.Main.y + (xInter2-r.Main.x)*math.Tan(r.Angle)

	xStep2 := tileSize
	if !r.FacingRight {
		xStep2 *= -1
	}

	yStep2 := tileSize * math.Tan(r.Angle)
	if !r.FacingDown && yStep2 > 0 {
		yStep2 *= -1
	} else if r.FacingDown && yStep2 < 0 {
		yStep2 *= -1
	}

	xNext2 := xInter2
	yNext2 := yInter2
	if !r.FacingRight {
		xNext2--
	}

	for xNext2 >= 0 && xNext2 <= windowWitdth && yNext2 >= 0 && yNext2 <= windowHeight {
		if r.Grid.HasWallAt(xNext2, yNext2) {
			vertWall = true
			xWallHit2 = xNext2
			yWallHit2 = yNext2
			break
		} else {
			xNext2 += xStep2
			yNext2 += yStep2
		}
	}

	horzDistance := math.MaxFloat64
	if horzWall {
		horzDistance = Pitagoras(r.Main.x, r.Main.y, xWallHit, yWallHit)
	}

	vertDistance := math.MaxFloat64
	if vertWall {
		vertDistance = Pitagoras(r.Main.x, r.Main.y, xWallHit2, yWallHit2)
	}

	if horzDistance < vertDistance {
		r.WallHitX = xWallHit
		r.WallHitY = yWallHit
		r.Distance = horzDistance
	} else {
		r.WallHitX = xWallHit2
		r.WallHitY = yWallHit2
		r.Distance = vertDistance
	}

	r.WasHitVert = vertDistance < horzDistance
	line(pixel.RGB(1, 0, 0), r.Main.x, r.Main.y, r.WallHitX, r.WallHitY).Draw(r.Win)
}

// Draw a ray in the screen
func (r *Ray) Draw() {
	x, y := r.Main.Position()

	xAlmost := x + math.Cos(r.Angle)*30
	yAlmost := y + math.Sin(r.Angle)*30

	almost := line(pixel.RGB(1, 0, 0), x, y, xAlmost, yAlmost)
	almost.Draw(r.Win)
}
