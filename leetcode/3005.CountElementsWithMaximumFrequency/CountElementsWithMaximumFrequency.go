// https://leetcode.com/problems/count-elements-with-maximum-frequency

package main

import (
	"fmt"
)

func maxFrequencyElements(nums []int) int {
	res := 0
	cnt := [101]int{}
	for _, v := range nums {
		cnt[v]++
	}

	mx := -1
	for _, v := range cnt {
		if mx < v {
			mx, res = v, v
		} else if mx == v {
			res += v
		}
	}

	return res
}

func main() {
	nums := []int{1, 2, 2, 3, 1, 4}

	fmt.Println(maxFrequencyElements(nums))
}
