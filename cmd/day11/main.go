package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"math"
)

func main() {
	input := parse.Ints("cmd/day11/input.txt", ",")
	fmt.Printf("Part 1: %d\n", calcNumberOfPanelsPainted(input))

	paintRegistrationIdentifier(input)
}

func paintRegistrationIdentifier(program []int) {
	panels := buildPanelMap(program, 1)
	var minX, maxX, minY, maxY int
	for p := range panels {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			if panels[point{x, y}] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

type point struct{ x, y int }

type robot struct {
	pos point
	dir int
}

func (r *robot) turnLeft() {
	switch r.dir {
	case 'U':
		r.dir = 'L'
	case 'L':
		r.dir = 'D'
	case 'D':
		r.dir = 'R'
	case 'R':
		r.dir = 'U'
	}
}

func (r *robot) turnRight() {
	r.turnLeft()
	r.turnLeft()
	r.turnLeft()
}

func (r *robot) step() {
	switch r.dir {
	case 'U':
		r.pos.y++
	case 'L':
		r.pos.x--
	case 'D':
		r.pos.y--
	case 'R':
		r.pos.x++
	}
}

func buildPanelMap(program []int, startingColor int) map[point]int {
	panels := map[point]int{}

	getPanelColor := func(p point) int {
		if color, exists := panels[p]; exists {
			return color
		}
		return 0
	}

	in, out, done := make(chan int), make(chan int), make(chan int)

	go runIntcodeProgram(program, in, out, done)

	robot := robot{point{0, 0}, 'U'}
	input := startingColor
	for {
		select {
		case in <- input:
		case <-done:
			return panels
		}
		color, turn := <-out, <-out
		panels[robot.pos] = color
		switch turn {
		case 0:
			robot.turnLeft()
		case 1:
			robot.turnRight()
		}
		robot.step()
		input = getPanelColor(robot.pos)
	}
}

func calcNumberOfPanelsPainted(program []int) int {
	panels := map[point]int{}
	paintedPoints := map[point]int{}

	getPanelColor := func(p point) int {
		if color, exists := panels[p]; exists {
			return color
		}
		return 0
	}

	in, out, done := make(chan int), make(chan int), make(chan int)

	go runIntcodeProgram(program, in, out, done)

	robot := robot{point{0, 0}, 'U'}
	for {
		input := getPanelColor(robot.pos)
		select {
		case in <- input:
		case <-done:
			return len(paintedPoints)
		}
		color, turn := <-out, <-out
		panels[robot.pos] = color
		paintedPoints[robot.pos]++
		switch turn {
		case 0:
			robot.turnLeft()
		case 1:
			robot.turnRight()
		}
		robot.step()
	}
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
