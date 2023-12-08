package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	scan := 0  // scan in chars[]
	write := 0 // write in result[]

	for scan < len(chars) {
		chars[write] = chars[scan]
		count := 0

		// check consecutive repeating characters
		for scan < len(chars) && chars[write] == chars[scan] {
			count++
			scan++
		}

		// check compresses
		if count > 1 {
			for _, c := range []byte(strconv.Itoa(count)) {
				write++
				chars[write] = c
			}
		}
		write++
	}
	return write
}

func main() {
	chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	fmt.Println(compress(chars))
}
