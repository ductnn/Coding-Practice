package main

import (
	"fmt"
)

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[l] && nums[mid] == nums[r] {
			r--
			l++
		} else {
			if nums[l] <= nums[mid] {
				if nums[l] <= target && nums[mid] > target {
					r = mid - 1
				} else {
					l = mid + 1
				}
			} else {
				if nums[mid] < target && nums[r] >= target {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
		}
	}

	return false
}

func QuickSort(array []int) []int {
	// Write your code here.
	helper(array, 0, len(array)-1)

	return array
}

func helper(array []int, left, right int) {
	if len(array) <= 1 {
		return
	}

	if left < right {
		pivot := Partition(array, left, right)
		helper(array, left, pivot-1)
		helper(array, pivot+1, right)
	}
}

func Partition(array []int, left, right int) int {
	index := left - 1
	// Defind pivot element
	p := array[right]

	for i := left; i < right; i++ {
		if array[i] <= p {
			index++
			Swap(array, index, i)
		}
	}

	Swap(array, index+1, right)

	return index + 1
}

func Swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	input := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}
	target := 2

	fmt.Println(search(input, target))
}
