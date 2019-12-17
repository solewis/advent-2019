package main

import (
	"container/list"
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"math"
	"time"
)

func main() {
	program := parse.Ints("cmd/day15/input.txt", ",")
	fmt.Printf("Part 1: %d\n", distanceToOxygenSystem(program))
	fmt.Printf("Part 2: %d\n", minutesToOxygenate(program))
}

func minutesToOxygenate(program []int) int {
	paths, start := mapSystem(program)
	currentLayer := []point{start}
	oxygenated := map[point]bool{start: true}
	minutes := 1
	// start at the system, add all of its neighbors as the next layer to get oxygenated
	// then visit all of those neighbors and add their neighbors as the next layer to get oxygenated.
	// continue until all open space is oxygenated, each layer taking one minute
	for {
		var nextLayer []point
		for _, p := range currentLayer {
			for _, n := range p.neighbors() {
				if !oxygenated[n] && paths[n] {
					nextLayer = append(nextLayer, n)
					oxygenated[n] = true
				}
			}
		}
		if len(oxygenated) == len(paths) {
			return minutes
		}
		currentLayer = nextLayer
		minutes++
	}
}

type point struct{ x, y int }

func (p point) neighbors() []point {
	return []point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y + 1},
		{p.x, p.y - 1},
	}
}

func mapSystem(program []int) (map[point]bool, point) {
	in, out, done := make(chan int), make(chan int), make(chan int)
	go runIntcodeProgram(program, in, out, done)

	toVisit := list.New()
	visited := map[point]point{} // map of node, with previous node as value for path finding
	walls := map[point]bool{}

	toVisit.PushBack(point{0, 0})
	currentLocation := point{0, 0}
	var oxygenSystem point
	visited[currentLocation] = currentLocation
	for {
		front := toVisit.Front()
		if front == nil {
			break
		}
		next := front.Value.(point)
		toVisit.Remove(front)
		move(currentLocation, next, visited, in, out)
		currentLocation = next

		for _, n := range next.neighbors() {
			if _, visited := visited[n]; visited || walls[n] {
				continue
			}

			neighborResult := step(currentLocation, n, in, out) //check the neighbor
			if neighborResult == 0 {                            //neighbor is wall, position unchanged
				walls[n] = true
			} else if neighborResult == 1 || neighborResult == 2 { //neighbor is open, moved to neighbor
				step(n, currentLocation, in, out) //return back
				visited[n] = currentLocation      //add path to neighbor
				toVisit.PushBack(n)
				if neighborResult == 2 {
					oxygenSystem = n
				}
			}
		}
	}
	paths := map[point]bool{}
	for p := range visited {
		paths[p] = true
	}
	//print(currentLocation, visited, walls)
	return paths, oxygenSystem
}

func distanceToOxygenSystem(program []int) int {
	in, out, done := make(chan int), make(chan int), make(chan int)
	go runIntcodeProgram(program, in, out, done)

	toVisit := list.New()
	visited := map[point]point{} // map of node, with previous node as value for path finding
	walls := map[point]bool{}

	toVisit.PushBack(point{0, 0})
	currentLocation := point{0, 0}
	visited[currentLocation] = currentLocation
	for {
		front := toVisit.Front()
		next := front.Value.(point)
		toVisit.Remove(front)
		move(currentLocation, next, visited, in, out)
		currentLocation = next

		for _, n := range next.neighbors() {
			if _, visited := visited[n]; visited || walls[n] {
				continue
			}

			neighborResult := step(currentLocation, n, in, out) //check the neighbor
			if neighborResult == 0 {                            //neighbor is wall, position unchanged
				walls[n] = true
			} else if neighborResult == 1 { //neighbor is open, moved to neighbor
				step(n, currentLocation, in, out) //return back
				visited[n] = currentLocation      //add path to neighbor
				toVisit.PushBack(n)
			} else { //oxygen system!
				steps := 1
				start := point{0, 0}
				ptr := currentLocation
				for {
					if ptr == start {
						return steps
					}
					ptr = visited[ptr]
					steps++
				}

			}
		}
		//print(currentLocation, visited, walls)
	}
}

func print(droid point, open map[point]point, walls map[point]bool) {
	var minX, minY, maxX, maxY int
	for p := range open {
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
	for p := range walls {
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
			if p == droid {
				fmt.Print("D")
			} else if _, o := open[p]; o {
				fmt.Print(".")
			} else if _, w := walls[p]; w {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	time.Sleep(100 * time.Millisecond)
}

func move(from, to point, paths map[point]point, in, out chan int) {
	toStart(from, paths, in, out)
	fromStart(to, paths, in, out)
}

//move from 0,0 to target, along a known path
//paths contains a map of each point and its previous point in the shortest path back to 0,0
func fromStart(to point, paths map[point]point, in, out chan int) {
	start := point{0, 0}
	if to == start {
		return
	}
	prev := paths[to]
	fromStart(prev, paths, in, out) //recurse to end first, since all paths work towards the start
	step(prev, to, in, out)
}

//move to 0,0 from given point along a known path
//paths contains a map of each point and its previous point in the shortest path back to 0,0
func toStart(from point, paths map[point]point, in, out chan int) {
	target := point{0, 0}
	if from == target {
		return
	}
	next := paths[from]
	step(from, next, in, out)
	toStart(next, paths, in, out)
}

//take one step from point to point, to a known open area
//assumes the two points are neighbors
func step(from, to point, in, out chan int) int {
	if from.x < to.x {
		return check(4, in, out) //east
	}
	if from.x > to.x {
		return check(3, in, out) //west
	}
	if from.y < to.y {
		return check(1, in, out) //north
	}
	if from.y > to.y {
		return check(2, in, out) //south
	}
	panic("you passed in the same two points")
}

func check(dir int, in, out chan int) int {
	in <- dir
	return <-out
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
