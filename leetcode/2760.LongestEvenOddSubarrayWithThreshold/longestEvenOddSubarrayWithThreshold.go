package main

import (
	"fmt"
)

func longestAlternatingSubarray(nums []int, threshold int) int {
	result := 0
	for l, n := 0, len(nums); l < n; {
		if nums[l]%2 == 0 && nums[l] <= threshold {
			r := l + 1
			for r < n && nums[r]%2 != nums[r-1]%2 && nums[r] <= threshold {
				r++
			}
			result = max(result, r-l)
			l = r
		} else {
			l++
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

func main() {
	nums := []int{3, 2, 5, 4}
	threshold := 5

	fmt.Println(longestAlternatingSubarray(nums, threshold))
}
