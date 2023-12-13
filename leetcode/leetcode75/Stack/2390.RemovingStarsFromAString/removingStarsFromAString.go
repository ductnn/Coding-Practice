package main

import (
	"fmt"
)

func removeStars(s string) string {
	result := []rune{}
	for _, c := range s {
		if c == '*' {
			result = result[:len(result)-1]
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}

func main() {
	s := "leet**cod*e"
	fmt.Println(removeStars(s))
}
