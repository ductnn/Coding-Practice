package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	result := ""

	// 2 pointers
	i := len(a) - 1
	j := len(b) - 1

	var carry byte = 0

	for i >= 0 || j >= 0 || carry == 1 {
		if i >= 0 {
			carry += a[i] - '0'
		}
		if j >= 0 {
			carry += b[j] - '0'
		}

		result = string(carry%2+'0') + result
		carry = carry / 2
		i--
		j--
	}

	return result
}

func addBinary1(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	result := []byte{}

	for carry := 0; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		if i >= 0 {
			carry += int(a[i] - '0')
		}
		if j >= 0 {
			carry += int(b[j] - '0')
		}
		result = append(result, byte(carry%2+'0'))
		carry /= 2
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func main() {
	a := "11"
	b := "1"

	fmt.Println(addBinary(a, b))
}
