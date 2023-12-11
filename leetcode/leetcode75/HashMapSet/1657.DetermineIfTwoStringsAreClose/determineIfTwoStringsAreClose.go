package main

import (
	"fmt"
	"sort"

	"slices"
)

func closeStrings(word1 string, word2 string) bool {
	map1 := make([]int, 26)
	map2 := make([]int, 26)
	for _, c := range word1 {
		map1[c-'a']++
	}
	for _, c := range word2 {
		map2[c-'a']++
	}
	if !slices.EqualFunc(map1, map2, func(v1, v2 int) bool { return (v1 == 0) == (v2 == 0) }) {
		return false
	}
	sort.Ints(map1)
	sort.Ints(map2)
	return slices.Equal(map1, map2)
}

func main() {
	word1 := "uau"
	word2 := "ssx"

	fmt.Println(closeStrings(word1, word2))
}
