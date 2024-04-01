package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}

	j := i
	for j >= 0 && s[j] != ' ' {
		j--
	}

	return i - j
}

func lengthOfLastWord1(s string) int {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")

	return len(words[len(words)-1])
}

func main() {
	s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord1(s))
}
