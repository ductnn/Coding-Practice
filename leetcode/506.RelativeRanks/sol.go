// https://leetcode.com/problems/relative-ranks

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	// Create a map to store the score-index pairs
	scoreMap := make(map[int]int)
	for i, score := range nums {
		scoreMap[score] = i
	}

	// Sort the scores in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	// Assign medals to the top three scores
	medals := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}

	// Assign rankings
	rankings := make([]string, len(nums))
	for i, score := range nums {
		if i < 3 {
			rankings[scoreMap[score]] = medals[i]
		} else {
			rankings[scoreMap[score]] = strconv.Itoa(i + 1)
		}
	}

	return rankings
}

func main() {
	score := []int{5, 4, 3, 2, 1}
	fmt.Println(findRelativeRanks(score))
}
