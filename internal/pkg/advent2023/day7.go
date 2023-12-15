package advent2023

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var orderedHandValues = map[string]int{
	"FiveOfAKind":  100,
	"FourOfAKind":  50,
	"FullHouse":    25,
	"ThreeOfAKind": 15,
	"TwoPair":      10,
	"OnePair":      5,
	"HighCard":     0,
}

type hand struct {
	cards string
	bid   int
}

func (h *hand) print() {
	fmt.Printf("Hand: %s | Bid: %d\n", h.cards, h.bid)
}

func calculateValueOfHand(cards string, jokersAreWild bool) (handValue int) {
	cardCounts := map[rune]int{}

	for _, card := range cards {
		cardCounts[card]++
	}

	// TODO revist this logic to be cleaner and flow better (probably by counting Jokers first and then
	// using that to determine the best hand to use them)
	switch len(cardCounts) {
	case 1:
		return orderedHandValues["FiveOfAKind"]
	case 2:
		four := false
		for _, value := range cardCounts {
			if value == 4 {
				four = true
			}
		}

		if countOfJokers, ok := cardCounts['J']; ok && jokersAreWild {
			if countOfJokers >= 1 {
				return orderedHandValues["FiveOfAKind"]
			}
		}

		if four {
			return orderedHandValues["FourOfAKind"]
		}

		return orderedHandValues["FullHouse"]
	case 3:
		three := false
		for _, value := range cardCounts {
			if value == 3 {
				three = true
			}
		}

		if countOfJokers, ok := cardCounts['J']; ok && jokersAreWild {
			if countOfJokers >= 2 {
				return orderedHandValues["FourOfAKind"]
			}

			if countOfJokers == 1 {
				for _, cardCount := range cardCounts {
					if cardCount == 3 {
						return orderedHandValues["FourOfAKind"]
					}
				}

				return orderedHandValues["FullHouse"]
			}
		}

		if three {
			return orderedHandValues["ThreeOfAKind"]
		}

		return orderedHandValues["TwoPair"]
	case 4:
		if countOfJokers, ok := cardCounts['J']; ok && jokersAreWild {
			if countOfJokers >= 1 {
				return orderedHandValues["ThreeOfAKind"]
			}
		}

		return orderedHandValues["OnePair"]
	case 5:
		if countOfJokers, ok := cardCounts['J']; ok && jokersAreWild {
			if countOfJokers >= 1 {
				return orderedHandValues["OnePair"]
			}
		}
	}

	return orderedHandValues["HighCard"]
}

func convertCardToValue(r rune, jokersAreWild bool) (value int) {
	switch r {
	case 'A':
		value = 14
	case 'K':
		value = 13
	case 'Q':
		value = 12
	case 'J':
		if jokersAreWild {
			value = 1
		} else {
			value = 11
		}
	case 'T':
		value = 10
	default:
		var err error
		value, err = strconv.Atoi(string(r))
		if err != nil {
			panic("Failed to convert to a string")
		}
	}

	return
}

func doesSecondBeatFirstInHighCard(first, second string, jokersAreWild bool) bool {
	for i := 0; i < len(first); i++ {
		firstsCard := convertCardToValue(rune(first[i]), jokersAreWild)
		secondsCard := convertCardToValue(rune(second[i]), jokersAreWild)

		if firstsCard > secondsCard {
			return false
		} else if firstsCard < secondsCard {
			return true
		}
	}

	return false
}

func calculateWinnings(handsData []string, jokersAreWild bool) (totalWinnings int) {
	hands := []hand{}

	for _, h := range handsData {
		split := strings.Split(h, " ")

		bid, err := strconv.Atoi(split[1])
		if err != nil {
			panic("Failed to convert to a string")
		}

		hands = append(hands, hand{
			cards: split[0],
			bid:   bid,
		})
	}

	sort.Slice(hands, func(first, second int) bool {
		firstsHandValue := calculateValueOfHand(hands[first].cards, jokersAreWild)
		secondsHandValue := calculateValueOfHand(hands[second].cards, jokersAreWild)

		if firstsHandValue == secondsHandValue {
			return doesSecondBeatFirstInHighCard(hands[first].cards, hands[second].cards, jokersAreWild)
		}

		if firstsHandValue < secondsHandValue {
			return true
		}

		return false
	})

	for i := range hands {
		totalWinnings += (hands[i].bid * (i + 1))
	}

	return
}

// Day7Part1 returns the total winnings with jokers being normal
func Day7Part1(handsData []string) (totalWinnings int) {
	return calculateWinnings(handsData, false)
}

// Day7Part2 returns the total winnings with jokers being wild
func Day7Part2(handsData []string) (totalWinnings int) {
	return calculateWinnings(handsData, true)
}
