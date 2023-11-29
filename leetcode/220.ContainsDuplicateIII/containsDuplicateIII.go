package main

import (
	"fmt"
)

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	// Input: nums = [1,2,3,1], indexDiff = 3, valueDiff = 0
	// Output: true
	// (i, j) = (0, 3)
	// abs(i-j) <= indexDiff, abs(nums[i] - nums[j]) <= valueDiff
	// return true

	if (indexDiff < 1) || (valueDiff < 0) {
		return false
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= min(len(nums)-1, i+indexDiff); j++ {
			if abs(nums[i], nums[j]) <= valueDiff {
				return true
			}
		}
	}

	return false
}

// Sliding window
func containsNearbyAlmostDuplicate1(nums []int, indexDiff int, valueDiff int) bool {
	if valueDiff < 0 {
		return false
	}

	buckets := make(map[int]int)
	width := valueDiff + 1

	for i, num := range nums {
		bucket := getBucket(num, width)

		if _, ok := buckets[bucket]; ok {
			return true
		}

		if val, ok := buckets[bucket-1]; ok && abs(num, val) < width {
			return true
		}

		if val, ok := buckets[bucket+1]; ok && abs(num, val) < width {
			return true
		}

		buckets[bucket] = num
		if i >= indexDiff {
			delete(buckets, getBucket(nums[i-indexDiff], width))
		}
	}

	return false
}

func getBucket(num, width int) int {
	if num >= 0 {
		return num / width
	}
	return (num+1)/width - 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	nums := []int{1, 2, 3, 1}
	indexDiff := 3
	valueDiff := 0

	fmt.Println(containsNearbyAlmostDuplicate1(nums, indexDiff, valueDiff))
}
