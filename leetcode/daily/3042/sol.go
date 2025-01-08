// https://leetcode.com/problems/count-prefix-and-suffix-pairs-i/

package main

import "fmt"

func countPrefixSuffixPairs(words []string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				count++
			}
		}
	}

	return count
}

func isPrefixAndSuffix(prefixSuffix, word string) bool {
	n := len(prefixSuffix)
	m := len(word)

	prefixMatch := m >= n && word[:n] == prefixSuffix
	suffixMatch := m >= n && word[m-n:] == prefixSuffix

	return prefixMatch && suffixMatch
}

func main() {
	words := []string{"a", "aba", "ababa", "aa"}
	fmt.Println(countPrefixSuffixPairs(words))
}
