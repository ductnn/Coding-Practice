// https://leetcode.com/problems/successful-pairs-of-spells-and-potions

package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	m := len(spells)
	n := len(potions)

	result := make([]int, m)
	for i := 0; i < m; i++ {
		left, right := 0, n
		for left < right {
			mid := left + (right-left)>>1
			if int64(potions[mid]*spells[i]) >= success {
				right = mid
			} else {
				left = mid + 1
			}
		}
		result[i] = n - left
	}
	return result
}

func successfulPairs1(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	m := len(spells)
	n := len(potions)

	result := make([]int, m)
	for i := 0; i < m; i++ {
		temp := binarySearch(potions, spells[i], int(success))
		result[i] = n - temp
	}
	return result
}

func binarySearch(nums []int, temp, success int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid]*temp >= success {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func main() {
	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	success := 7

	fmt.Println(successfulPairs1(spells, potions, int64(success)))
}
