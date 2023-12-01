package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIPAddresses(s string) []string {
	result := []string{}
	segments := make([]string, 4)

	var backtrack func(start, segmentNum int)
	backtrack = func(start, segmentNum int) {
		// If we've got 4 segments and used all characters in the string, it's a valid IP address
		if segmentNum == 4 && start == len(s) {
			ip := strings.Join(segments, ".")
			result = append(result, ip)
			return
		}

		// If we've got 4 segments but not all characters, it's not valid
		if segmentNum == 4 || start == len(s) {
			return
		}

		if s[start] == '0' {
			segments[segmentNum] = "0"
			backtrack(start+1, segmentNum+1)
		}

		addr := 0
		for i := start; i < len(s); i++ {
			addr = addr*10 + int(s[i]-'0')
			if addr > 0 && addr <= 255 {
				segments[segmentNum] = strconv.Itoa(addr)
				backtrack(i+1, segmentNum+1)
			} else {
				break
			}
		}
	}

	backtrack(0, 0)
	return result
}

func main() {
	ip := "25525511135"
	fmt.Println(restoreIPAddresses(ip))
}
