package main

import (
	"fmt"
)

const (
	MAX = 2147483647
	MIN = -2147483648
)

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x *= -1
	}

	res := 0
	for x > 0 {
		temp := x % 10
		res = res*10 + temp
		x /= 10
	}

	res = res * sign

	if res > MAX || res < MIN {
		res = 0
	}

	return res
}

func main() {
	x := 1534236469

	fmt.Println(reverse(x))
}
