package main

import (
	"fmt"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	s1, s2 := make(map[int]bool), make(map[int]bool)
	for _, v := range nums1 {
		s1[v] = true
	}
	for _, v := range nums2 {
		s2[v] = true
	}
	result := make([][]int, 2)
	for v := range s1 {
		if !s2[v] {
			result[0] = append(result[0], v)
		}
	}
	for v := range s2 {
		if !s1[v] {
			result[1] = append(result[1], v)
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	fmt.Println(findDifference(nums1, nums2))
}
