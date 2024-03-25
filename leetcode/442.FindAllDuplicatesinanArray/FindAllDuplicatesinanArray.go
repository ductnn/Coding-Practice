// https://leetcode.com/problems/find-all-duplicates-in-an-array/description/

package main

import "fmt"

func findDuplicates(nums []int) []int {
	cnt := make(map[int]int)
	res := []int{}

	for _, num := range nums {
		cnt[num]++
		// fmt.Println(cnt)
	}

	for num, count := range cnt {
		if count > 1 {
			res = append(res, num)
		}
	}

	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}
