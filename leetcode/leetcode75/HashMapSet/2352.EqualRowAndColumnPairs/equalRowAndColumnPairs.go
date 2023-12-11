package main

import (
	"fmt"
)

func equalPairs(grid [][]int) int {
	result := 0
	for i := range grid {
		for j := range grid {
			ok := 1
			for k := range grid {
				if grid[i][k] != grid[k][j] {
					ok = 0
					break
				}
			}
			result += ok
		}
	}
	return result
}

func main() {
	grid := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	fmt.Print(equalPairs(grid))
}
