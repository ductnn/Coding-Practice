//
// Author: Duc Tran
// Title: Two Sum
//

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	results := []int{}
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				results = []int{i, j}
			}
		}
	}
	return results
}

func twoSum1(nums []int, target int) []int {
	m := map[int]int{}

	for i := range nums {
		temp := nums[i]
		search := target - temp

		if j, ok := m[search]; ok {
			return []int{i, j}
		}

		m[temp] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum1(nums, target))
}
