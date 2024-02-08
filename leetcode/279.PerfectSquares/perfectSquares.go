// https://leetcode.com/problems/perfect-squares

package main

import (
	"math"
)

func numSquares(n int) int {
	f := make([]int, n+1)

	for i := 1; i <= n; i++ {
		minCount := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minCount = min(minCount, f[i-j*j]+1)
		}
		f[i] = minCount
	}

	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example usage:
	n := 12
	result := numSquares(n)
	println(result) // Output: 2 (12 = 4 + 4 + 4)
}
