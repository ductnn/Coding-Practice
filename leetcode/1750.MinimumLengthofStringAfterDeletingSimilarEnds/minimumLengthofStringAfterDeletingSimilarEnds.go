// https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends

package main

import (
	"fmt"
)

func minimumLength(s string) int {
	i, j := 0, len(s)-1

	for i < j && s[i] == s[j] {
		for i+1 < j && s[i] == s[i+1] {
			i++
		}
		for i < j-1 && s[j] == s[j-1] {
			j--
		}
		i, j = i+1, j-1
	}

	return max(0, j-i+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "aabccabba"

	fmt.Println(minimumLength(s))
}
