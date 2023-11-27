package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	result := uint32(0)
	for i := 0; i < 32; i++ {
		result |= (num & 1) << (31 - i)
		num >>= 1
	}

	return result
}

func reverseBits1(num uint32) uint32 {
	result := uint32(0)
	for i := 0; i < 32; i++ {
		result <<= 1
		result |= num & 1
		num >>= 1
	}

	return result
}

func main() {
	n := uint32(964176192)
	fmt.Printf("%032b", reverseBits1(n))
}
