package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	count := 0
	i, j := 0, len(nums)-1

	for i < j {
		if nums[i]+nums[j] == k {
			count++
			i++
			j--
		} else if nums[i]+nums[j] > k {
			j--
		} else {
			i++
		}
	}
	return count
}

func maxOperations1(nums []int, k int) int {
	result := 0
	count := make(map[int]int, len(nums))
	for _, v := range nums {
		if count[k-v] > 0 {
			count[k-v]--
			result++
		} else {
			count[v]++
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	k := 5
	fmt.Println(maxOperations1(nums, k))
}
