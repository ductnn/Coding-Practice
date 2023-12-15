package main

import (
	"fmt"
)

func destCity(paths [][]string) string {
	s := map[string]bool{}
	for _, p := range paths {
		s[p[0]] = true
	}
	for _, p := range paths {
		if !s[p[1]] {
			return p[1]
		}
	}
	return ""
}

func main() {
	paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	fmt.Println(destCity(paths))
}
