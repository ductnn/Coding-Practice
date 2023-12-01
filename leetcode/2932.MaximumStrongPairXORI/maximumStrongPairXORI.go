package main

import (
	"fmt"
	"sort"
)

func maximumStrongPairXor(nums []int) int {
	//
	// strong pair: |x - y| <= min(x, y)
	// max xor: max(result, x^y)
	//
	sort.Ints(nums)
	result := 0
	for _, x := range nums {
		for _, y := range nums {
			if abs(x, y) <= min(x, y) {
				result = max(result, x^y)
			}
		}
	}

	return result
}

func maximumStrongPairXor1(nums []int) int {
	//
	// strong pair: |x - y| <= min(x, y)
	// max xor: max(result, x^y)
	//
	sort.Ints(nums)

	result := 0
	i, j := 0, 0

	for i < len(nums) {
		if abs(nums[i], nums[j]) <= min(nums[i], nums[j]) {
			result = max(result, nums[i]^nums[j])
		}
		j++

		if j == len(nums) {
			i++
			j = i
		}
	}

	return result
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(maximumStrongPairXor1(nums))
}
