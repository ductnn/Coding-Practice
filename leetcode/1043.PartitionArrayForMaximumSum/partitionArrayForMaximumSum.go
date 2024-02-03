// https://leetcode.com/problems/partition-array-for-maximum-sum

package main

import (
	"fmt"
)

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		mx := 0
		for j := i; j > max(0, i-k); j-- {
			mx = max(mx, arr[j-1])
			f[i] = max(f[i], f[j-1]+mx*(i-j+1))
		}
	}
	return f[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 15, 7, 9, 2, 5, 10}
	k := 3

	fmt.Println(maxSumAfterPartitioning(arr, k))
}
