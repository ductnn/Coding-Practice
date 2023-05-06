package main

import (
	"fmt"
	"sort"
)

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	ans := []int{}

	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] == target-array[i] {
				ans = []int{array[i], array[j]}
			}
		}
	}

	return ans
}

func TwoNumberSum1(array []int, target int) []int {
	// Write your code here.
	ans := []int{}

	// bubbleSort(array)
	left, right := 0, len(array)-1

	for left < right {
		currentSum := array[left] + array[right]

		if currentSum == target {
			ans = []int{array[left], array[right]}
		} else if currentSum < target {
			left += 1
		} else if currentSum > target {
			right -= 1
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

func TwoNumberSum2(array []int, target int) []int {
	// Write your code here.
	ans := []int{}

	sort.Ints(array)
	left, right := 0, len(array)-1

	for left < right {
		currentSum := array[left] + array[right]

		if currentSum == target {
			ans = []int{array[left], array[right]}
		} else if currentSum < target {
			left += 1
		} else if currentSum > target {
			right -= 1
		}
	}

	return ans
}

func main() {
	arrays := []int{3, 5, -4, 8, 11, 1, -1, 6}
	targetSum := 10
	fmt.Println(TwoNumberSum(arrays, targetSum))
}
