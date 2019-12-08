package main

import "testing"

func TestOrbitCounter(t *testing.T) {
	t.Run("Sample 1", func(t *testing.T) {
		input := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
		result := countTotalOrbits(input)
		expected := 42
		if result != expected {
			t.Errorf("Expected %d total orbits, but got %d", expected, result)
		}
	})
}

func TestDistanceToSanta(t *testing.T) {
	t.Run("Sample 1", func(t *testing.T) {
		input := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}
		result := distanceToSanta(input)
		expected := 4
		if result != expected {
			t.Errorf("Expected %d orbital transfers, but got %d", expected, result)
		}
	})
}
