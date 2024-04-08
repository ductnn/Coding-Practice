// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch

package main

import (
	"fmt"
)

func countStudents(students []int, sandwiches []int) int {
	cnt := [2]int{}
	for _, v := range students {
		cnt[v]++
	}
	for _, v := range sandwiches {
		if cnt[v] == 0 {
			return cnt[v^1]
		}
		cnt[v]--
	}
	return 0
}

func main() {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}

	fmt.Println(countStudents(students, sandwiches))
}
