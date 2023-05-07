package main

import (
	"fmt"
)

func BubbleSort(array []int) []int {
	// Write your code here.
	// ==== Simple ====
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				// array[j], array[j+1] = array[j+1], array[j]
				Swap(array, j, j+1)
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

	fmt.Println(BubbleSort(array))
}
