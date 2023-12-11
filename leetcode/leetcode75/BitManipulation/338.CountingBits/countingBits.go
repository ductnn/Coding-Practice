package main

import (
	"fmt"
)

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}

func main() {
	n := 5
	fmt.Println(countBits(n))
}
