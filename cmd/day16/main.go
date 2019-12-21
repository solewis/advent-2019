package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"strconv"
	"strings"
)

func main() {
	input := parse.Lines("cmd/day16/input.txt")[0]
	fmt.Printf("Part 1: %s\n", runPhases(input, 100)[:8])
	fmt.Printf("Part 2: %s\n", p2(input))
}

func p2(signal string) string {
	offset, _ := strconv.Atoi(signal[:7])

	var fullSignal strings.Builder
	fullSignal.Grow(10000 * len(signal))
	for i := 0; i < 10000; i++ {
		_, _ = fmt.Fprintf(&fullSignal, "%s", signal)
	}

	//build slice only off the part of the string starting at the offset
	s := buildSlice(fullSignal.String()[offset:])
	for phase := 0; phase < 100; phase++ {
		nextSignal := make([]int, len(s))
		// for all numbers in the final 1/4 of the slice:
		//  the multiplier will be all 0 and 1, because of the repeating, so...
		//	last digit is always equal to previous last digit (because all the other values will be multiply by 0)
		//	second to last digit is always sum of previous second to last digit and last digit (taking only the 1's value of that sum)
		//  third to last digit is sum of previous third to last, second to last, and last (taking only the 1's value of that sum)
		//  so on...
		sum := 0
		for i := len(s) - 1; i >= 0; i-- {
			sum += s[i]
			nextSignal[i] = sum % 10
		}

		s = nextSignal
	}
	val := ""
	for i := 0; i < 8; i++ {
		val += strconv.Itoa(s[i])
	}
	return val
}

func runPhases(signal string, phases int) string {
	signalSlice := buildSlice(signal)
	pattern := []int{0, 1, 0, -1}

	for phase := 0; phase < phases; phase++ {
		nextSignal := make([]int, len(signalSlice))

		//calculate each output position in the next signal one by one
		for position := 0; position < len(signalSlice); position++ {
			positionValue := 0
			for i := 0; i < len(signalSlice); i++ {
				positionValue += signalSlice[i] * getMultiplier(pattern, position, i)
			}
			nextSignal[position] = abs(positionValue) % 10
		}
		signalSlice = nextSignal
	}
	valuesText := make([]string, len(signalSlice))
	for i := range signalSlice {
		text := strconv.Itoa(signalSlice[i])
		valuesText[i] = text
	}
	return strings.Join(valuesText, "")
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func buildSlice(signal string) []int {
	signalSlice := make([]int, len(signal))
	for i, b := range []rune(signal) {
		signalSlice[i] = int(b - '0')
	}
	return signalSlice
}

func getMultiplier(basePattern []int, calculatePosition, signalPosition int) int {
	signalPosition += 1
	calculatePosition += 1
	return basePattern[signalPosition/calculatePosition%len(basePattern)]
}
