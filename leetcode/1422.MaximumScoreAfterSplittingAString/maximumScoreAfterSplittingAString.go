// https://leetcode.com/problems/maximum-score-after-splitting-a-string

package main

import (
	"fmt"
)

func maxScore(s string) int {
	temp := 0
	if s[0] == '0' {
		temp++
	}
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i] == '1' {
			temp++
		}
	}
	result := temp
	for i := 1; i < n-1; i++ {
		if s[i] == '0' {
			temp++
		} else {
			temp--
		}
		result = max(result, temp)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxScore1(s string) int {
	result := 0
	n := len(s)
	for i := range s {
		if i == n-1 {
			break
		}
		left := s[:i+1]
		right := s[i+1:]

		scoreLeft := countScore(left, '0')
		scoreRight := countScore(right, '1')

		sum := scoreLeft + scoreRight
		if sum > result {
			result = sum
		}
	}

	return result
}

func countScore(s string, compare int) int {
	count := 0
	for _, v := range s {
		if v == rune(compare) {
			count++
		}
	}
	return count
}

func main() {
	s := "011101"
	fmt.Println(maxScore1(s))
}
