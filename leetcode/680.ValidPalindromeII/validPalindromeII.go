package main

import (
	"fmt"
)

// Accepted
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

// Nope =(((
func validPalindrome2(s string) bool {
	// 2 pointers
	i, j := 0, len(s)-1

	// stack for temp
	var tempI, tempJ string

	for i < j {
		if s[i] != s[j] {
			tempI = s[i+1 : j+1]
			tempJ = s[i:j]
			return (tempI == reverse(tempI)) || (tempJ == reverse(tempJ))
		}
		i++
		j--
	}
	return true
}

func reverse(s string) string {
	result := []byte{}
	// for i, j := 0, len(s)-1; i < j; i++ {
	// 	result[i], result[j] = result[j], result[i]
	// 	j--
	// }
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return string(result)
}

// Accepted
func validPalindrome1(s string) bool {
	check := func(i, j int) bool {
		for ; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return check(i+1, j) || check(i, j-1)
		}
	}
	return true
}

func validPalindrome3(s string) bool {
	check := func(s string) bool {
		length := len(s)
		for i := 0; i < length/2; i++ {
			if s[i] != s[length-1-i] {
				return false
			}
		}
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return (check(s[i+1 : j+1])) || (check(s[i:j]))
		}
		i++
		j--
	}
	return true
}

func isPalindrome(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	s := "abcax"
	fmt.Println(validPalindrome3(s))
}
