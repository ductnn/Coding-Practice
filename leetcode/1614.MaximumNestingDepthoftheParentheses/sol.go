// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/

package main

import (
	"fmt"
)

func maxDepth(s string) int {
	res := 0
	d := 0
	for _, c := range s {
		if c == '(' {
			d++
			res = max(res, d)
		} else if c == ')' {
			d--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "(1+(2*3)+((8)/4))+1"
	fmt.Println(maxDepth(s))
}
