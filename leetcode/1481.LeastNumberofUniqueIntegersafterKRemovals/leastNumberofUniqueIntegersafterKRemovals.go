// https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/

package main

import (
	"fmt"
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	cnt := make(map[int]int)
	for _, v := range arr {
		cnt[v]++
	}

	n := len(cnt)
	nums := make([]int, 0, n)
	for _, v := range cnt {
		nums = append(nums, v)
	}
	sort.Ints(nums)

	for i, v := range nums {
		k -= v
		if k < 0 {
			return len(nums) - i
		}
	}

	return 0
}

func main() {
	arr := []int{4, 3, 1, 1, 3, 3, 2}
	k := 3

	fmt.Println(findLeastNumOfUniqueInts(arr, k))
}
