package main

import "testing"

func TestRunPhases(t *testing.T) {
	t.Run("sample 1 - 1 phase", func(t *testing.T) {
		signal := "12345678"
		output := runPhases(signal, 1)

		expected := "48226158"
		if output != expected {
			t.Errorf("Expected %s, but was %s", expected, output)
		}
	})

	t.Run("sample 1 - 2 phases", func(t *testing.T) {
		signal := "12345678"
		output := runPhases(signal, 2)

		expected := "34040438"
		if output != expected {
			t.Errorf("Expected %s, but was %s", expected, output)
		}
	})

	t.Run("sample 1 - 3 phases", func(t *testing.T) {
		signal := "12345678"
		output := runPhases(signal, 3)

		expected := "03415518"
		if output != expected {
			t.Errorf("Expected %s, but was %s", expected, output)
		}
	})

	t.Run("sample 1 - 4 phases", func(t *testing.T) {
		signal := "12345678"
		output := runPhases(signal, 4)

		expected := "01029498"
		if output != expected {
			t.Errorf("Expected %s, but was %s", expected, output)
		}
	})

	t.Run("sample 2", func(t *testing.T) {
		signal := "80871224585914546619083218645595"
		output := runPhases(signal, 100)

		expected := "24176176"
		if output[:8] != expected {
			t.Errorf("Expected %s, but was %s", expected, output)
		}
	})

	t.Run("sample 3", func(t *testing.T) {
		signal := "19617804207202209144916044189917"
		output := runPhases(signal, 100)

		expected := "73745418"
		if output[:8] != expected {
			t.Errorf("Expected %s, but was %s", expected, output)
		}
	})

	t.Run("sample 4", func(t *testing.T) {
		signal := "69317163492948606335995924319873"
		output := runPhases(signal, 100)

		expected := "52432133"
		if output[:8] != expected {
			t.Errorf("Expected %s, but was %s", expected, output)
		}
	})
}

func Test_p2(t *testing.T) {
	t.Run("Sample 1", func(t *testing.T) {
		input := "03036732577212944063491565474664"
		expected := "84462026"
		actual := p2(input)
		if expected != actual {
			t.Errorf("Expected %s, but was %s", expected, actual)
		}
	})

	t.Run("Sample 2", func(t *testing.T) {
		input := "02935109699940807407585447034323"
		expected := "78725270"
		actual := p2(input)
		if expected != actual {
			t.Errorf("Expected %s, but was %s", expected, actual)
		}
	})

	t.Run("Sample 3", func(t *testing.T) {
		input := "03081770884921959731165446850517"
		expected := "53553731"
		actual := p2(input)
		if expected != actual {
			t.Errorf("Expected %s, but was %s", expected, actual)
		}
	})
}

func Test_GetMultiplier(t *testing.T) {
	basePattern := []int{0, 1, 0, -1}
	assertEqual(t, 1, getMultiplier(basePattern, 0, 0))
	assertEqual(t, 0, getMultiplier(basePattern, 0, 1))
	assertEqual(t, -1, getMultiplier(basePattern, 0, 2))
	assertEqual(t, 0, getMultiplier(basePattern, 0, 3))

	assertEqual(t, 0, getMultiplier(basePattern, 1, 0))
	assertEqual(t, 1, getMultiplier(basePattern, 1, 1))
	assertEqual(t, 1, getMultiplier(basePattern, 1, 2))
	assertEqual(t, 0, getMultiplier(basePattern, 1, 3))
	assertEqual(t, 0, getMultiplier(basePattern, 1, 4))
	assertEqual(t, -1, getMultiplier(basePattern, 1, 5))
	assertEqual(t, -1, getMultiplier(basePattern, 1, 6))
	assertEqual(t, 0, getMultiplier(basePattern, 1, 7))
}

func assertEqual(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Errorf("Expected %d, but was %d", expected, actual)
	}
}
