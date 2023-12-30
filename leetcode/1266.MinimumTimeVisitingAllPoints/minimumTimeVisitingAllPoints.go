// https://leetcode.com/problems/minimum-time-visiting-all-points

package main

import (
	"fmt"
)

func minTimeToVisitAllPoints(points [][]int) int {
	result := 0
	for i, p := range points[1:] {
		dx := abs(p[0] - points[i][0])
		dy := abs(p[1] - points[i][1])
		result += max(dx, dy)
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	points := [][]int{{1, 1}, {3, 4}, {-1, 0}}

	fmt.Println(minTimeToVisitAllPoints(points))
}
