package main

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	leftSum, rightSum := 0, 0
	for i := 0; i < len(nums); i++ {
		rightSum = sum - nums[i] - leftSum
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}

func pivotIndex1(nums []int) int {
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if prefix[i] == prefix[len(nums)-1]-prefix[i]+nums[i] {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}
