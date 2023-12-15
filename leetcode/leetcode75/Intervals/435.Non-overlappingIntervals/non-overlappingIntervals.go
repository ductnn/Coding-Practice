package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	temp, cnt := intervals[0][1], 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < temp {
			cnt++
		} else {
			temp = intervals[i][1]
		}
	}
	return cnt
}

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(intervals))
}
