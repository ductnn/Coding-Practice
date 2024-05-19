// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix

package main

import (
	"fmt"
)

func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i, j := m-1, 0; i >= 0 && j < n; {
		if grid[i][j] < 0 {
			res += n - j
			i--
		} else {
			j++
		}
	}
	return res
}

func main() {
	grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	fmt.Println(countNegatives(grid))
}
