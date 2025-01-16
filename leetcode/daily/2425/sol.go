// https://leetcode.com/problems/bitwise-xor-of-all-pairings/description/

package main

import "fmt"

// TLE
func xorAllNums(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)

	res := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res ^= nums1[i] ^ nums2[j]
		}
	}

	return res
}

func xorAllNums1(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)

	xor1, xor2 := 0, 0

	for i := 0; i < n; i++ {
		xor1 ^= nums1[i]
	}

	for i := 0; i < m; i++ {
		xor2 ^= nums2[i]
	}

	if n%2 == 0 && m%2 == 0 {
		return 0
	} else if n%2 == 0 {
		return xor1
	} else if m%2 == 0 {
		return xor2
	} else {
		return xor1 ^ xor2
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	// nums1 := []int{2, 1, 3}
	// nums2 := []int{10, 2, 5, 0}

	fmt.Println(xorAllNums1(nums1, nums2))
}
