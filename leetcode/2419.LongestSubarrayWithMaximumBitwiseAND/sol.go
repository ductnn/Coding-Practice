// https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/description/

package main

import (
	"fmt"
)

func longestSubarray(nums []int) int {
	// Step 1: Find the maximum value in the array
	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}

	// Step 2: Find the length of the longest subarray where all elements equal maxVal
	longest := 0
	currentLength := 0

	for _, num := range nums {
		if num == maxVal {
			currentLength++
			if currentLength > longest {
				longest = currentLength
			}
		} else {
			currentLength = 0
		}
	}

	return longest
}

func main() {
	nums := []int{1, 2, 3, 3, 2, 2, 3}

	fmt.Println(longestSubarray(nums))
}
