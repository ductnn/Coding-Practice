package main

import (
	"fmt"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	result := 0
	s := strconv.Itoa(num)

	for i := 0; i < len(s)-k+1; i++ {
		t, _ := strconv.Atoi(s[i : i+k])

		if t > 0 && num%t == 0 {
			result++
		}
	}

	return result
}

func main() {
	num := 430043
	k := 2

	fmt.Println(divisorSubstrings(num, k))
}
