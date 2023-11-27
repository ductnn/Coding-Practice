package main

import (
	"fmt"
)

func grayCode(n int) []int {
	ans := []int{}
	for i := 0; i < 1<<n; i++ {
		ans = append(ans, i^(i>>1))
	}
	return ans
}

func main() {
	n := 2
	fmt.Println(grayCode(n))
}
