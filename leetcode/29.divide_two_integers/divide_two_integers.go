package main

import (
	"fmt"
	"math"
)

const (
	MAX = math.MaxInt32
	MIN = math.MinInt32
)

func divide(dividend int, divisor int) int {
	if (divisor == 0) || (dividend == MIN && divisor == -1) {
		return MAX
	}
	if divisor == 1 {
		return dividend
	}
	if dividend == 0 {
		return 0
	}

	result := 0
	sign := -1

	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		sign = 1
	}

	absDividend, absDivisor := abs(dividend), abs(divisor)
	for absDividend >= absDivisor {
		temp := absDivisor
		m := 1

		for temp<<1 <= absDividend {
			temp <<= 1
			m <<= 1
		}

		absDividend -= temp
		result += m
	}

	return sign * result
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func main() {
	dividend := 7
	divisor := -3

	fmt.Println(divide(dividend, divisor))
}
