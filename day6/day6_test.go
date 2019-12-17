package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	tt := []struct {
		in       []string
		expected int
	}{
		{in: []string{"COM)XYZ", "XYZ)ABC"}, expected: 3},
		{in: []string{"COM)X", "X)Y", "Y)Z", "Z)XX"}, expected: 10},
		{in: []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}, expected: 42},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, part1(tc.in))
	}
}
