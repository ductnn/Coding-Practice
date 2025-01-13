// https://leetcode.com/problems/minimum-length-of-string-after-operations/description/

package main

import "fmt"

func minimumLength(s string) int {
	freq := [26]int{}
	for _, v := range s {
		freq[v-'a']++
	}

	res := 0
	for _, v := range freq {
		if v > 0 {
			if v&1 == 1 {
				res += 1
			} else {
				res += 2
			}
		}
	}

	return res
}

func main() {
	s := "aa"
	fmt.Println(minimumLength(s))
}
