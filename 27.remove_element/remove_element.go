package main

import (
	"fmt"
)

// 0  1  2  2  3  0  4  2
// i
// j

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	j := 0

	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	return i
}

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	fmt.Println(removeElement(arr, val))
}
