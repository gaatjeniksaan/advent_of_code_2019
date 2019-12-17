package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := readInput()
	r1 := part1(input)
	fmt.Printf("Answer day 6.1: %v\n", r1)
	r2 := part2(input)
	fmt.Printf("Answer day 6.2: %v\n", r2)
}

func part1(input []string) int {
	outerToInner := make(map[string]string)

	// Fill outerToInner map
	for _, v := range input {
		i := strings.Split(v, ")")
		inner := i[0]
		outer := i[1]
		outerToInner[outer] = inner
	}

	result := calculateTotalOrbits(outerToInner)
	return result
}

func calculateTotalOrbits(outerToInner map[string]string) int {
	connections := make(map[string]int)
	result := 0
	totalItems := len(outerToInner)

	for {
		for k, v := range outerToInner {
			// If already a logged connection, continue
			if _, ok := connections[k]; ok {
				continue
			}

			// If "COM" is the value, we have the initial connection, log it
			if v == "COM" {
				connections[k] = 1
				continue
			}

			// If inner object is logged, we are a connection, thus add ourselves with +1 orbit value
			if nOrbits, ok := connections[v]; ok {
				connections[k] = nOrbits + 1
			}
		}

		if len(connections) == totalItems {
			break
		}
	}

	for _, v := range connections {
		result += v
	}

	return result
}

func part2(input []string) int {
	outerToInner := make(map[string]string)

	// Fill outerToInner map
	for _, v := range input {
		i := strings.Split(v, ")")
		inner := i[0]
		outer := i[1]
		outerToInner[outer] = inner
	}

	result := calculateSantaYouOrbits(outerToInner)

	return result
}

func calculateSantaYouOrbits(outerToInner map[string]string) int {
	santaPath := []string{}
	yourPath := []string{}

	santa := "SAN"
	you := "YOU"

	result := 0

	for {
		changed := false
		for k, v := range outerToInner {

			if k == santa {
				santaPath = append(santaPath, v)
				santa = v
				changed = true
			}

			if k == you {
				yourPath = append(yourPath, v)
				you = v
				changed = true
			}
		}

		if !changed {
			break
		}
	}

	for is, vs := range santaPath {
		for iy, vy := range yourPath {
			if vs == vy {
				// Found first common orbit
				result = is + iy
				return result
			}
		}
	}

	panic("no common orbit found...")
}

func readInput() []string {
	inputRaw, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err.Error())
	}

	inputString := string(inputRaw)
	input := strings.Split(inputString, "\n")

	return input
}
