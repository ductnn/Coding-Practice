package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	prefix, suffix := 1, 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}
	return result
}

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}
