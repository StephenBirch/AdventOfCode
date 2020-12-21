package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type variance struct {
	x, y int
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	variances := []variance{
		variance{
			x: 1,
			y: 1,
		},
		// variance{
		// 	x: 3,
		// 	y: 1,
		// },
		variance{
			x: 5,
			y: 1,
		},
		variance{
			x: 7,
			y: 1,
		},
		variance{
			x: 1,
			y: 2,
		},
	}

	// Part 1 (x=3,y=1) was 200, use that as the start for the part2 answer
	part2 := 200

	for _, vari := range variances {
		var trees, x int

		for j := 0; j < len(lines); j += vari.y {
			if j == 0 {
				continue
			}
			x += vari.x
			if lines[j][x%len(lines[j])] == '#' {
				trees++
			}
		}
		part2 *= trees
	}
	fmt.Println("Part2: ", part2)
}
