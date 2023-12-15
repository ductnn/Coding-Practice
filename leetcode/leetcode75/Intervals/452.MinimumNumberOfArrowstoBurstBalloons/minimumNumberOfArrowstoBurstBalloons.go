package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	result := 0
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	last := -99999999999999
	for _, p := range points {
		a, b := p[0], p[1]
		if a > last {
			result++
			last = b
		}
	}

	return result
}

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points))
}
