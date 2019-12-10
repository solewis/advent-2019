package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"math"
)

func main() {
	program := parse.Ints("cmd/day9/input.txt", ",")
	in, out := make(chan int), make(chan int)
	go runIntcodeProgram(program, "part1", in, out)
	in <- 1
	p1 := <- out
	fmt.Printf("Part 1: %v\n", p1)
}

func runIntcodeProgram(program []int, name string, in, out chan int) {
	memory := make(map[int]int)
	for ix, val := range program {
		memory[ix] = val
	}

	ip := 0
	relativeBase := 0
	for {
		opCode := memory[ip] % 100

		getFromMemory := func(address int) int {
			if val, exists := memory[address]; exists {
				return val
			}
			return 0
		}

		getWriteAddress := func(paramNum int) int {
			dest := getFromMemory(ip+paramNum)
			destMode := getFromMemory(ip) / int(math.Pow10(paramNum + 1)) % 10
			if destMode == 2 {
				dest += relativeBase
			}
			return dest
		}

		getValue := func(paramNum int) int {
			param := getFromMemory(ip+paramNum)
			mode := getFromMemory(ip) / int(math.Pow10(paramNum+1)) % 10
			if mode == 0 { //position mode
				return getFromMemory(param)
			} else if mode == 1 { //immediate mode
				return param
			}
			return getFromMemory(param+relativeBase) //relative mode
		}

		switch opCode {

		case 99: //break
			//close(in)
			//close(out)
			return

		case 1: //add
			dest := getWriteAddress(3)
			val1, val2 := getValue(1), getValue(2)
			memory[dest] = val1 + val2
			ip += 4

		case 2: //multiply
			dest := getWriteAddress(3)
			val1, val2 := getValue(1), getValue(2)
			memory[dest] = val1 * val2
			ip += 4

		case 3: //input
			dest := getWriteAddress(1)
			fmt.Printf("%s waiting for input\n", name)
			input := <-in
			fmt.Printf("%s got input %d\n", name, input)
			memory[dest] = input
			ip += 2

		case 4: //output
			output := getValue(1)
			fmt.Printf("%s waiting to output %d\n", name, output)
			out <- output
			fmt.Printf("%s sent output\n", name)
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
			dest := getWriteAddress(3)
			val1, val2 := getValue(1), getValue(2)
			if val1 < val2 {
				memory[dest] = 1
			} else {
				memory[dest] = 0
			}
			ip += 4

		case 8: //equal
			dest := getWriteAddress(3)
			val1, val2 := getValue(1), getValue(2)
			if val1 == val2 {
				memory[dest] = 1
			} else {
				memory[dest] = 0
			}
			ip += 4

		case 9: //adjust relative base
			val := getValue(1)
			relativeBase += val
			ip += 2
		}
	}
}
