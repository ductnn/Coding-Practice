package main

import (
	"fmt"
)

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

var pick = 6

func guess(num int) int {
	if num == pick {
		return 0
	} else if num > pick {
		return -1
	}
	return 1
}

func guessNumber(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)>>1
		if guess(mid) <= 0 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	n := 10
	fmt.Println(guessNumber(n))
}
