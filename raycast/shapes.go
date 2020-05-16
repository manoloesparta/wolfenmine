package raycast

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func square(color pixel.RGBA, x float64, y float64) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = color
	imd.Push(pixel.V(x, y), pixel.V(x+tileSize, y+tileSize))
	imd.Rectangle(0)
	return imd
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
