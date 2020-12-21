package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	var highestID float64
	seats := make(map[float64]interface{})
	for _, line := range lines {
		var row, column float64
		for index, letter := range line[0:7] {
			switch letter {
			case 'F':
				continue
			case 'B':
				row += 64 / math.Pow(float64(2), float64(index))
			default:
				log.Fatal("Unexpected: ", letter)
			}
		}
		for index, letter := range line[7:10] {
			switch letter {
			case 'L':
				continue
			case 'R':
				column += 4 / math.Pow(float64(2), float64(index))
			default:
				log.Fatal("Unexpected: ", letter)
			}
		}
		id := row*8 + column

		if id > highestID {
			highestID = id
		}

		seats[id] = true
	}
	fmt.Println("Part1:", highestID)

	var previousExists bool
	for i := float64(0); i < highestID; i++ {
		_, nextExists := seats[i+1]
		if _, exists := seats[i]; !exists {
			if previousExists && nextExists {
				fmt.Println("Part2:", i)
			}
			previousExists = false
			continue
		}
		previousExists = true
	}

}
