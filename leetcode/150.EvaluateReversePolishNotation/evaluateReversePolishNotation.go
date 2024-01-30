// https://leetcode.com/problems/evaluate-reverse-polish-notation

package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	stackSize := 0

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			var res int
			op1, op2 := stack[stackSize-2], stack[stackSize-1]

			switch token {
			case "+":
				res = op1 + op2
			case "-":
				res = op1 - op2
			case "*":
				res = op1 * op2
			case "/":
				res = op1 / op2
			}

			stack = stack[0 : stackSize-2]
			stack = append(stack, res)

			stackSize--

			continue
		}

		converted, _ := strconv.Atoi(token)

		stack = append(stack, converted)
		stackSize++
	}

	return stack[0]
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}

	fmt.Println(evalRPN(tokens))
}
