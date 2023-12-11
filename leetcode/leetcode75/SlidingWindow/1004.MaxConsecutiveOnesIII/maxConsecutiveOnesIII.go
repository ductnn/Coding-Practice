package main

import (
	"fmt"
)

func longestOnes(nums []int, k int) int {
	result := 0
	j, count := 0, 0
	for i, v := range nums {
		if v == 0 {
			count++
		}
		for count > k {
			if nums[j] == 0 {
				count--
			}
			j++
		}
		result = max(result, i-j+1)
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
	nums := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k := 3
	fmt.Println(longestOnes(nums, k))
}
