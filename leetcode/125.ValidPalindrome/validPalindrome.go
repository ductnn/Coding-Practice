package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	// 2 pointer
	i, j := 0, len(s)-1
	for i < j {
		if !isAlnum(s[i]) {
			i++
		} else if !isAlnum(s[j]) {
			j--
		} else if toLower(s[i]) != toLower(s[j]) {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

func isAlnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || ch >= '0' && ch <= '9'
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}

	return ch
}

func main() {
	s := "A man, a plan, a canal: Panama X"
	fmt.Println(isPalindrome(s))
}
