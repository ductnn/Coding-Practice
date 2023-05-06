package main

import (
	"fmt"
)

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	ans := make([]int, len(array))

	for i, value := range array {
		ans[i] = value * value
	}

	return bubbleSort(ans)
}

func bubbleSort(array []int) []int {
	// Simple
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}

func main() {
	array := []int{-7, -3, 1, 9, 22, 30}

	fmt.Println(SortedSquaredArray(array))
}
