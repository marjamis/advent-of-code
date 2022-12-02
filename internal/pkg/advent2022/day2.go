package advent2022

import (
	"strings"

	"github.com/marjamis/advent-of-code/pkg/helpers"
)

var rockPaperScissorsMap = map[string]string{
	"A": "Rock",
	"X": "Rock",
	"B": "Paper",
	"Y": "Paper",
	"C": "Scissors",
	"Z": "Scissors",
}

var winningOrder = []interface{}{
	// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
	"Rock",     // 1 point for selection, References: A or X
	"Paper",    // 2 point for selection, References: B or Y
	"Scissors", // 3 point for selection, References: C or Z
}

func calculateRoundResults(opponent, yours string) (score int) {
	opponentIndex := helpers.GetArrayIndexForValue(winningOrder, opponent)
	yoursIndex := helpers.GetArrayIndexForValue(winningOrder, yours)

	// Loss score no points hence no action taken

	// Draw as the indexes are equal
	if opponentIndex == yoursIndex {
		score = 3
	}

	// Win as your index is the next level up OR -2 below (this accounts for Rock beating scissors and that index difference)
	if (yoursIndex == opponentIndex+1) || (yoursIndex == opponentIndex-2) {
		score = 6
	}

	return
}

func calculateFinalScore(input []string, decideHandShapes func(string, string) (string, string)) (score int) {
	for _, round := range input {
		split := strings.Split(round, " ")
		opponent, yours := decideHandShapes(split[0], split[1])

		// Result of the round
		score += calculateRoundResults(opponent, yours)

		// Points for hand shape the index of the hand shape +1 for the appropriate score
		score += helpers.GetArrayIndexForValue(winningOrder, yours) + 1
	}

	return
}

// Day2Part1 returns the points scored based on the exact actions provided
func Day2Part1(input []string) int {
	return calculateFinalScore(input, func(opponent, yours string) (string, string) {
		return rockPaperScissorsMap[opponent], rockPaperScissorsMap[yours]
	})
}

// Day2Part2 returns the points scored based on the plan provided
func Day2Part2(input []string) int {
	return calculateFinalScore(input, func(opponent, plan string) (string, string) {
		var yoursIndex int
		opp := rockPaperScissorsMap[opponent]

		switch plan {
		case "X":
			// Lose
			yoursIndex = helpers.GetValidAdjacentIndex(winningOrder, opp, false)
		case "Y":
			// Draw
			yoursIndex = helpers.GetArrayIndexForValue(winningOrder, opp)
		case "Z":
			// Win
			yoursIndex = helpers.GetValidAdjacentIndex(winningOrder, opp, true)
		}

		return opp, winningOrder[yoursIndex].(string)
	})
}
