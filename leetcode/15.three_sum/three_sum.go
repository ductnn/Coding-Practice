package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	result := [][]int{}

	for i := 0; i < length-1; i++ {
		if i > 0 && (nums[i-1] == nums[i]) {
			continue
		}
		left, right := i+1, length-1

		for left < right {
			sum := nums[left] + nums[i] + nums[right]

			switch {
			case sum < 0:
				left++
			case sum > 0:
				right--
			case sum == 0:
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for right > left && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return result
}

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}

	// sort.Ints(input)
	// fmt.Println(input)
	fmt.Println(threeSum(input))
}
