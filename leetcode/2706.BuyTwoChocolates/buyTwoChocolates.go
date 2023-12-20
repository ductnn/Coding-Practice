// https://leetcode.com/problems/buy-two-chocolates

package main

import (
	"fmt"
	"sort"
)

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	if money < prices[0]+prices[1] {
		return money
	}
	return money - prices[0] - prices[1]
}

func main() {
	prices := []int{1, 2, 2}
	money := 3
	fmt.Println(buyChoco(prices, money))
}
