package main

import (
	"fmt"
)

const HOME_TEAM_WON = 1
const POINTS = 3

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	currentBestTeam := ""

	// mapping score -> current team
	scores := map[string]int{currentBestTeam: 0}

	for i, competition := range competitions {
		homeTeam, awayTeam := competition[0], competition[1]
		result := results[i]

		winningTeam := awayTeam
		if result == HOME_TEAM_WON {
			winningTeam = homeTeam
		}

		scores[winningTeam] += POINTS

		if scores[winningTeam] > scores[currentBestTeam] {
			currentBestTeam = winningTeam
		}
	}

	return currentBestTeam
}

func main() {
	competitions := [][]string{
		{"HTML", "Java"},
		{"Java", "Golang"},
		{"Golang", "HTML"},
	}
	results := []int{0, 0, 1}

	fmt.Println(TournamentWinner(competitions, results))
}
