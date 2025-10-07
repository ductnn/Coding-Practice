package main

import (
	"fmt"
	"sort"
)

func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	full := make(map[int]int)
	dryDays := []int{}

	for i, lake := range rains {
		// Rain Day
		if lake > 0 {
			ans[i] = -1
			if day, ok := full[lake]; ok {
				// Lake full, need 1 dry day to dry it
				idx := findDryDayAfter(dryDays, day) // idx in dryDays slice
				if idx == -1 {
					return []int{} // No dry day found -> flood
				}
				// Use the actual dry day (value), not the index
				actualDay := dryDays[idx]
				ans[actualDay] = lake
				// Remove the used dry day from the list (by index)
				dryDays = append(dryDays[:idx], dryDays[idx+1:]...)
			}
			full[lake] = i
		} else {
			// Dry Day
			dryDays = append(dryDays, i)
			ans[i] = 1
		}
	}

	return ans
}

// Find the index in dryDays of the first day > day
// Ex: dryDays = [2, 4, 5], day = 3 => idx = 1
// Ex: dryDays = [2, 4, 5], day = 6 => idx = -1
func findDryDayAfter(dryDays []int, day int) int {
	idx := sort.Search(len(dryDays), func(i int) bool { return dryDays[i] > day })
	if idx == len(dryDays) {
		return -1
	}
	return idx
}

func main() {
	rains := []int{1, 2, 0, 0, 2, 1}
	fmt.Println(avoidFlood(rains)) // => [-1 -1 2 1 -1 -1]
}
