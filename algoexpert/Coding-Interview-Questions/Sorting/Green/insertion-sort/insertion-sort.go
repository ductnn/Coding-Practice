package main

import (
	"fmt"
)

func InsertionSort(array []int) []int {
	// Write your code here.
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			Swap(array, j, j-1)
		}
	}

	return array
}

func Swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}

	fmt.Println(InsertionSort(array))
}
