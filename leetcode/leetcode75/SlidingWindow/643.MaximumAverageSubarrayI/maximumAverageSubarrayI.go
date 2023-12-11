package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	result := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if result < sum {
			result = sum
		}
	}

	return float64(result) / float64(k)
}

func findMaxAverage1(nums []int, k int) float64 {
	j := 0
	res := average(nums[:k])
	for i := k; i <= len(nums); i++ {
		res = max(res, average(nums[j:i]))
		j++
	}
	return res
}

func average(temp []int) float64 {
	sum := 0
	for i := range temp {
		sum += temp[i]
	}
	return float64(sum) / float64(len(temp))
}

func max(a, b float64) float64 {
	if a < b {
		return b
	}
	return a
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	fmt.Println(findMaxAverage(nums, k))
}
