// https://leetcode.com/problems/fibonacci-number

package main

import (
	"fmt"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func main() {
	n := 10
	fmt.Println(fib(n))
}
