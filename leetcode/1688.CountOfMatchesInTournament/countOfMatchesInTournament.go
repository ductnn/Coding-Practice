package main

import (
	"fmt"
)

func numberOfMatches(n int) int {
	return n - 1
}

func numberOfMatches1(n int) int {
	var cnt int
	t := n
	for t > 1 {
		if t%2 == 0 {
			t /= 2
			cnt += t
		} else {
			cnt += (t - 1) / 2
			t = (t-1)/2 + 1
		}
	}
	return cnt
}

func numberOfMatches2(n int) int {
	cnt := 0
	t := n
	for t > 1 {
		if t%2 == 0 {
			t >>= 1
			cnt += t
		} else {
			cnt += (t - 1) >> 1
			t = (t-1)>>1 + 1
		}
	}
	return cnt
}

func numberOfMatches3(n int) int {
	sum := 0
	for n > 0 {
		sum = sum + n>>1
		if n%2 == 1 && n != 1 {
			n = n>>1 + 1
		} else {
			n >>= 1
		}
	}
	return sum
}

func main() {
	n := 10
	fmt.Println(numberOfMatches3(n))
}
