// https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram

package main

import (
	"fmt"
)

func minSteps(s string, t string) int {
	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}

	result := 0
	for i := 0; i < 26; i++ {
		if cnt[i] > 0 {
			result += cnt[i]
		}
	}

	return result
}

func main() {
	s := "leetcode"
	t := "practice"

	fmt.Println(minSteps(s, t))
}
