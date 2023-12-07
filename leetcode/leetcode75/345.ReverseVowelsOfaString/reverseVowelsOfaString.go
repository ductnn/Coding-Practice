package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	str := []byte(s)
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !strings.Contains(vowels, string(str[left])) {
			left++
		}
		for left < right && !strings.Contains(vowels, string(str[right])) {
			right--
		}
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
	return string(str)
}

func main() {
	s := "leetcode"
	fmt.Println(reverseVowels(s))
}
