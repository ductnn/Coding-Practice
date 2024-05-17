// https://leetcode.com/problems/binary-search

package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1

		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] == target {
		return left
	}

	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	fmt.Println(search(nums, target))
}
