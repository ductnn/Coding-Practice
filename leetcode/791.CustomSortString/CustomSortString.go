// https://leetcode.com/problems/custom-sort-string

package main

import (
	"fmt"
)

func customSortString(order string, s string) string {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}

	res := []rune{}
	for _, c := range order {
		for cnt[c-'a'] > 0 {
			res = append(res, c)
			cnt[c-'a']--
		}
	}
	for i, v := range cnt {
		for j := 0; j < v; j++ {
			res = append(res, rune('a'+i))
		}
	}

	return string(res)
}

func main() {
	order := "bcafg"
	s := "abcd"

	fmt.Println(customSortString(order, s))
}
