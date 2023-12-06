package main

import (
	"fmt"
	"slices"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := []bool{}
	temp := slices.Max(candies)
	for i := range candies {
		result = append(result, candies[i]+extraCandies >= temp)
	}
	return result
}

func kidsWithCandies1(candies []int, extraCandies int) []bool {
	result := []bool{}
	temp := 0
	for i := range candies {
		temp = max(temp, candies[i])
	}
	for i := range candies {
		result = append(result, candies[i]+extraCandies >= temp)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	fmt.Println(kidsWithCandies1(candies, extraCandies))
}
