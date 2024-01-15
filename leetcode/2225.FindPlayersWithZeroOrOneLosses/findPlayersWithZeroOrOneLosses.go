// https://leetcode.com/problems/find-players-with-zero-or-one-losses/

package main

import (
	"fmt"
	"sort"
)

func findWinners(matches [][]int) [][]int {
	cnt := map[int]int{}
	for _, m := range matches {
		a, b := m[0], m[1]
		if _, ok := cnt[a]; !ok {
			cnt[a] = 0
		}
		cnt[b]++
	}
	result := make([][]int, 2)
	for u, v := range cnt {
		if v < 2 {
			result[v] = append(result[v], u)
		}
	}
	sort.Ints(result[0])
	sort.Ints(result[1])
	return result
}

func main() {
	matches := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}

	fmt.Println(findWinners(matches))
}
