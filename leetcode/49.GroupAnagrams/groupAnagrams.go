// https://leetcode.com/problems/group-anagrams

package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	m := map[string][]string{}

	for _, s := range strs {
		tmp := sortString(s)
		m[tmp] = append(m[tmp], s)
		fmt.Println(m[tmp])
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
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

type Key [26]byte

func strKey(str string) (key Key) {
	for i := range str {
		key[str[i]-'a']++
	}
	return
}

func groupAnagrams1(strs []string) [][]string {
	m := make(map[Key][]string)
	for _, v := range strs {
		key := strKey(v)
		m[key] = append(m[key], v)
	}
	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(strs))
}
