// https://leetcode.com/problems/time-needed-to-buy-tickets/description/

package main

import (
	"fmt"
)

func timeRequiredToBuy(tickets []int, k int) int {
	res := 0

	for i, t := range tickets {
		if i <= k {
			res += min(tickets[k], t)
		} else {
			res += min(tickets[k]-1, t)
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return a
}

func main() {
	tickets := []int{2, 3, 2}
	k := 2

	fmt.Println(timeRequiredToBuy(tickets, k))
}
