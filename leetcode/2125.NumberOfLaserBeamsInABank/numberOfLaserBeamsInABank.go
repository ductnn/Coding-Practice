// https://leetcode.com/problems/number-of-laser-beams-in-a-bank/

package main

import (
	"fmt"
)

func numberOfBeams(bank []string) int {
	result, last := 0, 0
	for _, b := range bank {
		tmp := 0
		for _, c := range b {
			if c == '1' {
				tmp++
			}
		}

		if tmp > 0 {
			result += tmp * last
			last = tmp
		}
	}

	return result
}

func main() {
	bank := []string{"011001", "000000", "010100", "001000"}

	fmt.Println(numberOfBeams(bank))
}
