package main

import (
	"fmt"
)

func getLengthOfOptimalCompression(s string, k int) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = n
		}
	}

	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= min(i, k); j++ {
			if j > 0 {
				dp[i][j] = dp[i-1][j-1]
			}
			same, diff := 0, 0
			for m := i; m >= 1 && diff <= j; m-- {
				if s[m-1] == s[i-1] {
					same++
					dp[i][j] = min(dp[i][j], dp[m-1][j-diff]+calLength(same))
				} else {
					diff++
				}
				// if j+diff <= k {
				// 	dp[m][j+diff] = min(dp[m][j+diff], calLength(same)+1+dp[i-1][j])
				// } else {
				// 	break
				// }
			}
		}
	}

	return dp[n][k]
}

func calLength(x int) int {
	if x == 1 {
		return 1
	}
	if x < 10 {
		return 2
	}
	if x < 100 {
		return 3
	}
	return 4
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "aaabcccd"
	k := 2
	fmt.Println(getLengthOfOptimalCompression(s, k)) // Output: 4
}
