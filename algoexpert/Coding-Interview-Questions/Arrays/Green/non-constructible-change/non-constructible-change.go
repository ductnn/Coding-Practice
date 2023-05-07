package main

import (
	"fmt"
)

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	bubbleSort(coins)
	currentCoinCreated := 0

	for _, coin := range coins {
		if coin-currentCoinCreated > 1 {
			break
		}

		currentCoinCreated += coin
	}

	return currentCoinCreated + 1
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
	coins := []int{5, 7, 1, 1, 2, 3, 22}

	fmt.Println(NonConstructibleChange(coins))
}
