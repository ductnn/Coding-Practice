// https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/

package main

import (
	"fmt"
)

func minOperations(s string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			if s[i] == '0' {
				cnt++
			}
		} else {
			if s[i] == '1' {
				cnt++
			}
		}
	}
	return min(cnt, len(s)-cnt)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minOperations1(s string) int {
	cnt := 0
	for i, c := range s {
		if c != []rune("01")[i&1] {
			cnt++
		}
	}
	return min(cnt, len(s)-cnt)
}

func main() {
	s := "1111"

	fmt.Println(minOperations1(s))
}
