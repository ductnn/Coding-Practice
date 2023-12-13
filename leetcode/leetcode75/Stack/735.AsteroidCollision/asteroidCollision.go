package main

import (
	"fmt"
)

func asteroidCollision(asteroids []int) []int {
	result := []int{}
	for _, v := range asteroids {
		if v > 0 {
			result = append(result, v)
		} else {
			for len(result) > 0 && result[len(result)-1] > 0 && result[len(result)-1] < -v {
				result = result[:len(result)-1]
			}
			if len(result) > 0 && result[len(result)-1] == -v {
				result = result[:len(result)-1]
			} else if len(result) == 0 || result[len(result)-1] < 0 {
				result = append(result, v)
			}
		}
	}
	return result
}

func main() {
	asteroids := []int{5, 10, -5}
	fmt.Println(asteroidCollision(asteroids))
}
