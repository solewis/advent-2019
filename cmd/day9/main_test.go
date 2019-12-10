package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"testing"
	"time"
)

func TestIntcodeComputer(t *testing.T) {
	t.Run("Sample 1", func(t *testing.T) {
		//Takes no input and produces a copy of itself as output
		program := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
		in, out := make(chan int), make(chan int)
		go runIntcodeProgram(program, "my app", in, out)
		for i, expected := range program {
			result := <-out
			if result != expected {
				t.Errorf("Expected %d at output %d but was %d", expected, i, result)
			}
		}
	})

	t.Run("Sample 2", func(t *testing.T) {
		//Output a 16 digit number
		program := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
		in, out := make(chan int), make(chan int)
		go runIntcodeProgram(program, "my app", in, out)
		result := <-out
		expected := 1219070632396864
		if result != expected {
			t.Errorf("Expected %d but was %d", expected, result)
		}
	})

	t.Run("Sample 3", func(t *testing.T) {
		//Output number in middle
		program := []int{104, 1125899906842624, 99}
		in, out := make(chan int), make(chan int)
		go runIntcodeProgram(program, "my app", in, out)
		result := <-out
		expected := 1125899906842624
		if result != expected {
			t.Errorf("Expected %d but was %d", expected, result)
		}
	})

	t.Run("Day 5 sample 1", func(t *testing.T) {
		// program should output 1 if input == 8, otherwise output 0
		in1, in2, out1, out2 := make(chan int), make(chan int), make(chan int), make(chan int)
		go runIntcodeProgram([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, "run 1", in1, out1)
		go runIntcodeProgram([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, "run 2", in2, out2)
		in1 <- 8
		in2 <- 5
		if <-out1 != 1 {
			t.Errorf("Expected 1")
		}
		if <-out2 != 0 {
			t.Errorf("Expected 0")
		}
	})

	t.Run("Day 5 sample 2", func(t *testing.T) {
		// program should output 1 if input < 8, otherwise output 0
		in1, in2, out1, out2 := make(chan int), make(chan int), make(chan int), make(chan int)
		go runIntcodeProgram([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, "run 1", in1, out1)
		go runIntcodeProgram([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, "run 2", in2, out2)
		in1 <- 5
		in2 <- 8
		if <-out1 != 1 {
			t.Errorf("Expected 1")
		}
		if <-out2 != 0 {
			t.Errorf("Expected 0")
		}
	})

	t.Run("Run TEST diagnostics", func(t *testing.T) {
		program := parse.Ints("input.txt", ",")
		in, out := make(chan int), make(chan int)
		go runIntcodeProgram(program, "my app", in, out)
		in <- 1
		outputs := []int{}
	outer:
		for {
			select {
			case o := <-out:
				fmt.Println("reading")
				outputs = append(outputs, o)
			case <-time.After(1 * time.Second):
				fmt.Println("done")
				break outer
			}
		}
		for _, result := range outputs[:len(outputs)-1] {
			if result != 0 {
				t.Errorf("Error with opcode: %d", result)
			}
		}
	})
}
