package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	min, mid := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num > mid {
			return true
		}
		if num <= min {
			min = num
		} else {
			mid = num
		}
	}
	return false
}

func increasingTriplet1(nums []int) bool {
	min := nums[0]
	var mid, i int
	for i = 1; i < len(nums); i++ {
		if nums[i] > min {
			mid = nums[i]
			i++
			break
		}
		min = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] <= min {
			min = nums[i]
		} else if nums[i] <= mid {
			mid = nums[i]
		} else {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{2, 1, 5, 0, 4, 6}
	fmt.Println(increasingTriplet1(nums))
}
