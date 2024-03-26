// https://leetcode.com/problems/first-missing-positive

package main

import "fmt"

func firstMissingPositive(nums []int) int {
	m := make(map[int]bool)

	for _, num := range nums {
		if num > 0 {
			m[num] = true
		}
	}

	for i := 1; ; i++ {
		if !m[i] {
			return i
		}
	}
}

func main() {
	nums := []int{3, 4, -1, 1}

	fmt.Println(firstMissingPositive(nums))
}
