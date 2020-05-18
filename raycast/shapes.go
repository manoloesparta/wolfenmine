package raycast

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// Line draws a line in the screen
func Line(color pixel.RGBA, index, start, end float64) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = color
	imd.Push(pixel.V(index, start), pixel.V(index, end))
	imd.Line(1)
	return imd
}
