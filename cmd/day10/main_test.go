package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestMostAsteroidsDetected(t *testing.T) {
	t.Run("Sample 1", func(t *testing.T) {
		input := `.#..#
.....
#####
....#
...##`
		result := mostAsteroidsDetected(strings.Split(input, "\n"))
		expected := 8
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})

	t.Run("Sample 2", func(t *testing.T) {
		input := `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`
		result := mostAsteroidsDetected(strings.Split(input, "\n"))
		expected := 33
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})

	t.Run("Sample 4", func(t *testing.T) {
		input := `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`
		result := mostAsteroidsDetected(strings.Split(input, "\n"))
		expected := 35
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})

	t.Run("Sample 5", func(t *testing.T) {
		input := `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`
		result := mostAsteroidsDetected(strings.Split(input, "\n"))
		expected := 41
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})

	t.Run("Sample 6", func(t *testing.T) {
		input := `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`
		result := mostAsteroidsDetected(strings.Split(input, "\n"))
		expected := 210
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})
}

func TestGreatestCommonFactor(t *testing.T) {
	t.Run("0 and 5 is 5", func(t *testing.T) {
		result := greatestCommonFactor(0, 5)
		expected := 5
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})

	t.Run("6 and 9 is 3", func(t *testing.T) {
		result := greatestCommonFactor(6, 9)
		expected := 3
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})
	t.Run("14 and 21 is 7", func(t *testing.T) {
		result := greatestCommonFactor(14, 21)
		expected := 7
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})
	t.Run("4 and 2 is 2", func(t *testing.T) {
		result := greatestCommonFactor(4, 2)
		expected := 2
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})
}

func TestPointsBetween(t *testing.T) {
	t.Run("0,0 to 4,2", func(t *testing.T) {
		points := pointsBetween(point{0, 0}, point{4, 2})
		expected := []point{{2, 1}}
		if !reflect.DeepEqual(points, expected) {
			t.Errorf("Expected points %v, but got %v", expected, points)
		}
	})

	t.Run("4,2 to 0,0", func(t *testing.T) {
		points := pointsBetween(point{4, 2}, point{0, 0})
		expected := []point{{2, 1}}
		if !reflect.DeepEqual(points, expected) {
			t.Errorf("Expected points %v, but got %v", expected, points)
		}
	})
	t.Run("4,4 to 2,6", func(t *testing.T) {
		points := pointsBetween(point{4, 4}, point{2, 6})
		expected := []point{{3, 5}}
		if !reflect.DeepEqual(points, expected) {
			t.Errorf("Expected points %v, but got %v", expected, points)
		}
	})
	t.Run("0,0 to 3,12", func(t *testing.T) {
		points := pointsBetween(point{0, 0}, point{3, 12})
		expected := []point{{2, 8}, {1, 4}}
		if !listsEqualInAnyOrder(points, expected) {
			t.Errorf("Expected points %v, but got %v", expected, points)
		}
	})
}

func TestFind200thDestroyedAsteroid(t *testing.T) {
	t.Run("Sample 1", func(t *testing.T) {
		input := `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`
		result := find200thDestroyedAsteroid(strings.Split(input, "\n"))
		expected := point{8, 2}
		if result != expected {
			t.Errorf("Expected %v, but was %v", expected, result)
		}
	})
}

func listsEqualInAnyOrder(a, b []point) bool {
	aMap := make(map[point]int)
	bMap := make(map[point]int)

	for _, xElem := range a {
		aMap[xElem]++
	}
	for _, yElem := range b {
		bMap[yElem]++
	}

	for aMapKey, aMapVal := range aMap {
		if bMap[aMapKey] != aMapVal {
			return false
		}
	}
	return true
}
