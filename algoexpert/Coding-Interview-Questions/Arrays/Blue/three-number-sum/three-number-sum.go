package main

import (
	"fmt"
)

func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.
	bubbleSort(array)
	ans := [][]int{}

	for i := 0; i < len(array); i++ {
		left, right := i+1, len(array)-1

		for left < right {
			sum := array[left] + array[i] + array[right]
			switch {
			case sum < target:
				left++
			case sum > target:
				right--
			case sum == target:
				ans = append(ans, []int{array[i], array[left], array[right]})
				left++
				right--
			}
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

func main() {
	arrays := []int{12, 3, 1, 2, -6, 5, -8, 6}
	targetSum := 0

	fmt.Println(ThreeNumberSum(arrays, targetSum))
}
