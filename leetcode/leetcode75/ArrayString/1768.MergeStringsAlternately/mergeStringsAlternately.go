package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	m, n := len(word1), len(word2)
	result := make([]byte, 0, m+n)

	for i := 0; i < m || i < n; i++ {
		if i < m {
			result = append(result, word1[i])
		}
		if i < n {
			result = append(result, word2[i])
		}
	}

	return string(result)
}

func main() {
	word1 := "abc"
	word2 := "pqr"

	fmt.Println(mergeAlternately(word1, word2))
}
