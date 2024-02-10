// https://leetcode.com/problems/palindromic-substrings

package main

import (
	"fmt"
)

func countSubstrings(s string) int {
	res := 0
	n := len(s)

	for k := 0; k < n*2-1; k++ {
		i, j := k/2, (k+1)/2

		for i >= 0 && j < n && s[i] == s[j] {
			res++
			i, j = i-1, j+1
		}
	}

	return res
}

func main() {
	s := "aaa"

	fmt.Println(countSubstrings(s))
}
