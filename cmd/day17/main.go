package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := parse.Ints("cmd/day17/input.txt", ",")
	fmt.Printf("Part 1: %d\n", sumOfAlignmentParameters(input))
	fmt.Printf("Part 2: %d\n", runVacuum(input))
}

func runVacuum(program []int) int {
	r, world := readMap(program)
	path := buildPath(r, world)
	fmt.Println(path)

	//start program again in movement mode
	program[0] = 2
	in, out, done := make(chan int, 100), make(chan int), make(chan int)
	go runIntcodeProgram(program, in, out, done)

	// figure out by hand...
	mvmtRoutine := "A,B,A,C,A,B,A,C,B,C"
	aFunc := "R,4,L,12,L,8,R,4"
	bFunc := "L,8,R,10,R,10,R,6"
	cFunc := "R,4,R,10,L,12"

	inputToRobot(mvmtRoutine, in)
	inputToRobot(aFunc, in)
	inputToRobot(bFunc, in)
	inputToRobot(cFunc, in)
	inputToRobot("n", in) //video feed
	return flushOutput(out) //flush a the full path that is output afterwards, returning the final value
}

func flushOutput(out chan int) int {
	var finalValue int
	for {
		select {
		case finalValue = <-out: // ignore initial output
		case <-time.After(1 * time.Second):
			return finalValue
		}
	}
}

func inputToRobot(s string, in chan int) {
	for _, r := range s {
		in <- int(r)
	}
	in <- int('\n')
}

func buildPath(r robot, world map[point]rune) string {
	steps := 0
	var path strings.Builder
	for {
		nextTurn := "S"
		next := r.next() // try straight
		if world[next] != '#' {
			nextTurn = "L"
			r.turnLeft() // try left
			next = r.next()
		}
		if world[next] != '#' {
			nextTurn = "R"
			r.turnLeft()
			r.turnLeft() // try right
			next = r.next()
		}
		if world[next] != '#' {
			_, _ = fmt.Fprintf(&path, "%s", strconv.Itoa(steps))
			return path.String() // reached end of path
		}
		if nextTurn != "S" {
			if steps != 0 {
				_, _ = fmt.Fprintf(&path, "%s,", strconv.Itoa(steps))
			}
			_, _ = fmt.Fprintf(&path, "%s,", nextTurn)
			steps = 0
			continue
		}
		r.step()
		steps++
	}
}

func sumOfAlignmentParameters(input []int) int {
	intersections := intersections(readMap(input))
	sum := 0
	for _, i := range intersections {
		alignmentParam := i.x * i.y
		sum += alignmentParam
	}
	return sum
}

type point struct{ x, y int }

type robot struct {
	pos point
	dir rune
}

func (r *robot) step() {
	r.pos = r.next()
}

func (r *robot) turnLeft() {
	switch r.dir {
	case '^':
		r.dir = '<'
	case 'v':
		r.dir = '>'
	case '>':
		r.dir = '^'
	case '<':
		r.dir = 'v'
	}
}

func (r robot) next() point {
	switch r.dir {
	case '^':
		return point{r.pos.x, r.pos.y - 1}
	case 'v':
		return point{r.pos.x, r.pos.y + 1}
	case '>':
		return point{r.pos.x + 1, r.pos.y}
	case '<':
		return point{r.pos.x - 1, r.pos.y}
	default:
		panic("invalid direction")
	}
}

func readMap(program []int) (robot, map[point]rune) {
	in, out, done := make(chan int), make(chan int), make(chan int)
	go runIntcodeProgram(program, in, out, done)

	world := map[point]rune{}
	currentPtr := point{0, 0}
	var r robot
	for {
		select {
		//case in <- input:
		case val := <-out:
			if rune(val) == '\n' {
				currentPtr = point{0, currentPtr.y + 1}
			} else if rune(val) == '>' || rune(val) == '<' || rune(val) == '^' || rune(val) == 'v' {
				r = robot{currentPtr, rune(val)}
				world[currentPtr] = '#'
				currentPtr = point{currentPtr.x + 1, currentPtr.y}
			} else {
				world[currentPtr] = rune(val)
				currentPtr = point{currentPtr.x + 1, currentPtr.y}
			}
		case <-done:
			//print(world)
			return r, world
		}
	}
}

func intersections(r robot, world map[point]rune) []point {
	visited := map[point]bool{r.pos: true}
	var ints []point
	for {
		next := r.next() // try straight
		if world[next] != '#' {
			r.turnLeft() // try left
			next = r.next()
		}
		if world[next] != '#' {
			r.turnLeft()
			r.turnLeft() // try right
			next = r.next()
		}
		if world[next] != '#' {
			break // reached end of path
		}
		r.step()
		if visited[r.pos] {
			ints = append(ints, r.pos)
		} else {
			visited[r.pos] = true
		}
	}
	return ints
}

func print(world map[point]rune) {
	var minX, minY, maxX, maxY int
	for p := range world {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.y < minY {
			minY = p.y
		}
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := point{x, y}
			fmt.Print(string(world[p]))
		}
		fmt.Println()
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
