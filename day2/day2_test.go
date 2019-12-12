package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	var tt = []struct {
		in       []int
		expected int
	}{
		{[]int{1, 0, 0, 0, 99}, 2},
		{[]int{2, 3, 0, 3, 99}, 2},
		{[]int{2, 4, 4, 5, 99, 0}, 2},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 30},
	}

	for _, tc := range tt {
		assert.Equal(t, part1(tc.in), tc.expected)
	}
}
