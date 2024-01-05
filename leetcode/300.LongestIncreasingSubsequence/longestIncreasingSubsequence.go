// https://leetcode.com/problems/longest-increasing-subsequence/

package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	result := 1
	n := len(nums)
	f := make([]int, n+1)

	for i := range f {
		f[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j]+1)
				result = max(result, f[i])
			}
		}
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
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println(lengthOfLIS(nums))
}
