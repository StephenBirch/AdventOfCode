package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	blocks := strings.Split(string(input), "\n\n")

	letters := []byte{'a', 'b', 'c', 'x', 'y', 'z'}
	fmt.Println(letters[0], letters[5])

	var part1Score, part2Score int
	for _, block := range blocks {
		totalPeople := strings.Count(block, "\n") + 1

		var block1 int
		for i := byte('a'); i <= byte('z'); i++ {
			if strings.Contains(block, string(i)) {
				block1++
			}

			peopleBlocks := strings.Split(block, "\n")
			var people int
			for _, per := range peopleBlocks {
				if strings.Contains(per, string(i)) {
					people++
				}
			}

			if people == totalPeople {
				part2Score++
			}
		}
		part1Score += block1
	}
	fmt.Println("Part1:", part1Score)
	fmt.Println("Part2:", part2Score)
}
