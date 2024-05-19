// https://leetcode.com/problems/find-smallest-letter-greater-than-target/

package main

import (
	"fmt"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	for left <= right {
		mid := left + (right-left)>>1
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return letters[left%len(letters)]
}

func main() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('a')

	fmt.Println(nextGreatestLetter(letters, target))
}
