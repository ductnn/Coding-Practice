package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	result, tmp := 0, math.MaxInt64

	for i := 0; i < length; i++ {
		if i > 0 && (nums[i-1] == nums[i]) {
			continue
		}
		left, right := i+1, length-1

		for left < right {
			sum := nums[left] + nums[i] + nums[right]

			switch {
			case sum < target:
				left++
				if tmp > target-sum {
					tmp = target - sum
					result = sum
				}
			case sum > target:
				right--
				if tmp > sum-target {
					tmp = sum - target
					result = sum
				}
			default:
				return sum
			}
		}
	}

	return result
}

func main() {
	input := []int{-1, 2, 1, -4}
	target := 1

	fmt.Println(threeSumClosest(input, target))
	fmt.Println(math.MaxInt64)
}
