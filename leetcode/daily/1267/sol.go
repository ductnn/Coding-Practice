// https://leetcode.com/problems/count-servers-that-communicate/description/

package main

func countServers(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	rows := make([]int, m)
	cols := make([]int, n)

	// count servers in each row and column
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}

	// count servers that can communicate
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (rows[i] > 1 || cols[j] > 1) {
				count++
			}
		}
	}

	return count
}

func main() {
	grid := [][]int{
		{1, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}

	println(countServers(grid))
}
