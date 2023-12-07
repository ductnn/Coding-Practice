package main

import (
	"fmt"
)

func moveZeroes(nums []int) []int {
	i := -1
	for k, v := range nums {
		if v != 0 {
			i++
			nums[i], nums[k] = nums[k], nums[i]
		}
	}
	return nums
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroes(nums))
}
