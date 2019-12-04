package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", validPasswordCount(136760, 595730, validPassword))
	fmt.Printf("Part 2: %d\n", validPasswordCount(136760, 595730, validPasswordStrict))
}

func validPassword(pw string) bool {
	if len(pw) < 6 || len(pw) > 6 {
		return false
	}
	if !allAscending(pw) {
		return false
	}

	runes := []rune(pw)
	for i := 0; i < 5; i++ {
		if runes[i] == runes[i+1] {
			return true
		}
	}
	return false
}

func validPasswordStrict(pw string) bool {
	if len(pw) < 6 || len(pw) > 6 {
		return false
	}
	if !allAscending(pw) {
		return false
	}

	padded := fmt.Sprintf("x%sx", pw)
	runes := []rune(padded)
	for i := 1; i < 6; i++ {
		if runes[i] == runes[i+1] {
			if runes[i-1] != runes[i] && runes[i+2] != runes[i+1] {
				return true
			}
		}
	}
	return false
}

func allAscending(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s)-1; i++ {
		if runes[i] > runes[i+1] {
			return false
		}
	}
	return true
}

func validPasswordCount(min, max int, checker func(string) bool) int {
	count := 0
	for i := min; i <= max; i++ {
		if checker(strconv.Itoa(i)) {
			count++
		}
	}
	return count
}
