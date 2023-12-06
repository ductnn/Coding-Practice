package main

import (
	"fmt"
)

func totalMoney(n int) int {
	weeks := n / 7
	leftOverDays := n % 7

	return (weeks * 28) + (weeks * (weeks - 1) >> 1 * 7) + (leftOverDays * (2*weeks + leftOverDays + 1) >> 1)
}

func totalMoney1(n int) int {
	total := 0
	week := 0
	for i := 1; i <= n; i++ {
		day := i % 7
		if day == 0 {
			day = 7
		}
		total += day + week
		week = i / 7
	}
	return total
}

func main() {
	n := 10
	fmt.Println(totalMoney1(n))
}
