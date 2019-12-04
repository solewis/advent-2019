package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", validPasswordCount(136760, 595730, validPassword))
	fmt.Printf("Part 2: %d\n", validPasswordCount(136760, 595730, validPasswordStrict))
}

func validPassword(pw int) bool {
	digits := toSlice(pw)
	if len(digits) < 6 || len(digits) > 6 {
		return false
	}
	if !allAscending(digits) {
		return false
	}

	for i := 0; i < 5; i++ {
		if digits[i] == digits[i+1] {
			return true
		}
	}
	return false
}

func validPasswordStrict(pw int) bool {
	digits := toSlice(pw)
	if len(digits) < 6 || len(digits) > 6 {
		return false
	}
	if !allAscending(digits) {
		return false
	}

	padded := []int{-1}
	padded = append(padded, digits...)
	padded = append(padded, -1)
	for i := 1; i < 6; i++ {
		if padded[i] == padded[i+1] {
			if padded[i-1] != padded[i] && padded[i+2] != padded[i+1] {
				return true
			}
		}
	}
	return false
}

func allAscending(digits []int) bool {
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] > digits[i+1] {
			return false
		}
	}
	return true
}

func toSlice(i int) []int {
	var digits []int
	for i > 0 {
		digits = append([]int{i % 10}, digits...)
		i /= 10
	}
	return digits
}

func validPasswordCount(min, max int, checker func(int) bool) int {
	count := 0
	for i := min; i <= max; i++ {
		if checker(i) {
			count++
		}
	}
	return count
}
