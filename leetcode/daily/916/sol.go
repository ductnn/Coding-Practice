// https://leetcode.com/problems/word-subsets/description/

package main

import "fmt"

func wordSubsets(words1 []string, words2 []string) []string {
	// count freq
	countFrequencies := func(word string) [26]int {
		freq := [26]int{}
		for _, c := range word {
			freq[c-'a']++
		}
		return freq
	}

	// combine max words2
	maxFreq := [26]int{}
	for _, word := range words2 {
		freq := countFrequencies(word)
		for i := 0; i < 26; i++ {
			maxFreq[i] = max(maxFreq[i], freq[i])
		}
	}

	result := []string{}
	for _, word := range words1 {
		freq := countFrequencies(word)
		isUniversal := true
		for i := 0; i < 26; i++ {
			if freq[i] < maxFreq[i] {
				isUniversal = false
				break
			}
		}
		if isUniversal {
			result = append(result, word)
		}
	}

	return result
}

func main() {
	words1 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	words2 := []string{"e", "o"}
	fmt.Println(wordSubsets(words1, words2))
}
