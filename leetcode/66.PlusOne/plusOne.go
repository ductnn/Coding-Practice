package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	result := []int{}
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}

	result = append([]int{1}, digits...)

	return result
}

func plusOne1(digits []int) []int {
	result := []int{}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}

	result = append([]int{1}, digits...)

	return result
}

func main() {
	digits := []int{9}
	fmt.Println(plusOne1(digits))
}
