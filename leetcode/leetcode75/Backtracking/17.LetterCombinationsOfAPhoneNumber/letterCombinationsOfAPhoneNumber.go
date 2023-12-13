package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}

	phoneMap := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	result = append(result, "")
	for _, v := range digits {
		s := phoneMap[v-'2']
		temp := []string{}
		for _, i := range result {
			for _, j := range s {
				temp = append(temp, i+string(j))
			}
		}
		result = temp
	}

	return result
}

func main() {
	digits := "25"
	fmt.Println(letterCombinations(digits))
}
