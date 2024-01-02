// https://leetcode.com/problems/convert-an-array-into-a-2d-array-with-conditions/

package main

import (
	"fmt"
)

func findMatrix(nums []int) [][]int {
	result := make([][]int, 0)

	n := len(nums)
	cnt := make([]int, n+1)
	for _, v := range nums {
		cnt[v]++
	}

	for i, v := range cnt {
		for j := 0; j < v; j++ {
			if len(result) <= j {
				result = append(result, make([]int, 0))
			}
			result[j] = append(result[j], i)
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, 4, 1, 2, 3, 1}

	fmt.Println(findMatrix(nums))
}
