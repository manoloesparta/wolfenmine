package raycast

import (
	"io/ioutil"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	numRows float64 = 24
	numCols float64 = 24

	// WindowWidth constant
	WindowWidth float64 = 640
	// WindowHeight constant
	WindowHeight float64 = 400
)

// Colors for drawing a pixel
var Colors map[string]pixel.RGBA = map[string]pixel.RGBA{
	"0": pixel.RGB(1, 1, 1),
	"1": pixel.RGB(0.2, 0.2, 0.2),
	"2": pixel.RGB(1, 0, 0),
	"3": pixel.RGB(0, 0, 1),
	"4": pixel.RGB(0.5, 0.2, 0.8),
}

// LoadMap returns the content of the map file
func LoadMap() [][]string {
	file, err := ioutil.ReadFile("map")
	if err != nil {
		panic("Error at opening map file")
	}

	content := strings.Split(string(file), "\n")
	var result [][]string

	for _, val := range content {
		row := strings.Split(val, " ")
		result = append(result, row)
	}

	return result
}

// Setup returns the window object
func Setup() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Wolfenmine",
		Bounds: pixel.R(0, 0, WindowWidth, WindowHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic("Error at creating window")
	}

	return win
}

// KeyPressed is for checking if a update should be done
func KeyPressed(win *pixelgl.Window) bool {
	expr1 := win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyS)
	expr2 := win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyD)
	expr3 := win.Pressed(pixelgl.KeyRight) || win.Pressed(pixelgl.KeyLeft)
	return expr1 || expr2 || expr3
}
