package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
				flowerbed[i] = 1
				count++
			}
		}
	}
	return count >= n
}

func canPlaceFlowers1(flowerbed []int, n int) bool {
	m, j := 0, 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			j++
		} else {
			m += (max(j-1, 0) + 1) >> 1
			j = -1
		}
	}
	m += (max(j, 0) + 1) >> 1
	return n <= m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 2

	fmt.Println(canPlaceFlowers(flowerbed, n))
}
