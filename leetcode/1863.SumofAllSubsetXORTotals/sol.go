// https://leetcode.com/problems/sum-of-all-subset-xor-totals

package main

import (
	"fmt"
)

func subsetXORSum(nums []int) int {
	n := len(nums)
	res := 0
	for i := 1; i < 1<<n; i++ {
		sum := 0
		for k, v := range nums {
			if i>>k&1 == 1 {
				sum ^= v
			}
		}
		res += sum
	}
	return res
}

func main() {
	nums := []int{5, 1, 6}
	fmt.Println(subsetXORSum(nums))
}
