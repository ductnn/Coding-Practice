package main

import (
	"fmt"
)

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	i, j := 0, len(arr)-1
	for i+1 < len(arr)-1 && arr[i] < arr[i+1] {
		i++
	}
	for j-1 > 0 && arr[j-1] > arr[j] {
		j--
	}

	return i == j
}

func main() {
	arr := []int{0, 3, 2, 1}
	fmt.Println(validMountainArray(arr))
}
