package main

import (
	"fmt"
	"math"
)

const (
	MAX = math.MaxInt32
)

func SmallestDifference(array1, array2 []int) []int {
	// Write your code here.
	bubbleSort(array1)
	bubbleSort(array2)

	ans := []int{}
	smallestDiff, currentDiff := MAX, MAX
	point1, point2 := 0, 0

	for point1 < len(array1) && point2 < len(array2) {
		num1, num2 := array1[point1], array2[point2]
		currentDiff = abs(num1 - num2)

		if num1 == num2 {
			ans = []int{num1, num2}
		}

		if num1 < num2 {
			point1++
		} else {
			point2++
		}

		if currentDiff < smallestDiff {
			smallestDiff = currentDiff
			ans = []int{num1, num2}
		}
	}

	return ans
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

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func main() {
	arrayOne := []int{-1, 5, 10, 20, 28, 3}
	arrayTwo := []int{26, 134, 135, 15, 17}

	fmt.Println(SmallestDifference(arrayOne, arrayTwo))
}
