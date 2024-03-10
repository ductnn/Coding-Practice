// https://leetcode.com/problems/intersection-of-two-arrays

package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}

	mx := [1001]bool{}
	for _, x := range nums1 {
		mx[x] = true
	}

	for _, x := range nums2 {
		if mx[x] {
			res = append(res, x)
			mx[x] = false
		}
	}

	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersection(nums1, nums2))
}
