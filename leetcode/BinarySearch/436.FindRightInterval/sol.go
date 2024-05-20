// https://leetcode.com/problems/find-right-interval

package main

import (
	"fmt"
	"sort"
)

// Interval struct
type Interval struct {
	Start int
	End   int
}

// FindRightInterval function
func FindRightInterval(intervals []Interval) []int {
	n := len(intervals)
	result := make([]int, n)
	if n == 0 {
		return result
	}

	// Create a sorted array of start intervals with original indices
	startIntervals := make([][2]int, n)
	for i := 0; i < n; i++ {
		startIntervals[i] = [2]int{intervals[i].Start, i}
	}
	sort.Slice(startIntervals, func(i, j int) bool {
		return startIntervals[i][0] < startIntervals[j][0]
	})

	for i := 0; i < n; i++ {
		result[i] = -1
		end := intervals[i].End
		// Binary search to find the smallest start interval
		left, right := 0, n
		for left < right {
			mid := left + (right-left)>>1
			if startIntervals[mid][0] >= end {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if left < n {
			result[i] = startIntervals[left][1]
		}
	}

	return result
}

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	starts := make([][]int, n)
	for i := 0; i < n; i++ {
		starts[i] = make([]int, 2)
		starts[i][0] = intervals[i][0]
		starts[i][1] = i
	}
	sort.Slice(starts, func(i, j int) bool {
		return starts[i][0] < starts[j][0]
	})
	var res []int
	for _, interval := range intervals {
		left, right, end := 0, n-1, interval[1]
		for left < right {
			mid := left + (right-left)>>1
			if starts[mid][0] >= end {
				right = mid
			} else {
				left = mid + 1
			}
		}
		val := -1
		if starts[left][0] >= end {
			val = starts[left][1]
		}
		res = append(res, val)
	}
	return res
}

func main() {
	intervals := []Interval{
		{3, 4},
		{2, 3},
		{1, 2},
	}
	intervals2 := [][]int{
		{3, 4},
		{2, 3},
		{1, 2},
	}

	fmt.Println(FindRightInterval(intervals))
	fmt.Println(findRightInterval(intervals2))
}
