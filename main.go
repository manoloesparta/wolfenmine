package main

import (
	"io/ioutil"
	"strings"
	"wolfenmine/ray"
)

func main() {
	wolfmap := ray.Map{Form: loadMap()}
	wolfmap.Draw()
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
