package main

import (
	"fmt"
)

func findSpecialInteger(arr []int) int {
	for k, v := range arr {
		if v == arr[k+(len(arr)>>2)] {
			return v
		}
	}
	return 0
}

func findSpecialInteger1(arr []int) int {
	if len(arr) < 4 {
		return arr[0]
	}
	count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			count++
			if count > len(arr)/4 {
				return arr[i]
			}
		} else {
			count = 1
		}

	}
	return -1
}

func main() {
	arr := []int{1, 1, 1, 1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 12, 12, 12}
	fmt.Println(findSpecialInteger1(arr))
}
