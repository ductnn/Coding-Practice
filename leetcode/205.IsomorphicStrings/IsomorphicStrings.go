package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make([]int, 256)
	m2 := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}

		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}

	return true
}

func main() {
	s := "asdd"
	t := "qwed"

	fmt.Println(isIsomorphic(s, t))
}
