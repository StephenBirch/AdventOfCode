package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	execute(2)
}

func execute(input int) {
	intCodes := strings.Split(input9, ",")
	var codes []int
	for _, intCode := range intCodes {
		code, _ := strconv.Atoi(intCode)
		codes = append(codes, code)
	}

	for i := len(codes); i < math.MaxInt16; i++ {
		codes = append(codes, 0)
	}

	var relativeBase int
	var eof bool
	var currentPosition int
	for !eof {
		//fmt.Println(codes[currentPosition])
		currentInstruction := codes[currentPosition]
		location1 := currentPosition + 1
		location2 := currentPosition + 2
		location3 := currentPosition + 3

		if len(codes) > currentPosition+1 && currentInstruction/100%10 == 0 {
			location1 = codes[currentPosition+1]
		}
		if len(codes) > currentPosition+1 && currentInstruction/200%5 == 1 {
			fmt.Println(codes[currentPosition], " new mode for 100's")
			location1 = codes[currentPosition+1] + relativeBase
		}

		if len(codes) > currentPosition+2 && currentInstruction/1000%10 == 0 {
			location2 = codes[currentPosition+2]
		}
		if len(codes) > currentPosition+2 && currentInstruction/2000%5 == 1 {
			fmt.Println(codes[currentPosition], " new mode for 1000's")

			location2 = codes[currentPosition+2] + relativeBase
		}

		if len(codes) > currentPosition+3 && currentInstruction/10000%10 == 0 {
			location3 = codes[currentPosition+3]
		}
		if len(codes) > currentPosition+3 && currentInstruction/20000%5 == 1 {
			fmt.Println(codes[currentPosition], " new mode for 10000's")

			location3 = codes[currentPosition+3] + relativeBase
		}

		switch currentInstruction % 100 {
		case 99:
			eof = true
			continue
		case 1:
			codes[location3] = codes[location1] + codes[location2]
			currentPosition += 4
		case 2:
			codes[location3] = codes[location1] * codes[location2]
			currentPosition += 4
		case 3:
			codes[location1] = input
			currentPosition += 2
		case 4:
			fmt.Println(codes[location1])
			currentPosition += 2

			// if index == 4 {
			// 	if codes[location1] > goat {
			// 		goat = codes[location1]										Day7
			// 	}
			// 	ins[0] <- codes[location1]
			// 	continue
			// }

			// ins[index+1] <- codes[location1]

		case 5:
			if codes[location1] != 0 {
				currentPosition = codes[location2]
				continue
			}
			currentPosition += 3
		case 6:
			if codes[location1] == 0 {
				currentPosition = codes[location2]
				continue
			}
			currentPosition += 3
		case 7:
			if codes[location1] < codes[location2] {
				codes[location3] = 1
			} else {
				codes[location3] = 0
			}
			currentPosition += 4
		case 8:
			if codes[location1] == codes[location2] {
				codes[location3] = 1
			} else {
				codes[location3] = 0
			}
			currentPosition += 4
		case 9:
			relativeBase += codes[location1]
			currentPosition += 2
		default:
			log.Fatalf("Unexpected intCode at position %d: %d", currentPosition, codes[currentPosition])
		}
	}
}

var goat int

