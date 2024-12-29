// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/description

package main

import (
	"fmt"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	total := 0
	for i := range len(seats) {
		total += abs(seats[i] - students[i])
	}

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	seats := []int{4, 1, 5, 9}
	students := []int{1, 3, 2, 6}

	fmt.Println(minMovesToSeat(seats, students))
}
