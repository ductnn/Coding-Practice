// https://leetcode.com/problems/find-the-town-judge

package main

import (
	"fmt"
)

func findJudge(n int, trust [][]int) int {
	if n == 1 && len(trust) == 0 {
		return 1
	}

	m := make(map[int]int)

	for _, t := range trust {
		m[t[0]]--
		m[t[1]]++
	}

	for i := 1; i <= n; i++ {
		if m[i] == n-1 {
			return i
		}
	}

	return -1
}

func main() {
	n := 4
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}

	fmt.Println(findJudge(n, trust))
}
