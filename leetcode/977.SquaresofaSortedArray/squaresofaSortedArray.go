package main

import (
	"fmt"
)

func sortedSquares1(nums []int) []int {
	l := 0
	r := len(nums) - 1
	res := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		if abs(nums[l]) > nums[r] {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = res[i]
	}

	return res
}

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		res[i] = nums[i] * nums[i]
	}

	QuickSort(res)

	return res
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

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func main() {
	duma := []int{-7, -3, 2, 3, 11}

	fmt.Println(sortedSquares1(duma))
}
