// https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/

package main

import (
	"fmt"
)

const (
	INF = 9999999999
)

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, d+1)
		for j := range f[i] {
			f[i][j] = INF
		}
	}

	f[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, d); j++ {
			mx := 0
			for k := i; k > 0; k-- {
				mx = max(mx, jobDifficulty[k-1])
				f[i][j] = min(f[i][j], f[k-1][j-1]+mx)
			}
		}
	}

	if f[n][d] == INF {
		return -1
	}

	return f[n][d]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	jobDifficulty := []int{6, 5, 4, 3, 2, 1}
	d := 2

	fmt.Println(minDifficulty(jobDifficulty, d))
}
