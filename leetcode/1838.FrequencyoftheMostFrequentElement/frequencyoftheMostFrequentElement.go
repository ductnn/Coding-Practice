package main

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	// Sliding window
	result, window := 1, 0
	for l, r := 0, 1; r < len(nums); r++ {
		window += (nums[r] - nums[r-1]) * (r - l)
		for window > k {
			window -= nums[r] - nums[l]
			l++
		}

		result = max(result, r-l+1)
	}

	return result
}

func maxFrequency1(nums []int, k int) int {
	sort.Ints(nums)

	// Sliding window
	l, curr := 0, 0
	for r := 0; r < len(nums); r++ {
		target := nums[r]
		curr += target

		if (r-l+1)*target-curr > k {
			curr -= nums[l]
			l++
		}
	}

	return len(nums) - l
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	nums := []int{1, 4, 8, 13}
	k := 5

	fmt.Println(maxFrequency1(nums, k))
}
