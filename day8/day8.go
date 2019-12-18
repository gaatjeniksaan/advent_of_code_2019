package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	transparent = "2"
	white       = "1"
	black       = "0"
)

func main() {
	input := readInput()
	// r1 := part1(input, 25, 6)
	// fmt.Printf("Answer day 8.1: %v\n", r1)
	part2(input, 25, 6)
}

func part1(input []string, wide, tall int) int {
	result := 0

	pixelPerLayer := wide * tall

	layer := 0
	layers := make(map[int][]string)

	for i, v := range input {
		if i%pixelPerLayer == 0 {
			layer++
			layers[layer] = []string{v}
			continue
		}
		r := layers[layer]
		layers[layer] = append(r, v)
	}

	zeroDigits := -1
	for _, v := range layers {
		zeros := 0
		ones := 0
		twos := 0

		for _, vl := range v {
			if vl == "0" {
				zeros++
			}
			if vl == "1" {
				ones++
			}
			if vl == "2" {
				twos++
			}
		}

		if zeroDigits == -1 {
			zeroDigits = zeros
			result = ones * twos
			continue
		}

		if zeros < zeroDigits {
			zeroDigits = zeros
			result = ones * twos
		}
	}

	return result
}

func part2(input []string, wide, tall int) {
	pixelPerLayer := wide * tall
	pixels := make(map[int]string)

	// Fill pixels map with transparent pixels
	for i := 0; i < pixelPerLayer; i++ {
		pixels[i] = transparent
	}

	layer := 0
	layers := make(map[int][]string)
	var maxLayer = 0

	for i, v := range input {
		if i%pixelPerLayer == 0 {
			layer++
			layers[layer] = []string{v}
			continue
		}
		r := layers[layer]
		layers[layer] = append(r, v)
		maxLayer = layer
	}

	// Now loop over each layer and fill pixels accordingly
	for l := 1; l <= maxLayer; l++ {
		layer, ok := layers[l]
		if !ok {
			panic(fmt.Errorf("cannot find key '%v' in layers", l))
		}

		for i, v := range layer {
			current := pixels[i]
			if current == transparent {
				pixels[i] = v
			}

			if current == white || current == black {
				continue
			}
		}
	}

	// Now print the bloody thing
	for i := 0; i < pixelPerLayer; i++ {
		if i%wide == 0 {
			fmt.Println()
		}

		v := pixels[i]
		switch v {
		case transparent:
			fmt.Printf(" ")
		case black:
			fmt.Printf("%c", '\u2B1B')
		case white:
			fmt.Printf("%c", '\u2588')
		}
	}
	fmt.Println()
}

func readInput() []string {
	inputRaw, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err.Error())
	}

	input := strings.Split(string(inputRaw), "")

	return input
}
