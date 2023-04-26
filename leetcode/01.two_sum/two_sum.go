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

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
