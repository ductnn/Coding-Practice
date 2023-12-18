package main

import (
	"fmt"
	"sort"
)

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	return (nums[n-1] * nums[n-2]) - (nums[0] * nums[1])
}

func main() {
	nums := []int{4, 2, 5, 9, 7, 4, 8}
	fmt.Println(maxProductDifference(nums))
}
