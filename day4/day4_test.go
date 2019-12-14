package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	tt := []struct {
		in1      int
		in2      int
		expected int
	}{
		{in1: 111111, in2: 111111, expected: 1},
		{in1: 223450, in2: 223450, expected: 0},
		{in1: 123789, in2: 123789, expected: 0},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, part1(tc.in1, tc.in2))
	}
}

func TestPart2(t *testing.T) {
	tt := []struct {
		in1      int
		in2      int
		expected int
	}{
		{in1: 112233, in2: 112233, expected: 1},
		{in1: 123444, in2: 123444, expected: 0},
		{in1: 111122, in2: 111122, expected: 1},
		{in1: 111111, in2: 111111, expected: 0},
		{in1: 111101, in2: 111101, expected: 0},
		{in1: 123789, in2: 123789, expected: 0},
		{in1: 377888, in2: 377888, expected: 1},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, part2(tc.in1, tc.in2), fmt.Sprintf("Failed for index %v", tc.in1))
	}
}
