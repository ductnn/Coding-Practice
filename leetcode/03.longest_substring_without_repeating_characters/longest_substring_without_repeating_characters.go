package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLength := 0

	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	for i := range s {
		visited := false
		for j := i + 1; j < len(s); j++ {
			for z := i; z < j; z++ {
				if s[j] == s[z] {
					visited = true
					break
				}
			}

			if visited == true {
				if maxLength < j-i {
					maxLength = j - i
				}
				break
			}

			if j == len(s)-1 {
				if maxLength < j+1-i {
					maxLength = j + 1 - i
				}
			}
		}
	}

	return maxLength
}

func main() {
	s := "ductnnnnnnn"
	fmt.Println(lengthOfLongestSubstring(s))
}
