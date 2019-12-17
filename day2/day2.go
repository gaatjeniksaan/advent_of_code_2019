package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	rawInput := readInput()

	r1 := part1(setState(rawInput, 12, 2))
	fmt.Printf("Answer day 2.1: %v\n", r1)
	r2 := part2(rawInput)
	fmt.Printf("Answer day 2.2: %v\n", r2)
}

func setState(input []int, noun, verb int) []int {
	input[1] = noun
	input[2] = verb
	return input
}

func part1(rawInput []int) int {
	var result int
	pos := 0
	input := make([]int, len(rawInput))
	copy(input, rawInput)

	for {
		action := input[pos]
		if action == 99 {
			return input[0]
		}

		v1Pos := input[pos+1]
		v2Pos := input[pos+2]
		v1 := input[v1Pos]
		v2 := input[v2Pos]

		switch action {
		case 1:
			result = v1 + v2
		case 2:
			result = v1 * v2
		}

		resultPos := input[pos+3]
		input[resultPos] = result
		pos += 4
	}
}

func part2(rawInput []int) int {
	output := 19690720

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			input := make([]int, len(rawInput))
			copy(input, rawInput)

			out := part1(setState(input, noun, verb))
			if out == output {
				return 100*noun + verb
			}
		}
	}
	return -1
}

func readInput() []int {
	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err.Error())
	}
	rawInput := strings.Split(string(file), ",")
	input := make([]int, len(rawInput))
	// Convert str input to integers
	for i := range rawInput {
		input[i], err = strconv.Atoi(rawInput[i])
		if err != nil {
			panic(err.Error())
		}
	}

	return input
}
