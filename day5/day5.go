package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := readInput()
	fmt.Println(input)
	r1 := part1()
	fmt.Printf("Answer day 5.1: %v\n", r1)
	// r2 := part2()
	// fmt.Printf("Answer day 5.2: %v\n", r2)
}

func part1() int {
	return 1
}

func part2() int {
	return 1
}

func readInput() []int {
	var input []int
	inputRaw, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err.Error())
	}

	inputString := string(inputRaw)
	inputStrings := strings.Split(inputString, ",")
	for _, v := range inputStrings {
		i, _ := strconv.Atoi(v)
		input = append(input, i)
	}

	return input
}
