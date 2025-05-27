// https://leetcode.com/problems/divisible-and-non-divisible-sums-difference

package main

import "fmt"

func differenceOfSums(n int, m int) int {
	sum1 := 0
	sum2 := 0

	for i := 1; i <= n; i++ {
		if i%m == 0 {
			sum2 += i
		} else {
			sum1 += i
		}
	}

	return sum1 - sum2
}

func main() {
	fmt.Println(differenceOfSums(10, 3))
}
