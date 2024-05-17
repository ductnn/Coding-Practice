// https://leetcode.com/problems/running-sum-of-1d-array/description/

package main

import (
	"fmt"
)

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return nums
}

func main() {
	nums := []int{1, 1, 1, 1, 1}

	fmt.Println(runningSum(nums))
}
