package main

import (
	"fmt"
)

func countGoodSubstrings(s string) int {
	// s = "xyzzaz"
	count := 0

	for i := 0; i < len(s)-2; i++ {
		if s[i] != s[i+1] && s[i+1] != s[i+2] && s[i] != s[i+2] {
			count++
		}
	}

	return count
}

func countGoodSubstrings1(s string) int {
	count := 0

	for i := 0; i < len(s)-2; i++ {
		substring := s[i : i+3]
		if isDistinct(substring) {
			count++
		}
	}

	return count
}

func isDistinct(s string) bool {
	seen := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if seen[s[i]] {
			return false
		}
		seen[s[i]] = true
	}
	return true
}

func main() {
	s := "xyzzaz"
	fmt.Println(countGoodSubstrings1(s))
}
