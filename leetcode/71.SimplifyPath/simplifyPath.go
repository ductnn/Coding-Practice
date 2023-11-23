package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := []string{}

	for _, s := range strings.Split(path, "/") {
		if s == "" || s == "." {
			continue
		}
		if s == ".." {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack = append(stack, s)
		}
	}

	return "/" + strings.Join(stack, "/")
}

func main() {
	path := "/home//foo/"
	fmt.Println(simplifyPath(path))
}
