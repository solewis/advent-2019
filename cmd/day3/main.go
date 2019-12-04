package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"strings"
)

func main() {
	lines := parse.Lines("cmd/day3/input.txt")
	fmt.Printf("Part 1: %d\n", distanceToClosestIntersection(lines[0], lines[1]))

	fmt.Printf("Part 2: %d\n", fewestStepsToIntersection(lines[0], lines[1]))
}

type point struct {
	x, y int
}

func (p *point) move(dir uint8) {
	switch dir {
	case 'R':
		p.x += 1
	case 'L':
		p.x -= 1
	case 'U':
		p.y += 1
	case 'D':
		p.y -= 1
	}
}

func (p point) distanceTo(o point) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	return abs(p.x-o.x) + abs(p.y-o.y)
}

func buildPointPath(path string) []point {
	moves := strings.Split(path, ",")
	var points []point
	current := point{0, 0}
	for _, m := range moves {
		dir := m[0]
		dist := parse.Int(m[1:])
		for i := 1; i <= dist; i++ {
			current.move(dir)
			points = append(points, current)
		}
	}
	return points
}

type wireIntersection struct {
	p     point
	steps int
}

func findIntersections(pathA, pathB []point) []wireIntersection {
	var ints []wireIntersection
	pathAMap := map[point]int{}
	for i, a := range pathA {
		pathAMap[a] = i + 1
	}
	for i, b := range pathB {
		steps, exists := pathAMap[b]
		if exists {
			ints = append(ints, wireIntersection{p: b, steps: steps + i + 1})
		}
	}
	return ints
}

func distanceToClosestIntersection(pathA, pathB string) int {
	pointsA := buildPointPath(pathA)
	pointsB := buildPointPath(pathB)

	ints := findIntersections(pointsA, pointsB)

	closest := ints[0].p
	center := point{0, 0}
	for _, i := range ints {
		if i.p.distanceTo(center) < closest.distanceTo(center) {
			closest = i.p
		}
	}
	return closest.distanceTo(center)
}

func fewestStepsToIntersection(pathA, pathB string) int {
	pointsA := buildPointPath(pathA)
	pointsB := buildPointPath(pathB)

	ints := findIntersections(pointsA, pointsB)

	minInt := ints[0]
	for _, i := range ints {
		if i.steps < minInt.steps {
			minInt = i
		}
	}
	return minInt.steps
}
