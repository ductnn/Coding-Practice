package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[right] {
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] == nums[mid] {
				left++
			}
			if nums[right] == nums[mid] {
				right--
			}
		}
	}

	return -1
}

func main() {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	fmt.Println(search(input, target))
}
