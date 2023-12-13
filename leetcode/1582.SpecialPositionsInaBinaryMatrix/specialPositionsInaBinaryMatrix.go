package main

import (
	"fmt"
)

func numSpecial(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	row, col := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			row[i] += mat[i][j]
			col[j] += mat[i][j]
		}
	}

	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if row[i] == 1 && col[j] == 1 && mat[i][j] == 1 {
				result++
			}
		}
	}

	return result
}

func main() {
	mat := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(numSpecial(mat))
}
