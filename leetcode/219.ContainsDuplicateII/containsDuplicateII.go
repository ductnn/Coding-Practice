package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	if (len(nums) <= 1) || (k <= 0) {
		return false
	}

	m := map[int]int{}
	for i, n := range nums {
		if j, found := m[n]; found && i-j <= k {
			return true
		}
		m[n] = i
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}
