package main

import (
	"fmt"
)

func maxArea(height []int) int {
	result := 0
	i, j := 0, len(height)-1
	for i < j {
		temp := (j - i) * min(height[i], height[j])
		result = max(result, temp)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
