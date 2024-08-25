// https://leetcode.com/problems/three-consecutive-odds/description/

package main

import (
	"fmt"
)

func threeConsecutiveOdds(arr []int) bool {
    cnt := 0
	for _, v := range arr {
		if v%2 == 1 {
			cnt++
		} else {
			cnt = 0
		}
		if cnt == 3 {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{2,6,4,1}
	fmt.Println(threeConsecutiveOdds(arr))
}
