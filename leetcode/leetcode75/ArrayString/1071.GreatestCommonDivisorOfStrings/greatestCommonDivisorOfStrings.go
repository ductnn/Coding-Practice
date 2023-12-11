package main

import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	n := gcd(len(str1), len(str2))
	return str1[:n]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	str1 := "ABCABC"
	str2 := "ABC"

	fmt.Println(gcd(len(str1), len(str2)))
	fmt.Println(gcdOfStrings(str1, str2))
}
