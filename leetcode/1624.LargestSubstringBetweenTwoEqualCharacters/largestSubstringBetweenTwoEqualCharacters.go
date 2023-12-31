// https://leetcode.com/problems/largest-substring-between-two-equal-characters/

package main

import (
	"fmt"
)

func maxLengthBetweenEqualCharacters(s string) int {
	charIdx := make([]int, 26)
	for i := range charIdx {
		charIdx[i] = -1
	}

	result := -1

	for i, c := range s {
		if charIdx[c-'a'] == -1 {
			charIdx[c-'a'] = i
		} else {
			result = max(result, i-charIdx[c-'a']-1)
		}
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
	s := "cbzxy"

	fmt.Println(maxLengthBetweenEqualCharacters(s))
}
