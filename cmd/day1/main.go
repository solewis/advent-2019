package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
)

func main() {
	fmt.Printf("Part 1: %d\n", calculateTotalFuelRequirement(calculateFuel))
	fmt.Printf("Part 2: %d\n", calculateTotalFuelRequirement(calculateModuleTotalFuel))
}

func calculateTotalFuelRequirement(calculator func(int) int) int {
	masses := parse.Ints("cmd/day1/input.txt", "\n")
	sum := 0
	for _, m := range masses {
		sum += calculator(m)
	}
	return sum
}

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

func calculateModuleTotalFuel(mass int) int {
	total := 0
	for {
		mass = calculateFuel(mass)
		if mass <= 0 {
			return total
		}
		total += mass
	}
}