// https://leetcode.com/problems/find-polygon-with-the-largest-perimeter

package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	n := len(nums)
	s := make([]int, n+1)

	for i, v := range nums {
		s[i+1] = s[i] + v
	}

	res := -1
	for k := 3; k <= n; k++ {
		if s[k-1] > nums[k-1] {
			res = max(res, s[k])
		}
	}

	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 12, 1, 2, 5, 50, 3}

	fmt.Println(largestPerimeter(nums))
}
