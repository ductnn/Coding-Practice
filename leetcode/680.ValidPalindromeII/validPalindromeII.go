package main

import (
	"fmt"
)

func validPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return check(i+1, j, s) || check(i, j-1, s)
		}
	}
	return true
}

func check(i, j int, s string) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		} else {
			i++
			j--
		}
	}

	return true
}

// func validPalindrome(s string) bool {
//     check := func(i, j int) bool {
// 		for ; i < j; i, j = i+1, j-1 {
// 			if s[i] != s[j] {
// 				return false
// 			}
// 		}
// 		return true
// 	}
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		if s[i] != s[j] {
// 			return check(i+1, j) || check(i, j-1)
// 		}
// 	}
// 	return true
// }

func main() {
	s := "abca"
	fmt.Println(validPalindrome(s))
}
