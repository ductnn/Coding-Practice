// https://leetcode.com/problems/uncommon-words-from-two-sentences/description/

package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	words := strings.Fields(s1 + " " + s2)

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	var res []string
	for word, count := range wordCount {
		if count == 1 {
			res = append(res, word)
		}
	}

	return res
}

func main() {
	s1 := "this apple is sweet"
	s2 := "this apple is sour"
	result := uncommonFromSentences(s1, s2)
	fmt.Println(result)
}
