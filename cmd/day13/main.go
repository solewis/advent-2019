package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"math"
)

func main() {
	input := parse.Ints("cmd/day13/input.txt", ",")
	fmt.Printf("Part 1: %d\n", calcNumberOfBlockTiles(input))
	fmt.Printf("Part 2: %d\n", play(input))
}

func play(program []int) int {
	program[0] = 2

	in, out, done := make(chan int), make(chan int), make(chan int)
	go runIntcodeProgram(program, in, out, done)

	var ball, paddle point
	var input, score int
	for {
		//move the paddle with the ball
		if ball.x > paddle.x {
			input = 1
		} else if ball.x < paddle.x {
			input = -1
		} else {
			input = 0
		}
		select {
		case in <- input:
		case x := <-out:
			y := <-out
			tileId := <-out
			//keep track of paddle, ball, and score
			if tileId == 3 {
				paddle = point{x, y}
			}
			if tileId == 4 {
				ball = point{x, y}
			}
			if x == -1 && y == 0 {
				score = tileId
			}
		case <-done:
			return score
		}
	}
}

// Used initially to figure out what the game was
func printGame(panels map[point]int) {
	var maxX, maxY int
	for p := range panels {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	fmt.Print("  ")
	for x := 0; x <= maxX; x++ {
		fmt.Print(x)
	}
	fmt.Println()
	for y := 0; y <= maxY; y++ {
		fmt.Printf("%d ", y)
		for x := 0; x <= maxX; x++ {
			p := point{x, y}
			if panels[p] == 0 {
				fmt.Print(".")
			} else if panels[p] == 1 {
				fmt.Print("W")
			} else if panels[p] == 2 {
				fmt.Print("B")
			} else if panels[p] == 3 {
				fmt.Print("P")
			} else if panels[p] == 4 {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type point struct{ x, y int }

func calcNumberOfBlockTiles(program []int) int {
	in, out, done := make(chan int), make(chan int), make(chan int)
	go runIntcodeProgram(program, in, out, done)

	cnt := 0
	for {
		select {
		case <-out:
			<-out
			tileId := <-out
			if tileId == 2 {
				cnt++
			}
		case <-done:
			return cnt
		}
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
