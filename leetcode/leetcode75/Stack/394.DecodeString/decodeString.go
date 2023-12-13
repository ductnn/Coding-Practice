package main

import (
	"fmt"
)

func decodeString(s string) string {
	s1, s2 := []int{}, []string{}
	num, result := 0, ""

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			s1 = append(s1, num)
			s2 = append(s2, result)
			num = 0
			result = ""
		} else if s[i] == ']' {
			n := s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
			temp := ""
			for i := 0; i < n; i++ {
				temp += result
			}
			result = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
			result += temp
		} else {
			result += string(s[i])
		}
	}
	return result
}

func main() {
	s := "3[a2[c]]"
	fmt.Println(decodeString(s))
}
