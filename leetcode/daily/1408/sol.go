package main

import (
	"fmt"
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	// Sort words by length
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	var result []string

	// Check each word if it's a substring of any other word
	for i, word := range words {
		for _, other := range words[i+1:] {
			if strings.Contains(other, word) {
				result = append(result, word)
				break
			}
		}
	}

	return result
}

func main() {
	words := []string{"mass", "as", "hero", "superhero"}
	fmt.Println(stringMatching(words))
}
