package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"math"
)

func main() {
	input := parse.Ints("cmd/day19/input.txt", ",")
	fmt.Printf("Part 1: %d\n", pointsAffectedByBeam(input))
	p := edgeOfSquare(input)
	fmt.Printf("Part 2: %d\n", p.x*10000+p.y)
}

func edgeOfSquare(program []int) point {
	y := 10 //start at 10 to skip the empty rows
	for {
		x := 0
		firstBeam := 0
	column:
		for {
			p := point{x, y}
			if firstBeam == 0 {
				if beingPulled(p, program) {
					firstBeam = x
				}
			} else {
				if x-firstBeam == 99 { // 100 steps into current beam, check 100 steps up to see if that was also in the beam
					p2 := point{x, y - 99}
					if beingPulled(p2, program) {
						return point{x - 99, y - 99}
					}
				}
				if !beingPulled(p, program) { // end of beam for the given y
					break column
				}
			}
			x++
		}
		y++
	}
}

type point struct{ x, y int }

func pointsAffectedByBeam(program []int) int {
	beamMap := map[point]bool{}
	count := 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			p := point{x, y}
			beamMap[p] = beingPulled(p, program)
			if beamMap[p] {
				count++
			}
		}
	}
	printBeam(beamMap)
	return count
}

func printBeam(beamMap map[point]bool) {
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			p := point{x, y}
			if beamMap[p] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func beingPulled(p point, program []int) bool {
	in, out, done := make(chan int), make(chan int), make(chan int)
	go runIntcodeProgram(program, in, out, done)
	in <- p.x
	in <- p.y
	o := <-out
	<-done
	return o == 1
}

func runIntcodeProgram(program []int, in, out, done chan int) {
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
			dest := getFromMemory(ip + paramNum)
			destMode := getFromMemory(ip) / int(math.Pow10(paramNum+1)) % 10
			if destMode == 2 {
				dest += relativeBase
			}
			return dest
		}

		getValue := func(paramNum int) int {
			param := getFromMemory(ip + paramNum)
			mode := getFromMemory(ip) / int(math.Pow10(paramNum+1)) % 10
			if mode == 0 { //position mode
				return getFromMemory(param)
			} else if mode == 1 { //immediate mode
				return param
			}
			return getFromMemory(param + relativeBase) //relative mode
		}

		switch opCode {

		case 99: //break
			done <- 1
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
			input := <-in
			memory[dest] = input
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
