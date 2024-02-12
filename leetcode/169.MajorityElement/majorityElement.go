// https://leetcode.com/problems/majority-element/

package main

import "fmt"

// Moore Voting Algo
func majorityElement(nums []int) int {
	res := 0
	cnt := 0

	for _, num := range nums {
		if cnt == 0 {
			res = num
		}
		if num == res {
			cnt++
		} else {
			cnt--
		}
	}

	return res
}

// Hashmap
func majorityElement1(nums []int) int {
	n := len(nums)
	cnt := make(map[int]int)

	for _, num := range nums {
		cnt[num]++

		if cnt[num] > n/2 {
			return num
		}
	}

	return -1
}

func main() {
	nums := []int{2, 2, 2, 1, 1, 1, 2}
	fmt.Println(majorityElement1(nums))
}
