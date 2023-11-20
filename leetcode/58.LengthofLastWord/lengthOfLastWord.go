package main

import (
	"fmt"
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

func main() {
	s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}
