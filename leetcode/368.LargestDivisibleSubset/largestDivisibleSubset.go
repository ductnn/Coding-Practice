// https://leetcode.com/problems/largest-divisible-subset

package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	res := []int{}
	sort.Ints(nums)

	n := len(nums)
	f := make([]int, n)
	k := 0

	for i := 0; i < n; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				f[i] = max(f[i], f[j]+1)
			}
		}
		if f[k] < f[i] {
			k = i
		}
	}

	m := f[k]
	for i := k; m > 0; i-- {
		if nums[k]%nums[i] == 0 && f[i] == m {
			res = append(res, nums[i])
			k = i
			m--
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3}

	fmt.Println(largestDivisibleSubset(nums))
}
