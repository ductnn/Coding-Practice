// https://leetcode.com/problems/neighboring-bitwise-xor/

package main

func doesValidArrayExist(derived []int) bool {
	xor := 0
	for _, d := range derived {
		xor ^= d
	}

	return xor == 0
}

func main() {
	derived := []int{1, 1, 0}
	println(doesValidArrayExist(derived))
}
