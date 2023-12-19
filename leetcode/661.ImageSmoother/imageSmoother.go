package main

import (
	"fmt"
)

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	result := make([][]int, m)
	for i, row := range img {
		result[i] = make([]int, n)
		for j := range row {
			s, cnt := 0, 0
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x >= 0 && x < m && y >= 0 && y < n {
						cnt++
						s += img[x][y]
					}
				}
			}
			result[i][j] = s / cnt
		}
	}
	return result
}

func main() {
	img := [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}
	fmt.Println(imageSmoother(img))
}
