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

func fourSum1(nums []int, target int) [][]int {
	result := [][]int{}
	n := len(nums)

	if n < 4 {
		return result
	}

	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k, l := j+1, n-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum < target {
					k++
				} else if sum > target {
					l--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				}
			}
		}
	}
	return result
}

func main() {
	input := []int{1, 0, -1, 0, -2, 2}
	target := 0

	// sort.Ints(input)
	// fmt.Println(input)
	fmt.Println(fourSum1(input, target))
}
