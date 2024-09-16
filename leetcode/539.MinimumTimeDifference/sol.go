// https://leetcode.com/problems/minimum-time-difference/

package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	// Convert time to minute
	timeToMinutes := func(time string) int {
		parts := strings.Split(time, ":")
		hours, _ := strconv.Atoi(parts[0])
		minutes, _ := strconv.Atoi(parts[1])

		return hours*60 + minutes
	}

	// Sort list minute
	minuteList := make([]int, len(timePoints))
	for i, time := range timePoints {
		minuteList[i] = timeToMinutes(time)
	}
	sort.Ints(minuteList)

	// Caculate the difference
	minDiff := math.MaxInt32
	for i := 1; i < len(minuteList); i++ {
		minDiff = min(minDiff, minuteList[i]-minuteList[i-1])
	}

	// Find the min
	minDiff = min(minDiff, 24*60-minuteList[len(minuteList)-1]+minuteList[0])

	return minDiff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	timePoints := []string{"23:59", "00:00"}
	result := findMinDifference(timePoints)
	fmt.Println(result)
}
