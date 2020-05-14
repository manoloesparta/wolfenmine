package ray

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Map will be rendered at screen
type Map struct {
	Form [][]string
}

// Draw map into screen
func (m *Map) Draw() {
	pixelgl.Run(run)
}

func ray(color pixel.RGBA, start float64, end float64) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = color
	imd.Push(pixel.V(start, 300))
	imd.Push(pixel.V(end, 300))
	imd.Rectangle(600)
	return imd
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Wolfenmine",
		Bounds: pixel.R(0, 0, 600, 400),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic("Error at creating window")
	}

	win.Clear(colornames.Aliceblue)
	for !win.Closed() {
		(ray(pixel.RGBA{1, 0, 0, 1}, 0, 10)).Draw(win)
		win.Update()
	}
}
