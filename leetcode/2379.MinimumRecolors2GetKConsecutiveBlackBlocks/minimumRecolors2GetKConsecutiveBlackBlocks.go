package main

import (
	"fmt"
)

func minimumRecolors(blocks string, k int) int {
	count := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			count++
		}
	}

	result := count
	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			count++
		}
		if blocks[i-k] == 'W' {
			count--
		}
		if result > count {
			result = count
		}
	}

	return result
}

func main() {
	blocks := "WBBWWBBWBW"
	k := 7

	fmt.Println(minimumRecolors(blocks, k))
}
