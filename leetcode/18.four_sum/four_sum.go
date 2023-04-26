package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	result := [][]int{}

	for i := 0; i < length-3; i++ {
		if i > 0 && (nums[i-1] == nums[i]) {
			continue
		}

		for j := i + 1; j < length-2; j++ {
			if j > i+1 && (nums[j] == nums[j-1]) {
				continue
			}
			left, right := i+1, length-1

			for left < right {
				sum := nums[left] + nums[i] + nums[j] + nums[right]

				switch {
				case sum < target:
					left++
				case sum > target:
					right--
				default:
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for right > left && nums[right] == nums[right-1] {
						right--
					}
				}
			}
		}
	}

	return result
}

func main() {
	input := []int{-3, -1, 0, 2, 4, 5}
	target := 0

	// sort.Ints(input)
	// fmt.Println(input)
	fmt.Println(fourSum(input, target))
}
