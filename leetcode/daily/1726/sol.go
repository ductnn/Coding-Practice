package main

import "fmt"

func tupleSameProduct(nums []int) int {
	productCount := make(map[int]int)
	count := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			product := nums[i] * nums[j]
			if val, ok := productCount[product]; ok {
				count += val
			}
			productCount[product]++
		}
	}

	return count * 8
}

func tupleSameProduct1(nums []int) int {
	productCount := make(map[int]int)
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			product := nums[i] * nums[j]
			productCount[product]++
			fmt.Println(productCount)
		}

	}

	count := 0
	for _, v := range productCount {
		count += v * (v - 1) * 8
	}

	return count / 2
}

func main() {
	nums := []int{1, 2, 4, 5, 10}
	println(tupleSameProduct1(nums))
}
