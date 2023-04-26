package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	short := shortestString(strs)

	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}

	return short
}

func shortestString(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(res) > len(strs) {
			res = strs[i]
		}
	}

	return res
}

func main() {
	input := []string{"fl", "flcow", "flcight"}

	fmt.Println(longestCommonPrefix(input))
}
