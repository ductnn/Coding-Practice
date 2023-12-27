// https://leetcode.com/problems/minimum-time-to-make-rope-colorful/

package main

import (
	"fmt"
)

func minCost(colors string, neededTime []int) int {
	n := len(colors)
	prev := -1
	minTime := 0

	for i := 0; i < n; i++ {
		if prev == -1 || colors[prev] != colors[i] {
			prev = i
		} else {
			if neededTime[prev] < neededTime[i] {
				minTime += neededTime[prev]
				prev = i
			} else {
				minTime += neededTime[i]
			}
		}
	}

	return minTime
}

func minCost1(colors string, neededTime []int) int {
	n := len(colors)
	minTime := 0
	for i, j := 0, 0; i < n; i = j {
		j = i
		s, mx := 0, 0
		for j < n && colors[j] == colors[i] {
			s += neededTime[j]
			if mx < neededTime[j] {
				mx = neededTime[j]
			}
			j++
		}
		if j-i > 1 {
			minTime += s - mx
		}
	}
	return minTime
}

func main() {
	colors := "abaac"
	neededTime := []int{1, 2, 3, 4, 5}

	fmt.Println(minCost1(colors, neededTime))
}