func heapPermutation(a []int, size int, inputChan chan []int) {
	if size == 1 {
		temp := make([]int, len(a))
		copy(temp, a)
		inputChan <- temp
	}

	for i := 0; i < size; i++ {
		heapPermutation(a, size-1, inputChan)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

func createPermutations(input []int) [][]int {
	inputChan := make(chan []int)
	go func() {
		heapPermutation(input, len(input), inputChan)
		close(inputChan)
	}()

	var inputs [][]int
	for input := range inputChan {
		inputs = append(inputs, input)
	}

	return inputs
}

// func day7() {
// 	solutions := createPermutations([]int{5, 6, 7, 8, 9})

// 	for _, solution := range solutions {
// 		fmt.Println("New Solution: ", solution)
// 		inputChannels := []chan int{
// 			make(chan int, 1),
// 			make(chan int, 1),
// 			make(chan int, 1),
// 			make(chan int, 1),
// 			make(chan int, 1),
// 		}
// 		var wg sync.WaitGroup

// 		for index, s := range solution {
// 			wg.Add(1)

// 			go func(index, s int) {
// 				inputChannels[index] <- s
// 				wg.Done()
// 			}(index, s)
// 		}

// 		wg.Wait()

// 		for index := range solution {
// 			wg.Add(1)

// 			go func(index int) {
// 				execute(inputChannels, index)
// 				wg.Done()
// 			}(index)

// 			if index == 0 {
// 				inputChannels[index] <- 0
// 			}

// 		}
// 		wg.Wait()
// 	}

// 	fmt.Println(goat)
// }

var (
	testInput = "104,1125899906842624,99"
	input7    = "3,8,1001,8,10,8,105,1,0,0,21,30,47,64,81,98,179,260,341,422,99999,3,9,1001,9,5,9,4,9,99,3,9,1002,9,5,9,101,4,9,9,102,2,9,9,4,9,99,3,9,102,3,9,9,101,2,9,9,1002,9,3,9,4,9,99,3,9,1001,9,5,9,1002,9,3,9,1001,9,3,9,4,9,99,3,9,1002,9,3,9,101,2,9,9,102,5,9,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99"
	input9    = "1102,34463338,34463338,63,1007,63,34463338,63,1005,63,53,1101,3,0,1000,109,988,209,12,9,1000,209,6,209,3,203,0,1008,1000,1,63,1005,63,65,1008,1000,2,63,1005,63,904,1008,1000,0,63,1005,63,58,4,25,104,0,99,4,0,104,0,99,4,17,104,0,99,0,0,1102,1,344,1023,1101,0,0,1020,1101,0,481,1024,1102,1,1,1021,1101,0,24,1005,1101,0,29,1018,1102,39,1,1019,1102,313,1,1028,1102,1,35,1009,1101,28,0,1001,1101,26,0,1013,1101,0,351,1022,1101,564,0,1027,1102,1,32,1011,1101,23,0,1006,1102,1,25,1015,1101,21,0,1003,1101,0,31,1014,1101,33,0,1004,1102,37,1,1000,1102,476,1,1025,1101,22,0,1007,1102,30,1,1012,1102,1,27,1017,1102,1,34,1002,1101,38,0,1008,1102,1,36,1010,1102,1,20,1016,1102,567,1,1026,1102,1,304,1029,109,-6,2108,35,8,63,1005,63,201,1001,64,1,64,1106,0,203,4,187,1002,64,2,64,109,28,21101,40,0,-9,1008,1013,38,63,1005,63,227,1001,64,1,64,1105,1,229,4,209,1002,64,2,64,109,-2,1205,1,243,4,235,1105,1,247,1001,64,1,64,1002,64,2,64,109,-12,2102,1,-5,63,1008,63,24,63,1005,63,271,1001,64,1,64,1105,1,273,4,253,1002,64,2,64,109,8,2108,22,-9,63,1005,63,295,4,279,1001,64,1,64,1106,0,295,1002,64,2,64,109,17,2106,0,-5,4,301,1001,64,1,64,1106,0,313,1002,64,2,64,109,-21,21107,41,40,7,1005,1019,333,1001,64,1,64,1105,1,335,4,319,1002,64,2,64,109,1,2105,1,10,1001,64,1,64,1105,1,353,4,341,1002,64,2,64,109,10,1206,-3,371,4,359,1001,64,1,64,1105,1,371,1002,64,2,64,109,-5,21108,42,42,-7,1005,1011,393,4,377,1001,64,1,64,1105,1,393,1002,64,2,64,109,-8,2101,0,-4,63,1008,63,23,63,1005,63,415,4,399,1105,1,419,1001,64,1,64,1002,64,2,64,109,13,21102,43,1,-6,1008,1017,43,63,1005,63,441,4,425,1106,0,445,1001,64,1,64,1002,64,2,64,109,-21,1207,0,33,63,1005,63,465,1001,64,1,64,1106,0,467,4,451,1002,64,2,64,109,19,2105,1,3,4,473,1106,0,485,1001,64,1,64,1002,64,2,64,109,1,21101,44,0,-7,1008,1015,44,63,1005,63,511,4,491,1001,64,1,64,1106,0,511,1002,64,2,64,109,2,1206,-3,527,1001,64,1,64,1105,1,529,4,517,1002,64,2,64,109,-8,1201,-7,0,63,1008,63,35,63,1005,63,555,4,535,1001,64,1,64,1105,1,555,1002,64,2,64,109,1,2106,0,10,1105,1,573,4,561,1001,64,1,64,1002,64,2,64,109,4,21107,45,46,-7,1005,1014,591,4,579,1106,0,595,1001,64,1,64,1002,64,2,64,109,-12,1208,-6,21,63,1005,63,617,4,601,1001,64,1,64,1105,1,617,1002,64,2,64,109,-11,1208,6,31,63,1005,63,637,1001,64,1,64,1106,0,639,4,623,1002,64,2,64,109,16,2101,0,-7,63,1008,63,20,63,1005,63,659,1105,1,665,4,645,1001,64,1,64,1002,64,2,64,109,3,2102,1,-9,63,1008,63,38,63,1005,63,691,4,671,1001,64,1,64,1106,0,691,1002,64,2,64,109,4,1205,-1,703,1105,1,709,4,697,1001,64,1,64,1002,64,2,64,109,-14,21108,46,45,7,1005,1014,729,1001,64,1,64,1105,1,731,4,715,1002,64,2,64,109,7,21102,47,1,0,1008,1014,45,63,1005,63,755,1001,64,1,64,1106,0,757,4,737,1002,64,2,64,109,-12,2107,34,7,63,1005,63,775,4,763,1105,1,779,1001,64,1,64,1002,64,2,64,109,-5,1207,6,22,63,1005,63,797,4,785,1106,0,801,1001,64,1,64,1002,64,2,64,109,12,1202,0,1,63,1008,63,35,63,1005,63,827,4,807,1001,64,1,64,1105,1,827,1002,64,2,64,109,-5,1202,0,1,63,1008,63,36,63,1005,63,851,1001,64,1,64,1105,1,853,4,833,1002,64,2,64,109,-2,1201,4,0,63,1008,63,20,63,1005,63,873,1105,1,879,4,859,1001,64,1,64,1002,64,2,64,109,2,2107,22,-1,63,1005,63,899,1001,64,1,64,1106,0,901,4,885,4,64,99,21102,1,27,1,21101,0,915,0,1105,1,922,21201,1,53897,1,204,1,99,109,3,1207,-2,3,63,1005,63,964,21201,-2,-1,1,21101,0,942,0,1106,0,922,21202,1,1,-1,21201,-2,-3,1,21101,0,957,0,1105,1,922,22201,1,-1,-2,1105,1,968,22102,1,-2,-2,109,-3,2105,1,0"
)
