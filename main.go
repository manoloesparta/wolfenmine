package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	wolfmap := loadMap()
	fmt.Println(wolfmap)
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

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		win.Update()
	}
}

func loadMap() [][]string {
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
