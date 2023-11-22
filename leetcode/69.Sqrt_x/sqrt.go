package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	l, r := 0, x

	for l < r {
		mid := l + (r-l+1)>>1
		if mid <= x/mid {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l
}

func main() {
	x := 4
	fmt.Println(mySqrt(x))
}
