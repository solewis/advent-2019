package main

import (
	"testing"
)

func TestFuelCalculator(t *testing.T) {
	t.Run("mass 12 needs 2 fuel", func(t *testing.T) {
		result := calculateFuel(12)
		if result != 2 {
			t.Errorf("Expected fuel of 2 but was %d", result)
		}
	})

	t.Run("mass 14 needs 2 fuel", func(t *testing.T) {
		result := calculateFuel(14)
		if result != 2 {
			t.Errorf("Expected fuel of 2 but was %d", result)
		}
	})

	t.Run("mass 1969 needs 654 fuel", func(t *testing.T) {
		result := calculateFuel(1969)
		if result != 654 {
			t.Errorf("Expected fuel of 654 but was %d", result)
		}
	})

	t.Run("mass 100756 needs 33583 fuel", func(t *testing.T) {
		result := calculateFuel(100756)
		if result != 33583 {
			t.Errorf("Expected fuel of 33583 but was %d", result)
		}
	})
}

func TestModuleFuelCalculator(t *testing.T) {
	t.Run("mass 14 needs 2 fuel", func(t *testing.T) {
		result := calculateModuleTotalFuel(14)
		if result != 2 {
			t.Errorf("Expected fuel of 2 but was %d", result)
		}
	})

	t.Run("mass 1969 needs 966 fuel", func(t *testing.T) {
		result := calculateModuleTotalFuel(1969)
		if result != 966 {
			t.Errorf("Expected fuel of 966 but was %d", result)
		}
	})

	t.Run("mass 100756 needs 50346 fuel", func(t *testing.T) {
		result := calculateModuleTotalFuel(100756)
		if result != 50346 {
			t.Errorf("Expected fuel of 50346 but was %d", result)
		}
	})
}


