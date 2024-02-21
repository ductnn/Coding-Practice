// https://leetcode.com/problems/bitwise-and-of-numbers-range

package main

import (
	"fmt"
)

func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right &= (right - 1)
	}
	return right
}

func main() {
	left := 1
	right := 2147483647

	fmt.Println(rangeBitwiseAnd(left, right))
}
