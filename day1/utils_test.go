package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateFuel(t *testing.T) {
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
		assert.Equal(t, tc.fuel, calculateFuel(tc.mass))
	}
}

func TestFuelTheFuel(t *testing.T) {
	tt := []struct {
		fuel       int
		additional int
	}{
		{2, 0},
		{654, 312},
		{33583, 16763},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.additional, fuelTheFuel(tc.fuel))
	}
}
