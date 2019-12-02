package main

func calculateFuel(mass int) int {
	fuel := (mass / 3) - 2
	return fuel
}

// Recursively calculate how much fuel is needed to fuel the fuel
func fuelTheFuel(mass int) int {
	fuel := calculateFuel(mass)
	if fuel <= 0 {
		return 0
	}

	return fuel + fuelTheFuel(fuel)
}
