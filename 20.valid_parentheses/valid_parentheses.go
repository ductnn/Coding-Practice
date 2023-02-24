package main

func isValid(s string) bool {
	stack := []rune{}
	if len(stack)%2 != 0 {
		return false
	}

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else if v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else if v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else if v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
}
