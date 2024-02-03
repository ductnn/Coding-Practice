// https://leetcode.com/problems/sequential-digits

package main

import (
	"fmt"
	"sort"
)

func sequentialDigits(low int, high int) []int {
	ans := []int{}
	for i := 1; i < 9; i++ {
		x := i
		for j := i + 1; j < 10; j++ {
			x = x*10 + j
			if low <= x && x <= high {
				ans = append(ans, x)
			}
		}
	}
	sort.Ints(ans)
	return ans
}

func main() {
	low := 100
	high := 300

	fmt.Println(sequentialDigits(low, high))
}
