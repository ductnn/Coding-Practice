// https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/

package main

import (
	"fmt"
)

func findTheLongestSubstring(s string) int {
	// Map to store the first occurrence of each bitmask
	vowelMap := map[int]int{0: -1}
	bitmask := 0
	maxLength := 0

	// Mapping vowels to bit positions
	vowelToBit := map[rune]int{
		'a': 0,
		'e': 1,
		'i': 2,
		'o': 3,
		'u': 4,
	}

	for i, char := range s {
		// If the character is a vowel, toggle its corresponding bit
		if bit, exists := vowelToBit[char]; exists {
			bitmask ^= (1 << bit)
		}

		// Check if this bitmask has been seen before
		if pos, exists := vowelMap[bitmask]; exists {
			maxLength = max(maxLength, i-pos)
		} else {
			vowelMap[bitmask] = i
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	s := "eleetminicoworoep"

	fmt.Println(findTheLongestSubstring(s))
}
