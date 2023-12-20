// https://leetcode.com/problems/number-of-provinces

package main

import (
	"fmt"
)

func findCircleNum(isConnected [][]int) int {
	result := 0
	n := len(isConnected)
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(i int) {
		vis[i] = true
		for k, v := range isConnected[i] {
			if !vis[k] && v == 1 {
				dfs(k)
			}
		}
	}
	for i, v := range vis {
		if !v {
			result++
			dfs(i)
		}
	}
	return result
}

func main() {
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(isConnected))
}
