// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent

package main

import (
	"fmt"
	"strings"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func arrayStringsAreEqual1(word1 []string, word2 []string) bool {
	w1, w2 := "", ""

	for _, w := range word1 {
		w1 += w
	}

	for _, w := range word2 {
		w2 += w
	}

	return w1 == w2
}

func main() {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}

	fmt.Println(arrayStringsAreEqual1(word1, word2))
}
