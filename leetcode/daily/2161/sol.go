// https://leetcode.com/problems/partition-array-according-to-given-pivot/description/

package main

import "fmt"

func pivotArray(nums []int, pivot int) []int {
	left := []int{}
	equal := []int{}
	right := []int{}

	for _, num := range nums {
		if num < pivot {
			left = append(left, num)
		} else if num == pivot {
			equal = append(equal, num)
		} else {
			right = append(right, num)
		}
	}

	return append(left, append(equal, right...)...)
}

func main() {
	nums := []int{9, 12, 5, 10, 14, 3, 10}
	pivot := 10

	fmt.Print(pivotArray(nums, pivot))
}
