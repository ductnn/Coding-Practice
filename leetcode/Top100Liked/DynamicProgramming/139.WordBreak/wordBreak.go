// https://leetcode.com/problems/word-break

package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	words := map[string]bool{}

	for _, w := range wordDict {
		words[w] = true
	}

	f := make([]bool, n+1)
	f[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if f[j] && words[s[j:i]] {
				f[i] = true
				break
			}
		}
	}
	return f[n]
}

func main() {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}

	fmt.Println(wordBreak(s, wordDict))
}
