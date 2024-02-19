// https://leetcode.com/problems/power-of-two

package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	for n%2 == 0 {
		n /= 2
	}

	return n == 1
}

func main() {
	n := 16

	fmt.Print(isPowerOfTwo(n))
}
