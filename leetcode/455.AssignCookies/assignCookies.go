// https://leetcode.com/problems/assign-cookies

package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	j := 0
	for i, v := range g {
		for j < len(s) && s[j] < v {
			j++
		}
		if j >= len(s) {
			return i
		}
		j++
	}

	return len(g)
}

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}

	fmt.Println(findContentChildren(g, s))
}
