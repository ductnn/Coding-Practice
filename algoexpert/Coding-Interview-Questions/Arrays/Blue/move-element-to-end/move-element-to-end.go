package main

import (
	"fmt"
)

func MoveElementToEnd(array []int, toMove int) []int {
	// Write your code here.
	for i, j := 0, 0; i < len(array); i++ {
		if array[i] != toMove {
			array[i], array[j] = array[j], array[i]
			j++
		}
	}

	return array
}

func main() {
	array := []int{2, 1, 2, 2, 2, 3, 4, 2}
	toMove := 2

	fmt.Println(MoveElementToEnd(array, toMove))
}
