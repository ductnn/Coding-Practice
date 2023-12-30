// https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/

package main

import (
	"fmt"
)

func makeEqual(words []string) bool {
	n := len(words)
	count := make([]int, 26)

	for _, word := range words {
		for _, w := range word {
			count[w-'a']++
		}
	}

	for i := 0; i < 26; i++ {
		if count[i]%n != 0 {
			return false
		}
	}

	return true
}

func main() {
	words := []string{"abc", "aabc", "bc"}

	fmt.Println(makeEqual(words))
}
