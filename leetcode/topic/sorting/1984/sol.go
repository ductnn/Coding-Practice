// https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}

	sort.Ints(nums)

	minDiff := math.MaxInt32

	for i := 0; i <= len(nums)-k; i++ {
		diff := nums[i+k-1] - nums[i]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

func main() {
	nums := []int{9, 4, 1, 7}
	k := 2
	result := minimumDifference(nums, k)
	fmt.Println(result)
}
