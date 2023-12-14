package main

import (
	"fmt"
)

func onesMinusZeros(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])

	rows := make([]int, m)
	cols := make([]int, n)

	diff := make([][]int, m)

	for i, row := range grid {
		diff[i] = make([]int, n)
		for j, v := range row {
			rows[i] += v
			cols[j] += v
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diff[i][j] = rows[i] + cols[j] - (n - rows[i]) - (m - cols[j])
		}
	}

	return diff
}

func onesMinusZeros1(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	oneRows, oneCols := make([]int, m), make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			oneRows[i] += grid[i][j] + grid[i][j]
			oneCols[j] += grid[i][j] + grid[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = oneRows[i] + oneCols[j] - m - n
		}
	}

	return grid
}

func main() {
	grid := [][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}
	fmt.Println(onesMinusZeros(grid))
}
