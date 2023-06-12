package main

import (
	"fmt"
)

func sortColors4(nums []int) {
	tmp := make([]int, 3)

	for _, num := range nums {
		tmp[num]++
	}

	i := 0
	for color := 0; color <= 2; color++ {
		for tmp[color] > 0 {
			nums[i] = color
			i++
			tmp[color]--
		}
	}
}

func sortColors3(nums []int) {
	isDone := false

	for !isDone {
		isDone = true

		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[i-1] {
				Swap(nums, i, i-1)
				isDone = false
			}
		}
	}
}

func sortColors2(nums []int) {
	size := len(nums) - 1
	i, j, k := 0, 0, size

	for i <= k {
		if nums[i] == 0 {
			Swap(nums, j, i)
			j++
		} else if nums[i] == 2 {
			Swap(nums, i, k)
			k--
			i--
		}
		i++
	}
}

func sortColors1(nums []int) {
	size := len(nums) - 1
	i, j, k := 0, 0, size

	for j <= k {
		switch nums[j] {
		case 0:
			Swap(nums, i, j)
			i++
			j++
		case 1:
			j++
		case 2:
			Swap(nums, j, k)
			k--
		}
	}
}

func sortColors(nums []int) {
	QuickSort(nums)
}

func QuickSort(array []int) []int {
	// Write your code here.
	helper(array, 0, len(array)-1)

	return array
}

func helper(array []int, left, right int) {
	if len(array) <= 1 {
		return
	}

	if left < right {
		pivot := Partition(array, left, right)
		helper(array, left, pivot-1)
		helper(array, pivot+1, right)
	}
}

func Partition(array []int, left, right int) int {
	index := left - 1
	// Defind pivot element
	p := array[right]

	for i := left; i < right; i++ {
		if array[i] <= p {
			index++
			Swap(array, index, i)
		}
	}

	Swap(array, index+1, right)

	return index + 1
}

func Swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	nums := []int{1, 2, 0, 0, 0, 2, 1, 2, 1, 2}
	sortColors(nums)

	fmt.Println(nums)

	nums1 := []int{1, 2, 0, 0, 0, 2, 1, 2, 1, 2, 0, 0, 0, 0, 2}
	sortColors1(nums1)

	fmt.Println(nums1)

	nums2 := []int{1, 2, 0, 0, 0, 2, 1, 2, 1, 2, 0, 0, 0, 0, 2}
	sortColors2(nums2)

	fmt.Println(nums2)

	nums3 := []int{1, 2, 0, 0, 0, 2, 1, 2, 1, 2, 0, 0, 0, 0, 2}
	sortColors3(nums3)

	fmt.Println(nums3)

	nums4 := []int{1, 2, 0, 0, 0, 2, 1, 2, 1, 2, 0, 0, 0, 0, 2}
	sortColors4(nums4)

	fmt.Println(nums4)
}
