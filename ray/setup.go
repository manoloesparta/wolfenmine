package ray

import (
	"io/ioutil"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

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
		Bounds: pixel.R(0, 0, windowWitdth, windowHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic("Error at creating window")
	}

	win.Clear(colornames.Aliceblue)
	return win
}

// KeyPressed is for checking if a update should be done
func KeyPressed(win *pixelgl.Window) bool {
	expr1 := win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyS)
	expr2 := win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyD)
	return expr1 || expr2
}
