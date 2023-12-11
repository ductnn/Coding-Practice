package main

import (
	"fmt"
	"math"
)

const (
	MAX = math.MaxInt32
)

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, MAX
	for left < right {
		mid := left + (right-left)>>1
		sum := 0
		for _, v := range piles {
			sum += (v + mid - 1) / mid
		}
		if sum <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	piles := []int{30, 11, 23, 4, 20}
	h := 6

	fmt.Println(minEatingSpeed(piles, h))
}
