package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / pow(x, -n)
	}

	return pow(x, n)
}

func duma(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if x == -1 && n&1 == 1 {
		return -1
	}

	if n == 0 || x == 1 {
		return 1
	}

	var res float64 = 1

	for i := 0; i < n; i++ {
		res = res * x
	}

	return res
}

func pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if x == -1 && n&1 == 1 {
		return -1
	}

	if n == 0 || x == 1 {
		return 1
	}

	res := pow(x, n>>1)
	if n&1 == 0 {
		return res * res
	}

	return res * res * x
}

func main() {
	x := -1.00000
	n := -3

	fmt.Println(myPow(x, n))
}
