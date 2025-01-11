// https://leetcode.com/problems/construct-k-palindrome-strings/description/

package main

import "fmt"

func canConstruct(s string, k int) bool {
	// count freq
	freq := make(map[rune]int)
	for _, c := range s {
		freq[c]++
	}

	// count odd freq
	odd := 0
	for _, v := range freq {
		if v%2 != 0 {
			odd++
		}
	}

	// check conditions
	return k <= len(s) && odd <= k
}

func canConstruct2(s string, k int) bool {
	if len(s) < k {
		return false
	}

	// count freq
	freq := make([]int, 26)
	for _, c := range s {
		freq[c-'a']++
	}

	// count odd freq
	odd := 0
	for _, v := range freq {
		if v%2 == 1 {
			odd++
		}
	}

	if odd > k {
		return false
	}

	return true
}

func main() {
	s := "annabelle"
	k := 2

	fmt.Println(canConstruct2(s, k))
}
