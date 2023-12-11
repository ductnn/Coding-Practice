package main

import (
	"fmt"
)

func largestAltitude(gain []int) int {
	result := 0
	sum := 0

	for i := 0; i < len(gain); i++ {
		sum += gain[i]
		result = max(result, sum)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(largestAltitude(gain))
}
