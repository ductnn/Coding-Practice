package main

import (
	"fmt"
)

func ThreeNumberSort(array []int, order []int) []int {
	// Write your code here.
	orderFirst, orderSecond := order[0], order[1]

	// Three index in Array
	i, j, k := 0, 0, len(array)-1

	for j <= k {
		value := array[j]
		switch {
		case value == orderFirst:
			Swap(array, i, j)
			i++
			j++
		case value == orderSecond:
			j++
		default:
			Swap(array, j, k)
			k--
		}
	}
	return array
}

func Swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	array := []int{1, 0, 0, -1, -1, 0, 1, 1}
	order := []int{0, 1, -1}

	fmt.Println(ThreeNumberSort(array, order))
}
