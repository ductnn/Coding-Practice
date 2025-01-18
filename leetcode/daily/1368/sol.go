// https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/description/

package main

import (
	"container/list"
	"fmt"
)

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// define directions
	dirs := [4][2]int{
		{0, 1},  // right
		{0, -1}, // left
		{1, 0},  // down
		{-1, 0}, // up
	}

	// init cost
	costs := make([][]int, m)
	for i := range costs {
		costs[i] = make([]int, n)
		for j := range costs[i] {
			costs[i][j] = 1 << 30
		}
	}
	costs[0][0] = 0 // start point cost is 0

	// define dqueue
	dq := list.New()
	dq.PushBack([3]int{0, 0, 0}) // [x, y, cost]

	// bfs
	for dq.Len() > 0 {
		node := dq.Remove(dq.Front()).([3]int)
		x, y, cost := node[0], node[1], node[2]

		// if reach the end point, return the cost
		if x == m-1 && y == n-1 {
			return cost
		}

		// try all directions
		for i, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n { // Nếu ô mới nằm trong lưới
				newCost := cost
				if grid[x][y] != i+1 { // Nếu hướng sai, tăng chi phí
					newCost++
				}

				// Nếu tìm được đường đi với chi phí thấp hơn
				if newCost < costs[nx][ny] {
					costs[nx][ny] = newCost
					if grid[x][y] == i+1 {
						dq.PushFront([3]int{nx, ny, newCost}) // Ưu tiên chi phí 0
					} else {
						dq.PushBack([3]int{nx, ny, newCost}) // Chi phí 1 thêm vào sau
					}
				}
			}
		}
	}

	return -1
}

func main() {
	grid1 := [][]int{
		{1, 1, 3},
		{3, 2, 2},
		{1, 1, 4},
	}
	fmt.Println(minCost(grid1)) // Output: 0

	grid2 := [][]int{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{1, 1, 1, 1},
		{2, 2, 2, 2},
	}
	fmt.Println(minCost(grid2)) // Output: 3
}
