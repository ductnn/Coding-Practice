package main

import (
	"fmt"
)

func moveZeroes(nums []int) []int {
	i := -1
	for k, v := range nums {
		if v != 0 {
			i++
			Swap(nums, i, k)
		}
	}
	return nums
}

func Swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroes(nums))
}
