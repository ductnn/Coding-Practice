package main

import (
	"fmt"
)

func maxVowels(s string, k int) int {
	count := 0
	for i := 0; i < k; i++ {
		if isVowels(s[i]) {
			count++
		}
	}
	result := count
	for i := k; i < len(s); i++ {
		if isVowels(s[i]) {
			count++
		}
		if isVowels(s[i-k]) {
			count--
		}
		result = max(result, count)
	}
	return result
}

func maxVowels1(s string, k int) int {
	count := 0
	for i := 0; i < k; i++ {
		if isVowels1(s[i]) != 0 {
			count++
		}
	}
	result := count
	for i := k; i < len(s); i++ {
		count = count - isVowels1(s[i-k]) + isVowels1(s[i])
		if count > result {
			result = count
		}
	}
	return result
}

func isVowels1(c byte) int {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return 1
	}
	return 0
}

func isVowels(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	s := "abciiidef"
	k := 3
	fmt.Println(maxVowels1(s, k))
}
