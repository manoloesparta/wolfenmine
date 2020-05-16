package ray

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const tileSize float64 = 32
const numRows float64 = 20
const numCols float64 = 30

const windowWitdth float64 = numCols * tileSize
const windowHeight float64 = numRows * tileSize

var colors map[string]pixel.RGBA = map[string]pixel.RGBA{
	"0": pixel.RGB(1, 1, 1),
	"1": pixel.RGB(0, 0, 0),
}

// Map will be rendered at screen
type Map struct {
	Grid   [][]string
	Window *pixelgl.Window
}

// Draw map into screen
func (m *Map) Draw() {
	for row := 0; float64(row) < numRows; row++ {
		for col := 0; float64(col) < numCols; col++ {
			key := m.Grid[row][col]

			xCoor := float64(col) * tileSize
			yCoor := float64(row) * tileSize

			square(colors[key], xCoor, yCoor).Draw(m.Window)
		}
	}
}

func square(color pixel.RGBA, x float64, y float64) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = color
	imd.Push(pixel.V(x, y), pixel.V(x+tileSize, y+tileSize))
	imd.Rectangle(0)
	return imd
}
