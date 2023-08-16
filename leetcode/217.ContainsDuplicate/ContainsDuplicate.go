package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	sort.Ints(nums)

	disct := make(map[int]bool, len(nums))
	for _, number := range nums {
		if disct[number] {
			return true
		} else {
			disct[number] = true
		}
	}

	return false
}

func main() {
	nums := []int{8, 5, 2, 9, 10, 6, 3, 3}

	fmt.Println(containsDuplicate(nums))
}
