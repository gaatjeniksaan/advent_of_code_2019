package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	tt := []struct {
		in1      []string
		in2      []string
		expected int
	}{
		{
			in1:      []string{"R8", "U5", "L5", "D3"},
			in2:      []string{"U7", "R6", "D4", "L4"},
			expected: 6,
		},
		{
			in1:      []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			in2:      []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			expected: 159,
		},
		{
			in1:      []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			in2:      []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			expected: 135,
		}}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, part1(tc.in1, tc.in2))
	}
}

func TestPart2(t *testing.T) {
	tt := []struct {
		in1      []string
		in2      []string
		expected int
	}{
		{
			in1:      []string{"R8", "U5", "L5", "D3"},
			in2:      []string{"U7", "R6", "D4", "L4"},
			expected: 30,
		},
		{
			in1:      []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			in2:      []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			expected: 610,
		},
		{
			in1:      []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			in2:      []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			expected: 410,
		}}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, part2(tc.in1, tc.in2))
	}
}
