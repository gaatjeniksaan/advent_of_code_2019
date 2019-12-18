package main

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	tt := []struct {
		input    []string
		wide     int
		tall     int
		expected int
	}{
		{input: []string{"1", "2", "3", "4", "0", "0"}, expected: 1, wide: 3, tall: 2},
		{input: []string{"1", "2", "1", "2", "0", "0"}, expected: 4, wide: 3, tall: 2},
		{input: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "1", "2"}, expected: 1, wide: 3, tall: 1},
		{input: []string{"1", "2", "0", "5", "0", "0"}, expected: 1, wide: 3, tall: 1},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, part1(tc.input, tc.wide, tc.tall))
	}

}
