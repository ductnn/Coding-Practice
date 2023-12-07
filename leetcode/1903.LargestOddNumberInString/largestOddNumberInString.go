package main

import (
	"fmt"
)

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		c := num[i] - '0'
		if (c & 1) == 1 {
			return num[:i+1]
		}
	}
	return ""
}

func largestOddNumber1(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if (num[i]-'0')%2 != 0 {
			return num[:i+1]
		}
	}
	return ""
}

func main() {
	num := "354274"
	fmt.Println(largestOddNumber1(num))
}
