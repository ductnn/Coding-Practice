// https://leetcode.com/problems/score-of-a-string/description/

package main

import (
	"fmt"
)

func scoreOfString(s string) int {
	res := 0
	for i := 1; i < len(s); i++ {
		res += abs(int(s[i-1]) - int(s[i]))
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	s := "zaz"
	fmt.Println(scoreOfString(s))
}
