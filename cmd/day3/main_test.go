package main

import (
	"testing"
)

func TestDistanceToClosestIntersection(t *testing.T) {
	t.Run("sample 1", func(t *testing.T) {
		a := "R8,U5,L5,D3"
		b := "U7,R6,D4,L4"
		dist := distanceToClosestIntersection(a, b)
		expected := 6
		if dist != expected {
			t.Errorf("Expected distance %d but was %d", expected, dist)
		}
	})

	t.Run("sample 2", func(t *testing.T) {
		a := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
		b := "U62,R66,U55,R34,D71,R55,D58,R83"
		dist := distanceToClosestIntersection(a, b)
		expected := 159
		if dist != expected {
			t.Errorf("Expected distance %d but was %d", expected, dist)
		}
	})

	t.Run("sample 3", func(t *testing.T) {
		a := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
		b := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
		dist := distanceToClosestIntersection(a, b)
		expected := 135
		if dist != expected {
			t.Errorf("Expected distance %d but was %d", expected, dist)
		}
	})
}

func TestFewestStepsToIntersection(t *testing.T) {
	t.Run("sample 1", func(t *testing.T) {
		a := "R8,U5,L5,D3"
		b := "U7,R6,D4,L4"
		dist := fewestStepsToIntersection(a, b)
		expected := 30
		if dist != expected {
			t.Errorf("Expected steps %d but was %d", expected, dist)
		}
	})

	t.Run("sample 2", func(t *testing.T) {
		a := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
		b := "U62,R66,U55,R34,D71,R55,D58,R83"
		dist := fewestStepsToIntersection(a, b)
		expected := 610
		if dist != expected {
			t.Errorf("Expected steps %d but was %d", expected, dist)
		}
	})

	t.Run("sample 3", func(t *testing.T) {
		a := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
		b := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
		dist := fewestStepsToIntersection(a, b)
		expected := 410
		if dist != expected {
			t.Errorf("Expected steps %d but was %d", expected, dist)
		}
	})
}
