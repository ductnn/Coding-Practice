// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero

package main

import (
	"fmt"
)

func minReorder(n int, connections [][]int) int {
	g := make([][][2]int, n)
	for _, e := range connections {
		a, b := e[0], e[1]
		g[a] = append(g[a], [2]int{b, 1})
		g[b] = append(g[b], [2]int{a, 0})
	}
	var dfs func(int, int) int
	dfs = func(a, fa int) (ans int) {
		for _, e := range g[a] {
			if b, c := e[0], e[1]; b != fa {
				ans += c + dfs(b, a)
			}
		}
		return
	}
	return dfs(0, -1)
}

func main() {
	n := 5
	connections := [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}

	fmt.Println(minReorder(n, connections))
}
