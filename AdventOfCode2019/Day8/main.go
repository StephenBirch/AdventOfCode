package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	// input = ""
	layerSize := 25 * 6
	var currentLayer string
	var allLayers []string
	for index, c := range input {
		if index%layerSize == 0 && index != 0 {
			allLayers = append(allLayers, currentLayer)
			currentLayer = ""
		}
		currentLayer += string(c)
	}
	allLayers = append(allLayers, currentLayer)

	goat := math.MaxInt8
	goatValue := 0
	for _, layer := range allLayers {

		count := strings.Count(layer, "0")
		if count < goat {
			goat = count
			goatValue = strings.Count(layer, "1") * strings.Count(layer, "2")
		}
	}

	fmt.Println(goatValue)
	for y := 0; y < 6; y++ {

		for x := 0; x < 25; x++ {

			fmt.Printf(getPixelColour(y*25+x, allLayers))
		}
		fmt.Printf("\n")
	}
}

func getPixelColour(index int, layers []string) string {
	for _, layer := range layers {
		if string(layer[index]) == "2" {
			continue
		}
		if string(layer[index]) == "1" {
			return "#"
		}
		if string(layer[index]) == "0" {
			return "-"
		}
		return string(layer[index])
	}

	return "invalid"
}
