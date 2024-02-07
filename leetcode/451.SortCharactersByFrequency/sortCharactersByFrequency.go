// https://leetcode.com/problems/sort-characters-by-frequency

package main

import (
	"bytes"
	"fmt"
	"sort"
)

func frequencySort(s string) string {

	// count character
	cnt := map[byte]int{}
	for i := range s {
		cnt[s[i]]++
	}

	// sort character
	cs := make([]byte, 0, len(s))
	for c := range cnt {
		cs = append(cs, c)
		fmt.Println(cs)
	}
	sort.Slice(cs, func(i, j int) bool { return cnt[cs[i]] > cnt[cs[j]] })
	fmt.Println(cs)

	for _, s := range cs {
		tmp := sortString(string(s))
		fmt.Println(tmp)
	}

	// append to res
	res := make([]byte, 0, len(s))
	for _, c := range cs {
		res = append(res, bytes.Repeat([]byte{c}, cnt[c])...)
	}
	return string(res)
}

func sortString(s string) string {
	bytes := []byte(s)
	tmp := make([]int, len(bytes))
	for i, b := range bytes {
		tmp[i] = int(b)
	}
	sort.Ints(tmp)
	for i, v := range tmp {
		bytes[i] = byte(v)
	}
	return string(bytes)
}

func main() {
	s := "aaabccc"
	fmt.Println(frequencySort(s))
}
