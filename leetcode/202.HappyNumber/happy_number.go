package main

import (
	"fmt"
)

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	s, f := n, getSquareNumber(n)
	for s != f {
		s = getSquareNumber(s)
		f = getSquareNumber(getSquareNumber(f))
	}

	if s == 1 {
		return true
	}

	return false
}

func getSquareNumber(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}

func main() {
	n := 19
	fmt.Println(isHappy(n))
}
