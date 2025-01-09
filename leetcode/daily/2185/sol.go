// https://leetcode.com/problems/counting-words-with-a-given-prefix/description/

package main

import "fmt"

func prefixCount(words []string, pref string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		if isPrefix(pref, words[i]) {
			count++
		}
	}

	return count
}

func isPrefix(prefix, word string) bool {
	n := len(prefix)
	m := len(word)

	prefixMatch := m >= n && word[:n] == prefix

	return prefixMatch
}

func main() {
	words := []string{"pay", "attention", "practice", "attend"}
	prefix := "at"
	fmt.Println(prefixCount(words, prefix))
}
