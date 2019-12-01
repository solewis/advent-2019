package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", calculateTotalFuelRequirement(calculateFuel))
	fmt.Printf("Part 2: %d\n", calculateTotalFuelRequirement(calculateModuleTotalFuel))
}

func calculateTotalFuelRequirement(calculator func(int) int) int {
	masses := parse("cmd/day1/input.txt")
	sum := 0
	for _, m := range masses {
		mass, _ := strconv.Atoi(m)
		sum += calculator(mass)
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

func parse(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(dat), "\n")
}