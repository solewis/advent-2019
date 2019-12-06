package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"math"
)

func main() {
	program := parse.Ints("cmd/day5/input.txt", ",")

	output := runIntcodeProgram(program, 1)
	fmt.Printf("Part 1: %d\n", output)

	output = runIntcodeProgram(program, 5)
	fmt.Printf("Part 2: %d\n", output)
}

func runIntcodeProgram(program []int, input int) int {
	memory := make([]int, len(program))
	copy(memory, program)

	var output int
	ip := 0
	for {
		opCode := memory[ip] % 100

		getValue := func(paramNum int) int {
			param := memory[ip+paramNum]
			mode := memory[ip] / int(math.Pow10(paramNum+1)) % 10
			if mode == 0 {
				return memory[param]
			}
			return param
		}

		switch opCode {

		case 99: //break
			return output

		case 1: //add
			dest := memory[ip+3]
			val1, val2 := getValue(1), getValue(2)
			memory[dest] = val1 + val2
			ip += 4

		case 2: //multiply
			dest := memory[ip+3]
			val1, val2 := getValue(1), getValue(2)
			memory[dest] = val1 * val2
			ip += 4

		case 3: //input
			parameter1 := memory[ip+1]
			memory[parameter1] = input
			ip += 2

		case 4: //output
			output = getValue(1)
			ip += 2

		case 5: //jump if true
			val1, val2 := getValue(1), getValue(2)
			if val1 != 0 {
				ip = val2
			} else {
				ip += 3
			}

		case 6: //jump if false
			val1, val2 := getValue(1), getValue(2)
			if val1 == 0 {
				ip = val2
			} else {
				ip += 3
			}

		case 7: //less than
			dest := memory[ip+3]
			val1, val2 := getValue(1), getValue(2)
			if val1 < val2 {
				memory[dest] = 1
			} else {
				memory[dest] = 0
			}
			ip += 4

		case 8: //equal
			dest := memory[ip+3]
			val1, val2 := getValue(1), getValue(2)
			if val1 == val2 {
				memory[dest] = 1
			} else {
				memory[dest] = 0
			}
			ip += 4
		}
	}
}
