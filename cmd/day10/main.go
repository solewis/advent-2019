package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"math"
	"sort"
)

func main() {
	input := parse.Lines("cmd/day10/input.txt")
	fmt.Printf("Part 1: %d\n", mostAsteroidsDetected(input))
	p2 := find200thDestroyedAsteroid(input)
	fmt.Printf("Part 2: %d\n", p2.x*100+p2.y)
}

func find200thDestroyedAsteroid(asteroids []string) point {
	asteroidMap := buildAsteroidMap(asteroids)

	//find base
	var base point
	max := 0
	for asteroid := range asteroidMap {
		c := countDetectedAsteroids(asteroid, asteroidMap)
		if c > max {
			max = c
			base = asteroid
		}
	}

	//build list of all asteroids (minus base)
	var asteroidList []point
	for a := range asteroidMap {
		if a == base {
			continue
		}
		asteroidList = append(asteroidList, a)
	}

	//sort asteroids based on laser line order, using proximity to base as secondary sort
	sort.Slice(asteroidList, func(i, j int) bool {
		//return true if i is before j
		a1 := asteroidList[i]
		a2 := asteroidList[j]
		a1Quad := getQuadrant(base, a1)
		a2Quad := getQuadrant(base, a2)
		if a1Quad == a2Quad {
			a1Slope := float64(a1.y-base.y) / float64(a1.x-base.x)
			a2Slope := float64(a2.y-base.y) / float64(a2.x-base.x)
			if a1Slope == a2Slope { //same slope, return closest one first
				return base.distanceTo(a1) < base.distanceTo(a2)
			}
			if a1Quad == 1 || a1Quad == 3 {
				if a1Slope == math.Inf(-1) || a1Slope == math.Inf(1) {
					//handle asteroids on same x coordinate as base
					return true
				}
				if a2Slope == math.Inf(-1) || a2Slope == math.Inf(1) {
					//handle asteroids on same x coordinate as base
					return false
				}
			}
			return a1Slope < a2Slope
			//return a1Slope > a2Slope
		}

		return a1Quad < a2Quad
	})

	var asteroidCircle [][]point
	//loop through sorted asteroids, and group all the ones that are in the same line
	//asteroidCircle represents each line of asteroids in order of how they will be destroyed
	var currentLine []point
	slopePtr := math.Inf(-1) //hard code to initial slope
	for _, asteroid := range asteroidList {
		slope := float64(asteroid.y-base.y) / float64(asteroid.x-base.x)
		if slope != slopePtr {
			asteroidCircle = append(asteroidCircle, currentLine)
			currentLine = []point{}
		}
		slopePtr = slope
		currentLine = append(currentLine, asteroid)
	}
	asteroidCircle = append(asteroidCircle, currentLine)

	deadAsteroidCount := 0
	// loop the circle and get the first asteroid in each line
	// then continue looping getting the next in each line if exists, until you have destroyed 200, then return that point
	for i := 0; true; i++ {
		for _, line := range asteroidCircle {
			if i < len(line) {
				deadAsteroidCount++
				if deadAsteroidCount == 200 {
					return line[i]
				}
			}
		}
	}
	return point{0, 0} //error
}

func getQuadrant(base, asteroid point) int {
	if asteroid.x >= base.x && asteroid.y < base.y {
		return 1
	}
	if asteroid.x > base.x && asteroid.y >= base.y {
		return 2
	}
	if asteroid.x <= base.x && asteroid.y > base.y {
		return 3
	}
	return 4
}

func mostAsteroidsDetected(asteroids []string) int {
	asteroidMap := buildAsteroidMap(asteroids)
	max := 0
	for asteroid := range asteroidMap {
		c := countDetectedAsteroids(asteroid, asteroidMap)
		if c > max {
			max = c
		}
	}
	return max
}

type point struct{ x, y int }

func (p point) distanceTo(o point) int {
	return abs(p.x-o.x) + abs(p.y-o.y)
}

func buildAsteroidMap(asteroids []string) map[point]bool {
	asteroidMap := map[point]bool{}
	for y, xRow := range asteroids {
		for x, val := range xRow {
			if val == '#' {
				asteroidMap[point{x, y}] = true
			}
		}
	}
	return asteroidMap
}

func countDetectedAsteroids(asteroid point, asteroidMap map[point]bool) int {
	count := 0
	for other := range asteroidMap {
		if other == asteroid {
			continue
		}
		if asteroidCanDetect(asteroid, other, asteroidMap) {
			count++
		}
	}
	return count
}

func asteroidCanDetect(a, b point, asteroidMap map[point]bool) bool {
	for _, p := range pointsBetween(a, b) {
		if asteroidMap[p] {
			return false
		}
	}
	return true
}

func pointsBetween(a, b point) []point {
	xd := b.x - a.x
	yd := b.y - a.y
	gcf := greatestCommonFactor(abs(xd), abs(yd))
	if gcf == 1 {
		return []point{}
	}

	//add the point that is one step from b to a, then recurse
	point := point{b.x - xd/gcf, b.y - yd/gcf}
	return append(pointsBetween(a, point), point)
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func greatestCommonFactor(a, b int) int {
	small := a
	large := b
	if a > b {
		small = b
		large = a
	}
	if small == 0 {
		return large
	}
	result := large - (small * (large / small))
	if result == 0 {
		return small
	}
	return greatestCommonFactor(small, result)
}
