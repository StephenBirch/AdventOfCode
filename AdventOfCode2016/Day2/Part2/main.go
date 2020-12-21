package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	numberpad = [][]rune{
		{0, 0, 1, 0, 0},
		{0, 2, 3, 4, 0},
		{5, 6, 7, 8, 9},
		{0, 'A', 'B', 'C', 0},
		{0, 0, 'D', 0, 0},
	}
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Boom: %v", err)
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Kaboom: %v", err)
	}

	readable := string(contents)
	fmt.Println(readable)

	var x rune
	fmt.Println(x)

	// 	readable = `ULL
	// RRDDD
	// LURDL
	// UUUUD`
	lines := strings.Split(readable, "\n")

	fmt.Println(lines)
	currentX := 1
	currentY := 1
	for _, line := range lines {
		for _, step := range line {
			switch step {
			case 'U':
				if currentY < 4 && numberpad[currentY+1][currentX] != 0 {
					currentY++
				}
			case 'D':
				if currentY > 0 && numberpad[currentY-1][currentX] != 0 {
					currentY--
				}
			case 'L':
				if currentX > 0 && numberpad[currentY][currentX-1] != 0 {
					currentX--
				}
			case 'R':
				if currentX < 4 && numberpad[currentY][currentX+1] != 0 {
					currentX++
				}
			default:
				log.Fatalf("Unexpected char: %v", step)
			}
		}
		fmt.Println(numberpad[currentY][currentX])
	}
}
