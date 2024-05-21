// https://leetcode.com/problems/first-bad-version

package main

import (
	"fmt"
)

// Assume this is a pre-defined API that checks if a version is bad
var firstBad int

func isBadVersion(version int) bool {
	return version >= firstBad
}

// FirstBadVersion finds the first bad version among n versions
func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)>>1
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	n := 5
	firstBad = 4
	fmt.Println(firstBadVersion(n))
}
