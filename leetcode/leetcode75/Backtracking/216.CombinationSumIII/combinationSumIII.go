package main

import (
	"fmt"
)

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	t := []int{}
	var dfs func(i, s int)
	dfs = func(i, s int) {
		if s == 0 {
			if len(t) == k {
				cp := make([]int, len(t))
				copy(cp, t)
				result = append(result, cp)
			}
			return
		}
		if i > 9 || i > s || len(t) >= k {
			return
		}
		t = append(t, i)
		dfs(i+1, s-i)
		t = t[:len(t)-1]
		dfs(i+1, s)
	}
	dfs(1, n)
	return result
}

func main() {
	k := 3
	n := 7
	fmt.Println(combinationSum3(k, n))
}
