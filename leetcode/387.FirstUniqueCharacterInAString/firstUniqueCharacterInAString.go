// https://leetcode.com/problems/first-unique-character-in-a-string

package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	cnt := make([]int, 26)

	for _, c := range s {
		cnt[c-'a']++
		fmt.Println(cnt[c-'a'])
	}

	for i, c := range s {
		if cnt[c-'a'] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}
