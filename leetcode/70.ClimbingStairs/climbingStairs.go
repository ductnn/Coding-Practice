package main

import (
	"fmt"
)

func climbStairs(n int) int {
	// Fibonaci
	if n == 1 || n == 2 {
		return n
	}

	a, b := 0, 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return b
}

func main() {
	n := 3
	fmt.Println(climbStairs(n))
}
