// https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/

package main

import "fmt"

func findThePrefixCommonArray(A []int, B []int) []int {
	seen := make(map[int]bool) // To track elements seen so far
	result := make([]int, len(A))
	commonCount := 0

	for i := 0; i < len(A); i++ {
		if seen[A[i]] {
			commonCount++
		} else {
			seen[A[i]] = true
		}

		if seen[B[i]] {
			commonCount++
		} else {
			seen[B[i]] = true
		}

		result[i] = commonCount
	}

	return result
}

func main() {
	A := []int{1, 3, 2, 4}
	B := []int{3, 1, 2, 4}
	fmt.Println(findThePrefixCommonArray(A, B))
}
