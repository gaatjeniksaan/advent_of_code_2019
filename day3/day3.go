package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x     int
	y     int
	steps int
}

func main() {
	i1, i2 := readInput()
	// r1 := part1(i1, i2)
	// fmt.Printf("Answer day 3.1: %v\n", r1)
	r2 := part2(i1, i2)
	fmt.Printf("Answer day 3.2: %v\n", r2)
}

func part1(i1, i2 []string) int {
	var smallest int
	started := false

	p1 := generatePositions(i1)
	p2 := generatePositions(i2)

	for k1, v1 := range p1 {
		for k2 := range p2 {
			if k1 == k2 {
				d := manhattenDistance(v1)
				if !started {
					smallest = d
					started = true
					continue
				}

				if d < smallest {
					smallest = d
				}
			}
		}
	}

	return smallest
}

func part2(i1, i2 []string) int {
	var smallest int
	started := false

	p1 := generatePositions(i1)
	p2 := generatePositions(i2)

	for k1, v1 := range p1 {
		for k2, v2 := range p2 {
			if k1 == k2 {
				d := v1.steps + v2.steps
				if !started {
					smallest = d
					started = true
					continue
				}

				if d < smallest {
					smallest = d
				}
			}
		}
	}

	return smallest
}

func (p *point) string() string {
	x := strconv.Itoa(p.x)
	y := strconv.Itoa(p.y)
	return x + "-" + y
}

func generatePositions(vals []string) map[string]point {
	positions := make(map[string]point)
	pos := point{0, 0, 0}
	n := 0

	for _, v := range vals {
		var horizontal bool
		var sign int

		dir := string(v[0])
		steps, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err.Error())
		}

		switch dir {
		case "L":
			sign = -1
			horizontal = true
		case "R":
			sign = 1
			horizontal = true
		case "U":
			sign = 1
			horizontal = false
		case "D":
			sign = -1
			horizontal = false
		default:
			panic("unknown direction " + dir)
		}

		for i := 1; i <= steps; i++ {
			n++
			if horizontal {
				pos = point{pos.x + sign, pos.y, n}
			} else {
				pos = point{pos.x, pos.y + sign, n}
			}

			positions[pos.string()] = pos
		}
	}

	return positions
}

func manhattenDistance(p1 point) int {
	p2 := point{0, 0, 0}
	x := abs(p1.x - p2.x)
	y := abs(p1.y - p2.y)
	return x + y
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func readInput() ([]string, []string) {
	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err.Error())
	}

	raw := strings.Split(string(file), "\n")
	i1 := raw[0]
	i2 := raw[1]

	i1 = strings.TrimSpace(i1)
	i2 = strings.TrimSpace(i2)

	i1s := strings.Split(i1, ",")
	i2s := strings.Split(i2, ",")

	return i1s, i2s
}
