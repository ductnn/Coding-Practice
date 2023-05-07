package main

import (
	"fmt"
)

func IsMonotonic(array []int) bool {
	// Write your code here.
	inc, dec := true, true

	for i := 1; i < len(array); i++ {
		inc = inc && array[i-1] <= array[i]
		dec = dec && array[i-1] >= array[i]
	}

	return inc || dec
}

func main() {
	array := []int{-1, -5, -10, -1100, -1100, -1101, -1102, 1}

	fmt.Println(IsMonotonic(array))
}
