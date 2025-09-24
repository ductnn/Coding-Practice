// https://leetcode.com/problems/fraction-to-recurring-decimal/

package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	sign := ""
	if numerator < 0 && denominator > 0 || numerator > 0 && denominator < 0 {
		sign = "-"
	}
	numerator = abs(numerator)
	denominator = abs(denominator)

	integerPart := numerator / denominator
	remainder := numerator % denominator

	fractionPart := ""
	if remainder != 0 {
		fractionPart = "."
	}

	remainderMap := make(map[int]int)
	repeatIndex := -1

	for remainder != 0 {
		if index, ok := remainderMap[remainder]; ok {
			repeatIndex = index
			break
		}
		remainderMap[remainder] = len(fractionPart)
		remainder *= 10
		fractionPart += strconv.Itoa(remainder / denominator)
		remainder %= denominator
	}

	if repeatIndex != -1 {
		fractionPart = fractionPart[:repeatIndex] + "(" + fractionPart[repeatIndex:] + ")"
	}

	return sign + strconv.Itoa(integerPart) + fractionPart
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(fractionToDecimal(-1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(4, 333))
	fmt.Println(fractionToDecimal(1, 5))
	fmt.Println(fractionToDecimal(1, 17))
	fmt.Println(fractionToDecimal(1, 18))
	fmt.Println(fractionToDecimal(1, 19))
}
