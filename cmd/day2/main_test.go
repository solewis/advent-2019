package main

import (
	"reflect"
	"testing"
)

func TestIntcodeRunner(t *testing.T) {
	t.Run("sample 1", func(t *testing.T) {
		program := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
		result := runIntcodeProgram(program, 9, 10)

		expected := []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected final program %v, but was %v", expected, result)
		}
	})

	t.Run("sample 2", func(t *testing.T) {
		program := []int{1, 0, 0, 0, 99}
		result := runIntcodeProgram(program, 0, 0)

		expected := []int{2, 0, 0, 0, 99}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected final program %v, but was %v", expected, result)
		}
	})

	t.Run("sample 3", func(t *testing.T) {
		program := []int{2, 3, 0, 3, 99}
		result := runIntcodeProgram(program, 3, 0)

		expected := []int{2, 3, 0, 6, 99}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected final program %v, but was %v", expected, result)
		}
	})

	t.Run("sample 4", func(t *testing.T) {
		program := []int{2, 4, 4, 5, 99, 0}
		result := runIntcodeProgram(program, 4, 4)

		expected := []int{2, 4, 4, 5, 99, 9801}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected final program %v, but was %v", expected, result)
		}
	})

	t.Run("sample 5", func(t *testing.T) {
		program := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
		result := runIntcodeProgram(program, 1, 1)

		expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected final program %v, but was %v", expected, result)
		}
	})
}
