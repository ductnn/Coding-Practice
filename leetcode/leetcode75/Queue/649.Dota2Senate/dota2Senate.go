package main

import (
	"fmt"
)

func predictPartyVictory(senate string) string {
	n := len(senate)
	queueR := []int{}
	queueD := []int{}
	for i, c := range senate {
		if c == 'R' {
			queueR = append(queueR, i)
		} else {
			queueD = append(queueD, i)
		}
	}
	for len(queueR) > 0 && len(queueD) > 0 {
		r, d := queueR[0], queueD[0]
		queueR, queueD = queueR[1:], queueD[1:]
		if r < d {
			queueR = append(queueR, r+n)
		} else {
			queueD = append(queueD, d+n)
		}
	}
	if len(queueR) > 0 {
		return "Radiant"
	}
	return "Dire"
}

func main() {
	senate := "RD"
	fmt.Println(predictPartyVictory(senate))
}
