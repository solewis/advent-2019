package main

import (
	"reflect"
	"testing"
)

func TestTotalEnergy(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		m1 := moon{point{-1, 0, 2}, point{0, 0, 0}}
		m2 := moon{point{2, -10, -7}, point{0, 0, 0}}
		m3 := moon{point{4, -8, 8}, point{0, 0, 0}}
		m4 := moon{point{3, 5, -1}, point{0, 0, 0}}
		moons := []*moon{&m1, &m2, &m3, &m4}
		result := calcTotalEnergyAfterSteps(moons, 10)
		expected := 179
		if result != expected {
			t.Errorf("Expected energy %d, but was %d", expected, result)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		m1 := moon{point{-8, -10, 0}, point{0, 0, 0}}
		m2 := moon{point{5, 5, 10}, point{0, 0, 0}}
		m3 := moon{point{2, -7, 3}, point{0, 0, 0}}
		m4 := moon{point{9, -8, -3}, point{0, 0, 0}}
		moons := []*moon{&m1, &m2, &m3, &m4}
		result := calcTotalEnergyAfterSteps(moons, 100)
		expected := 1940
		if result != expected {
			t.Errorf("Expected energy %d, but was %d", expected, result)
		}
	})
}

func TestCalcStepsUntilRepeat(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		m1 := moon{point{-1, 0, 2}, point{0, 0, 0}}
		m2 := moon{point{2, -10, -7}, point{0, 0, 0}}
		m3 := moon{point{4, -8, 8}, point{0, 0, 0}}
		m4 := moon{point{3, 5, -1}, point{0, 0, 0}}
		moons := []*moon{&m1, &m2, &m3, &m4}
		result := calcStepsUntilRepeat(moons)
		expected := 2772
		if result != expected {
			t.Errorf("Expected energy %d, but was %d", expected, result)
		}
	})
	t.Run("Example 2", func(t *testing.T) {
		m1 := moon{point{-8, -10, 0}, point{0, 0, 0}}
		m2 := moon{point{5, 5, 10}, point{0, 0, 0}}
		m3 := moon{point{2, -7, 3}, point{0, 0, 0}}
		m4 := moon{point{9, -8, -3}, point{0, 0, 0}}
		moons := []*moon{&m1, &m2, &m3, &m4}
		result := calcStepsUntilRepeat(moons)
		expected := 4686774924
		if result != expected {
			t.Errorf("Expected energy %d, but was %d", expected, result)
		}
	})
}

func TestGetMoonPairs(t *testing.T) {
	t.Run("2 moons has 1 pair", func(t *testing.T) {
		m1 := moon{point{1, 0, 2}, point{0, 0, 0}}
		m2 := moon{point{2, 10, 7}, point{0, 0, 0}}
		result := getMoonPairs([]*moon{&m1, &m2})
		expected := [][]moon{{m1, m2}}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but was %v", expected, result)
		}
	})
	t.Run("3 moons has 3 pair", func(t *testing.T) {
		m1 := moon{point{1, 0, 2}, point{0, 0, 0}}
		m2 := moon{point{2, 10, 7}, point{0, 0, 0}}
		m3 := moon{point{4, 8, 8}, point{0, 0, 0}}
		result := getMoonPairs([]*moon{&m1, &m2, &m3})
		expected := [][]moon{{m1, m2}, {m1, m3}, {m2, m3}}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but was %v", expected, result)
		}
	})
}
