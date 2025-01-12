// https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/

package main

import "fmt"

func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}

	// Forward pass
	balance := 0
	for i := 0; i < len(s); i++ {
		if locked[i] == '1' && s[i] == ')' {
			balance--
		} else {
			balance++
		}
		if balance < 0 {
			return false
		}
	}

	// Backward pass
	balance = 0
	for i := len(s) - 1; i >= 0; i-- {
		if locked[i] == '1' && s[i] == '(' {
			balance--
		} else {
			balance++
		}
		if balance < 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "))()))"
	locked := "010100"

	fmt.Println(canBeValid(s, locked))
}
