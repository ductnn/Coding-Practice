package main

import (
	"fmt"
)

func longestSubarray(nums []int) int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == 1 {
			left[i] = left[i-1] + 1
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i+1] == 1 {
			right[i] = right[i+1] + 1
		}
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		result = max(result, left[i]+right[i])
	}
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	fmt.Println(longestSubarray(nums))
}
