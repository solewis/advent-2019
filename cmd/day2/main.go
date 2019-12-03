package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
)

func runIntcodeProgram(program []int, noun, verb int) []int {
	memory := make([]int, len(program))
	copy(memory, program)
	memory[1], memory[2] = noun, verb
	ip := 0
	for {
		opCode := memory[ip]
		if opCode == 99 {
			return memory
		}
		addressA, addressB, addressC := memory[ip+1], memory[ip+2], memory[ip+3]
		switch opCode {
		case 1:
			memory[addressC] = memory[addressA] + memory[addressB]
		case 2:
			memory[addressC] = memory[addressA] * memory[addressB]
		}
		ip += 4
	}
}

func main() {
	memory := parse.Ints("cmd/day2/input.txt", ",")
	finalMemory := runIntcodeProgram(memory, 12, 2)
	fmt.Printf("Part 1: %d\n", finalMemory[0])

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			finalMemory := runIntcodeProgram(memory, noun, verb)
			if finalMemory[0] == 19690720 {
				fmt.Printf("Part 2: %d\n", 100*noun+verb)
				return
			}
		}
	}
}
