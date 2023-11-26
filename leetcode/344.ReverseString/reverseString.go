package main

import (
	"bytes"
	"fmt"
)

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

func reverseString1(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
}

func reverseString2(s []byte) {
	i := 0
	j := len(s) - i - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseString3(s []byte) {
	for i, j := 0, len(s)-1; i < j; j = len(s) - i - 1 {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseString4(s []byte) {
	b := bytes.Runes(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	copy(s, string(b))
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}

	reverseString3(s)

	fmt.Println(string(s))
}
