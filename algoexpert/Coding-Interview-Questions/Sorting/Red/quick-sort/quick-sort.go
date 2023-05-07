package main

import (
	"fmt"
)

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
	array := []int{8, 5, 2, 9, 5, 6, 3}

	fmt.Println(QuickSort(array))
}
