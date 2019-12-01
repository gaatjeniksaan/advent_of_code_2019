package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testingSomething(t *testing.T) {
	tt := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tc := range tt {
		fmt.Println(tc)
		assert.Equal(t, tc.fuel, calculateFuel(tc.mass))
	}
}
