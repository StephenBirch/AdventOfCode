package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	numberpad = [3][3]int{
		{7, 8, 9},
		{4, 5, 6},
		{1, 2, 3},
	}
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Boom: ", err)
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Kaboom: ", err)
	}

	readable := string(contents)
	fmt.Println(readable)

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
				if currentY < 2 {
					currentY++
				}
			case 'D':
				if currentY > 0 {
					currentY--
				}
			case 'L':
				if currentX > 0 {
					currentX--
				}
			case 'R':
				if currentX < 2 {
					currentX++
				}
			default:
				log.Fatalf("Unexpected char: %v", step)
			}
		}
		fmt.Println(numberpad[currentY][currentX])
	}
}
