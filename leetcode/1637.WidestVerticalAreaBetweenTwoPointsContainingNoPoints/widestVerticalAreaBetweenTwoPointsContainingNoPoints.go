// https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/

package main

import (
	"fmt"
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })

	result := 0
	for i, p := range points[1:] {
		result = max(result, p[0]-points[i][0])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	points := [][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}
	fmt.Println(maxWidthOfVerticalArea(points))
}
