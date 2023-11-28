package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	result := 0
	if len(nums) == 1 {
		return nums[0]
	}

	for _, v := range nums {
		result ^= v
	}

	return result
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
