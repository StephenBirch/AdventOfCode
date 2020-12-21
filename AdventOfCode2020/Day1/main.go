package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	var intArray []int
	for _, line := range lines {
		if line == "" {
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		intArray = append(intArray, i)
	}

	//part1
	for indexI, i := range intArray {
		for indexJ, j := range intArray {
			if indexI == indexJ {
				continue
			}

			if i+j == 2020 {
				fmt.Println(i * j)
			}
		}
	}

	//part2
	for indexI, i := range intArray {
		for indexJ, j := range intArray {
			if indexI == indexJ {
				continue
			}
			for indexL, l := range intArray {
				if indexL == indexJ || indexL == indexI {
					continue
				}

				if i+j+l == 2020 {
					fmt.Println(i * j * l)
				}
			}
		}
	}
}
