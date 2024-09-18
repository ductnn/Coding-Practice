// https://leetcode.com/problems/largest-number/description/

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Custom comparator function
func compare(x, y string) bool {
	return x+y > y+x
}

func largestNumber(nums []int) string {
	// Convert numbers to string
	numsStr := make([]string, len(nums))
	for i, num := range nums {
		numsStr[i] = strconv.Itoa(num)
	}
	// fmt.Print(numsStr)

	// Sort numbers using custom comparator
	sort.Slice(numsStr, func(i, j int) bool {
		return compare(numsStr[i], numsStr[j])
	})

	// fmt.Print(numsStr)

	// Join the sorted strings
	result := strings.Join(numsStr, "")

	// Handle edge case where the result is "0" (e.g., when input is [0, 0])
	if result[0] == '0' {
		return "0"
	}

	return result
}

func main() {
	// Example usage
	nums := []int{3, 30, 34, 5, 9}
	result := largestNumber(nums)
	fmt.Println("Largest number:", result)
}
