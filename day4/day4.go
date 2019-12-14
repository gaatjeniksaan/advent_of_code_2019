package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	i1, i2 := readInput()
	r1 := part1(i1, i2)
	fmt.Printf("Answer day 4.1: %v\n", r1)
	r2 := part2(i1, i2)
	fmt.Printf("Answer day 4.2: %v\n", r2)
}

func part1(i1, i2 int) int {
	var nPasswords int

	for i := i1; i <= i2; i++ {
		ruleDouble := false
		ruleDecrease := true

		s := strconv.Itoa(i)
		for j := 0; j < 5; j++ {
			c := s[j]
			n := s[j+1]

			if c > n {
				ruleDecrease = false
				break
			}

			if c == n {
				ruleDouble = true
			}
		}

		if ruleDouble && ruleDecrease {
			nPasswords++
		}
	}
	return nPasswords
}

func part2(i1, i2 int) int {
	var nPasswords int

	for i := i1; i <= i2; i++ {
		ruleDouble := false
		ruleDecrease := true

		s := strconv.Itoa(i)
		for j := 0; j < 5; j++ {
			c := s[j]
			n := s[j+1]

			if c > n {
				ruleDecrease = false
				break
			}

			if c == n {
				if (j == 0 || s[j-1] != c) && (j == 4 || c != s[j+2]) {
					ruleDouble = true
				}
			}
		}

		if ruleDouble && ruleDecrease {
			nPasswords++
			// fmt.Printf("considering '%v' a password\n", i)
		}
	}
	return nPasswords
}

func readInput() (int, int) {
	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err.Error())
	}

	inputRaw := string(file)
	input := strings.Split(inputRaw, "-")

	i1, _ := strconv.Atoi(input[0])
	i2, _ := strconv.Atoi(input[1])

	return i1, i2
}
