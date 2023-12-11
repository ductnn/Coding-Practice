package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	result := []string{}
	w := strings.Split(s, " ")
	for i := len(w) - 1; i >= 0; i-- {
		if w[i] != "" {
			result = append(result, w[i])
		}
	}
	return strings.Join(result, " ")
}

func main() {
	s := "   the sky is blue    "
	fmt.Println(reverseWords(s))
}
