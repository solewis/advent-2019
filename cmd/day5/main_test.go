package main

import (
	"testing"
)

func Test(t *testing.T) {

	t.Run("Sample 1", func(t *testing.T) {
		// program should output 1 if input == 8, otherwise output 0
		output1 := runIntcodeProgram([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 8)
		output2 := runIntcodeProgram([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 5)
		if output1 != 1 {
			t.Errorf("Expected 1")
		}
		if output2 != 0 {
			t.Errorf("Expected 0")
		}
	})
	t.Run("Sample 2", func(t *testing.T) {
		// program should output 1 if input < 8, otherwise output 0
		output1 := runIntcodeProgram([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 5)
		output2 := runIntcodeProgram([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 8)
		if output1 != 1 {
			t.Errorf("Expected 1")
		}
		if output2 != 0 {
			t.Errorf("Expected 0")
		}
	})
	t.Run("Sample 3", func(t *testing.T) {
		// program should output 1 if input == 8, otherwise output 0
		output1 := runIntcodeProgram([]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 8)
		output2 := runIntcodeProgram([]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 5)
		if output1 != 1 {
			t.Errorf("Expected 1")
		}
		if output2 != 0 {
			t.Errorf("Expected 0")
		}
	})
	t.Run("Sample 4", func(t *testing.T) {
		// program should output 1 if input < 8, otherwise output 0
		output1 := runIntcodeProgram([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 5)
		output2 := runIntcodeProgram([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 8)
		if output1 != 1 {
			t.Errorf("Expected 1")
		}
		if output2 != 0 {
			t.Errorf("Expected 0")
		}
	})
	t.Run("Sample 5", func(t *testing.T) {
		// program should output 0 if input is 0, otherwise output 1
		output1 := runIntcodeProgram([]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 5)
		output2 := runIntcodeProgram([]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 0)
		if output1 != 1 {
			t.Errorf("Expected 1")
		}
		if output2 != 0 {
			t.Errorf("Expected 0")
		}
	})
	t.Run("Sample 6", func(t *testing.T) {
		// program should output 0 if input is 0, otherwise output 1
		output1 := runIntcodeProgram([]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 5)
		output2 := runIntcodeProgram([]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 0)
		if output1 != 1 {
			t.Errorf("Expected 1")
		}
		if output2 != 0 {
			t.Errorf("Expected 0")
		}
	})
	t.Run("Sample 7", func(t *testing.T) {
		// program should output 999 if the input value is below 8, output 1000 if the input value is equal to 8, or output 1001 if the input value is greater than 8.
		output1 := runIntcodeProgram([]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 5)
		output2 := runIntcodeProgram([]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 8)
		output3 := runIntcodeProgram([]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 13)
		if output1 != 999 {
			t.Errorf("Expected 999")
		}
		if output2 != 1000 {
			t.Errorf("Expected 1000")
		}
		if output3 != 1001 {
			t.Errorf("Expected 1001")
		}
	})
}
