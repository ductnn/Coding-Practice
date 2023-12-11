package main

import (
	"fmt"
)

func minFlips(a int, b int, c int) int {
	result := 0
	for a > 0 || b > 0 || c > 0 {
		am := a % 2
		bm := b % 2
		cm := c % 2

		a = a >> 1
		b = b >> 1
		c = c >> 1

		if cm == 1 && am == 0 && bm == 0 {
			result++
		} else if cm == 0 {
			if am == 1 {
				result++
			}
			if bm == 1 {
				result++
			}
		}
	}
	return result
}

func minFlips1(a int, b int, c int) int {
	result := 0
	for i := 0; i < 30; i++ {
		x, y, z := a>>i&1, b>>i&1, c>>i&1
		if (x | y) != z {
			if x == 1 && y == 1 {
				result += 2
			} else {
				result++
			}
		}
	}
	return result
}

func main() {
	a, b, c := 2, 4, 7
	fmt.Println(minFlips(a, b, c))
}
