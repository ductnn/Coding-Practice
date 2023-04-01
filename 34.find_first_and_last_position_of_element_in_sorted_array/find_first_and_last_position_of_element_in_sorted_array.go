package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	return []int{searchFirstElement(nums, target), searchLastElement(nums, target)}
}

func searchFirstElement(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + ((right - left) / 2)

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if (mid == 0) || (nums[mid-1] != target) {
				return mid
			}
			right = mid - 1
		}
	}

	return -1
}

func searchLastElement(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + ((right - left) / 2)
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) {
				return mid
			}
			left = mid + 1
		}
	}

	return -1
}

func main() {
	input := []int{5, 7, 7, 8, 8, 10}
	target := 8

	fmt.Println(searchRange(input, target))
}
