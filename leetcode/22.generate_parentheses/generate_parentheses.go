package main

import "fmt"

func generateParenthesis(n int) []string {
	result := make([]string, 0)

	generate("", n, 0, 0, &result)
	return result
}

func generate(current string, n, open, close int, result *[]string) {
	if close == n {
		*result = append(*result, current)
		return
	} else {
		if open < n {
			generate(current+"(", n, open+1, close, result)
		}

		if open > close {
			generate(current+")", n, open, close+1, result)
		}
	}
}

func main() {
	n := 20
	fmt.Println(generateParenthesis(n))
}
