// https://leetcode.com/problems/number-of-good-pairs/

package main

import (
	"fmt"
)

func numIdenticalPairs(nums []int) int {
	result := 0
	count := [101]int{}

	for _, num := range nums {
		result += count[num]
		count[num]++
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}

	fmt.Println(numIdenticalPairs(nums))
}
