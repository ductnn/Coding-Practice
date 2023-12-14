package main

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				result[i] = j - i
				break
			}
		}
	}
	return result
}

func dailyTemperatures1(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stk := []int{}
	for i, t := range temperatures {
		for len(stk) > 0 && temperatures[stk[len(stk)-1]] < t {
			j := stk[len(stk)-1]
			result[j] = i - j
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	return result
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}
