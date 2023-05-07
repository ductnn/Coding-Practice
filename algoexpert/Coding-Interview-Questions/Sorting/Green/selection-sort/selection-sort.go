package main

import (
	"fmt"
)

func SelectionSort(array []int) []int {
	// Write your code here.
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				Swap(array, i, j)
			}
		}
	}
	return array
}

func Swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}

	fmt.Println(SelectionSort(array))
}
