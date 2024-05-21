// https://leetcode.com/problems/subsets/description/

package main

import "fmt"

func subsets(nums []int) [][]int {
	n := len(nums)
	totalSubsets := 1 << n // 2^n subsets
	result := make([][]int, 0, totalSubsets)

	for subsetMask := 0; subsetMask < totalSubsets; subsetMask++ {
		subset := []int{}
		for i := 0; i < n; i++ {
			// Check if the i-th bit in subsetMask is set
			if subsetMask&(1<<i) != 0 {
				subset = append(subset, nums[i])
			}
		}
		result = append(result, subset)
	}

	return result
}

func main() {
	nums := []int{1, 2, 3}
	subsets := subsets(nums)
	for _, subset := range subsets {
		fmt.Println(subset)
	}
}
