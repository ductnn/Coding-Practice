// https://leetcode.com/problems/missing-number/

package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	n := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return n*(n+1)/2 - sum
}

func missingNumber1(nums []int) int {
	n := len(nums)
	m := make([]bool, n+1)

	for _, v := range nums {
		m[v] = true
	}

	for i, exists := range m {
		if !exists {
			return i
		}
	}

	return 0
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber1(nums))
}
