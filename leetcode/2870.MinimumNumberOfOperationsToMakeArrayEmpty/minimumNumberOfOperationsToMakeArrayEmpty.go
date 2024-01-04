// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-empty

package main

import (
	"fmt"
)

func minOperations(nums []int) int {
	result := 0
	count := map[int]int{}

	for _, num := range nums {
		count[num]++
	}

	for _, c := range count {
		if c < 2 {
			return -1
		}
		result += (c + 2) / 3
	}

	return result
}

func main() {
	nums := []int{2, 3, 3, 2, 2, 4, 2, 3, 4}

	fmt.Println(minOperations(nums))
}
