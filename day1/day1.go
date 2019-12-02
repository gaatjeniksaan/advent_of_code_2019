package main

import (
	"bufio"
	"fmt"
	"github.com/gaatjeniksaan/advent_of_code_2019/day1/utils"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error reading input: %v", err)
	}
	defer file.Close()

	c1 := make(chan int)
	c2 := make(chan int)
	var total int
	var additional int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := scanner.Text()

		go func(input string) {
			mass, err := strconv.Atoi(input)
			if err != nil {
				panic(fmt.Sprintf("cannot convert input '%v' to integer: %v", mass, err))
			}
			fuel := calculateFuel(mass)
			c1 <- fuel
			c2 <- fuelTheFuel(fuel)
		}(input)

		total += <-c1
		additional += <-c2
	}
	fmt.Printf("Answer day 1.1: %v\n", total)
	fmt.Printf("Answer day 1.2: %v\n", total+additional)
}
