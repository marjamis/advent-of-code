package advent2023

import (
	"strings"
)

func isWinningNumberOnTicket(winningNumber string, ticketNumbers []string) bool {
	for _, ticketNumber := range ticketNumbers {
		if strings.Compare(winningNumber, ticketNumber) == 0 {
			return true
		}
	}

	return false
}

func getMatchesPerGame(cards []string) (matchesPerGame []int) {
	for _, card := range cards {
		parts := strings.Split(card, "|")
		winningNumbers := strings.Fields(parts[0][2:])
		ticketNumbers := strings.Fields(parts[1])

		count := 0
		for _, winningNumber := range winningNumbers {
			if isWinningNumberOnTicket(winningNumber, ticketNumbers) {
				count++
			}
		}

		matchesPerGame = append(matchesPerGame, count)
	}

	return
}

// Day4Part1 return the total number of point per winning card
func Day4Part1(cards []string) (totalPoints int) {
	matchesPerGame := getMatchesPerGame(cards)

	for _, matches := range matchesPerGame {
		if matches > 0 {
			points := 1
			for i := matches - 2; i >= 0; i-- {
				points *= 2
			}
			totalPoints += points
		}
	}

	return
}

// Day4Part2 returns the total number of scratchcards after the copies are made
func Day4Part2(cards []string) (totalScratchcards int) {
	matchesPerGame := getMatchesPerGame(cards)
	countPerCard := make([]int, len(matchesPerGame))

	for i := range countPerCard {
		countPerCard[i] = 1
	}

	for index, aheadCount := range matchesPerGame {
		for i := index + 1; i <= (index+aheadCount) && i < len(countPerCard); i++ {
			countPerCard[i] += (1 * countPerCard[index])
		}
	}

	for _, count := range countPerCard {
		totalScratchcards += count
	}

	return
}
