package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	temp1 := 0
	temp2 := 0

	for _, v := range nums {
		temp1 = ^temp2 & (temp1 ^ v)
		temp2 = ^temp1 & (temp2 ^ v)
	}

	return temp1
}

func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(nums))
}
