package main

import (
	"fmt"
)

func countCharacters(words []string, chars string) int {
	result := 0
	countChar := [26]int{}

	for _, char := range chars {
		countChar[char-'a']++
	}

	for _, word := range words {
		countWord := [26]int{}
		for _, w := range word {
			countWord[w-'a']++
		}

		ok := true
		for i := 0; i < len(countWord); i++ {
			if countChar[i] < countWord[i] {
				ok = false
				break
			}
		}

		if ok {
			result += len(word)
		}
	}

	return result
}

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"

	countChar := [26]int{}

	for _, char := range chars {
		countChar[char-'a']++
		fmt.Println(countChar)
	}

	fmt.Println(countCharacters(words, chars))
}
