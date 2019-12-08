package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"math"
	"sync"
)

func main() {
	program := parse.Ints("cmd/day7/input.txt", ",")
	fmt.Printf("Part 1: %d\n", findMaxThrusterSignal(program))
	fmt.Printf("Part 2: %d\n", findMaxThrusterSignalWithFeedback(program))
}

func findThrusterSignal(program, phaseSettings []int) int {
	output := 0
	for i := 0; i < 5; i++ {
		output = runIntcodeProgram(program, []int{phaseSettings[i], output})
	}
	return output
}

func findThrusterSignalWithFeedback(program, phaseSettings []int) int {
	ic1, ic2, ic3, ic4, ic5 := make(chan int), make(chan int), make(chan int), make(chan int), make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		runContinuousIntcodeProgram(program, ic1, ic2)
		wg.Done()
	}()
	go func() {
		runContinuousIntcodeProgram(program, ic2, ic3)
		wg.Done()
	}()
	go func() {
		runContinuousIntcodeProgram(program, ic3, ic4)
		wg.Done()
	}()
	go func() {
		runContinuousIntcodeProgram(program, ic4, ic5)
		wg.Done()
	}()
	go func() {
		runContinuousIntcodeProgram(program, ic5, ic1)
	}()
	ic1 <- phaseSettings[0]
	ic2 <- phaseSettings[1]
	ic3 <- phaseSettings[2]
	ic4 <- phaseSettings[3]
	ic5 <- phaseSettings[4]
	ic1 <- 0

	//wait for the first for amplifiers to finish, then grab the final output of amp5 and return it
	wg.Wait()
	return <-ic1
}

func findMaxThrusterSignal(program []int) int {
	max := 0
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			for c := 0; c < 5; c++ {
				for d := 0; d < 5; d++ {
					for e := 0; e < 5; e++ {
						m := map[int]bool{a: true, b: true, c: true, d: true, e: true}
						if len(m) < 5 {
							continue
						}
						phaseSettings := []int{a, b, c, d, e}
						s := findThrusterSignal(program, phaseSettings)
						if s > max {
							max = s
						}
					}
				}
			}
		}
	}
	return max
}

func findMaxThrusterSignalWithFeedback(program []int) int {
	max := 0
	for a := 5; a < 10; a++ {
		for b := 5; b < 10; b++ {
			for c := 5; c < 10; c++ {
				for d := 5; d < 10; d++ {
					for e := 5; e < 10; e++ {
						m := map[int]bool{a: true, b: true, c: true, d: true, e: true}
						if len(m) < 5 {
							continue
						}
						phaseSettings := []int{a, b, c, d, e}
						s := findThrusterSignalWithFeedback(program, phaseSettings)
						if s > max {
							max = s
						}
					}
				}
			}
		}
	}
	return max
}

func runIntcodeProgram(program []int, input []int) int {
	memory := make([]int, len(program))
	copy(memory, program)

	var output int
	ip := 0
	inputPtr := 0
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
			memory[parameter1] = input[inputPtr]
			inputPtr++
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

func runContinuousIntcodeProgram(program []int, in, out chan int) {
	memory := make([]int, len(program))
	copy(memory, program)

	ip := 0
	inputPtr := 0
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
			return

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
			input := <-in
			memory[parameter1] = input
			inputPtr++
			ip += 2

		case 4: //output
			output := getValue(1)
			out <- output
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
