package main

import "fmt"

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	seqIdx := 0

	for _, num := range array {
		if num == sequence[seqIdx] {
			seqIdx++
		}

		if seqIdx == len(sequence) {
			return true
		}
	}

	return false
}

func IsValidSubsequence1(array []int, sequence []int) bool {
	// Write your code here.
	seqIdx := 0
	arrIdx := 0

	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx++
		}

		arrIdx++
	}

	return seqIdx == len(sequence)
}

func main() {
	arrays := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 110}

	fmt.Println(IsValidSubsequence1(arrays, sequence))
}
