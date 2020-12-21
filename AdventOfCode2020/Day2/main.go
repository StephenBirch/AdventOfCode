package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	var part1Valid, part2Valid int
	for _, line := range lines {
		var min, max int
		var c byte
		var pw string
		fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &c, &pw)

		count := strings.Count(pw, string(c))

		if count <= max && count >= min {
			part1Valid++
		}

		if len(pw) < max-1 || min-1 < 0 {
			continue
		}

		if (string(pw[min-1]) == string(c)) != (string(pw[max-1]) == string(c)) {
			part2Valid++
		}
	}

	fmt.Println(part1Valid)
	fmt.Println(part2Valid)
}
