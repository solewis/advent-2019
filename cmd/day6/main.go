package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"strings"
)

func main() {
	lines := parse.Lines("cmd/day6/input.txt")
	fmt.Printf("Part 1: %d\n", countTotalOrbits(lines))
	fmt.Printf("Part 2: %d\n", distanceToSanta(lines))
}

func distanceToSanta(orbitData []string) int {
	orbitMap := buildOrbitMap(orbitData)

	myPath := map[string]int{}
	santaPath := map[string]int{}

	steps := 0
	myPtr := "YOU"
	santaPtr := "SAN"
	for {
		//walk YOU and SAN back one step at a time
		myPtr = orbitMap[myPtr]
		myPath[myPtr] = steps

		santaPtr = orbitMap[santaPtr]
		santaPath[santaPtr] = steps

		//check if either YOU or SAN has reached a "common" orbit object, then just add the steps it took each to get there
		if steps, exists := myPath[santaPtr]; exists {
			return steps + santaPath[santaPtr]
		}
		if steps, exists := santaPath[myPtr]; exists {
			return steps + myPath[myPtr]
		}
		steps++
	}
}

func countTotalOrbits(orbitData []string) int {
	orbitMap := buildOrbitMap(orbitData)

	countObjectOrbits := func(o string) int {
		count := 0
		for {
			parent, hasParent := orbitMap[o]
			if !hasParent {
				return count
			}
			o = parent
			count++
		}
	}

	count := 0
	for obj := range orbitMap {
		count += countObjectOrbits(obj)
	}
	return count
}

func buildOrbitMap(orbitData []string) map[string]string {
	orbitMap := map[string]string{}
	for _, d := range orbitData {
		objects := strings.Split(d, ")")
		orbitMap[objects[1]] = objects[0]
	}
	return orbitMap
}
