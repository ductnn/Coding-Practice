package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := append(nums1[0:m], nums2[0:n]...)

	QuickSort(temp)
}

func QuickSort(array []int) []int {
	// Write your code here.
	helper(array, 0, len(array)-1)

	return array
}

func helper(array []int, left, right int) {
	if len(array) <= 1 {
		return
	}

	if left < right {
		pivot := Partition(array, left, right)
		helper(array, left, pivot-1)
		helper(array, pivot+1, right)
	}
}

func Partition(array []int, left, right int) int {
	index := left - 1
	// Defind pivot element
	p := array[right]

	for i := left; i < right; i++ {
		if array[i] <= p {
			index++
			Swap(array, index, i)
		}
	}

	Swap(array, index+1, right)

	return index + 1
}

func Swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)

	fmt.Println(nums1)
}
