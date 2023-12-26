// https://leetcode.com/problems/pascals-triangle/

package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	f := [][]int{{1}}
	for i := 0; i < numRows-1; i++ {
		g := []int{1}
		for j := 0; j < len(f[i])-1; j++ {
			g = append(g, f[i][j]+f[i][j+1])
		}
		g = append(g, 1)
		f = append(f, g)
	}
	return f
}

func main() {
	numRows := 5

	fmt.Println(generate(numRows))
}
