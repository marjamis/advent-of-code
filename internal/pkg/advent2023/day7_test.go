package advent2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day7ProvidedTestInput = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestDay7Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day7ProvidedTestInput,
			6440,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day7Part1(test.input))
	}
}

func TestDay7Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day7ProvidedTestInput,
			5905,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day7Part2(test.input))
	}
}

func TestCalculateHandValue(t *testing.T) {
	tests := []struct {
		input         string
		jokersAreWild bool
		expected      int
	}{
		{"AAAAA", false, orderedHandValues["FiveOfAKind"]},
		{"AAAAK", false, orderedHandValues["FourOfAKind"]},
		{"AAAAJ", false, orderedHandValues["FourOfAKind"]},
		{"AAAKK", false, orderedHandValues["FullHouse"]},
		{"AAAJJ", false, orderedHandValues["FullHouse"]},
		{"AAJJJ", false, orderedHandValues["FullHouse"]},
		{"AJJJJ", false, orderedHandValues["FourOfAKind"]},
		{"AQQQQ", false, orderedHandValues["FourOfAKind"]},
		{"AAAJK", false, orderedHandValues["ThreeOfAKind"]},
		{"AAJJK", false, orderedHandValues["TwoPair"]},
		{"AAJQK", false, orderedHandValues["OnePair"]},
		{"AJQK9", false, orderedHandValues["HighCard"]},
		{"AAAAA", true, orderedHandValues["FiveOfAKind"]},
		{"AAAAK", true, orderedHandValues["FourOfAKind"]},
		{"AAAAJ", true, orderedHandValues["FiveOfAKind"]},
		{"AAAKK", true, orderedHandValues["FullHouse"]},
		{"AAAJJ", true, orderedHandValues["FiveOfAKind"]},
		{"AAAJQ", true, orderedHandValues["FourOfAKind"]},
		{"AAQJJ", true, orderedHandValues["FourOfAKind"]},
		{"AAATK", true, orderedHandValues["ThreeOfAKind"]},
		{"AATTK", true, orderedHandValues["TwoPair"]},
		{"AAT9K", true, orderedHandValues["OnePair"]},
		{"AT4KQ", true, orderedHandValues["HighCard"]},
		{"ATJKQ", true, orderedHandValues["OnePair"]},
		{"ATJJQ", true, orderedHandValues["ThreeOfAKind"]},
		{"ATJJJ", true, orderedHandValues["FourOfAKind"]},
		{"ATTJJ", true, orderedHandValues["FourOfAKind"]},
		{"ATTAJ", true, orderedHandValues["FullHouse"]},
		{"TTJJJ", true, orderedHandValues["FiveOfAKind"]},
		{"TTTJJ", true, orderedHandValues["FiveOfAKind"]},
		{"TTTTJ", true, orderedHandValues["FiveOfAKind"]},
		{"TTTTQ", true, orderedHandValues["FourOfAKind"]},
		{"AAAAK", true, orderedHandValues["FourOfAKind"]},
		{"AAAAJ", true, orderedHandValues["FiveOfAKind"]},
		{"AAAKK", true, orderedHandValues["FullHouse"]},
		{"AAAJJ", true, orderedHandValues["FiveOfAKind"]},
		{"AAJJJ", true, orderedHandValues["FiveOfAKind"]},
		{"AJJJJ", true, orderedHandValues["FiveOfAKind"]},
		{"AQQQQ", true, orderedHandValues["FourOfAKind"]},
		{"AKQJT", true, orderedHandValues["OnePair"]},
		{"AKQTT", true, orderedHandValues["OnePair"]},
		{"AKJJT", true, orderedHandValues["ThreeOfAKind"]},
		{"AATJK", true, orderedHandValues["ThreeOfAKind"]},
		{"AKJTT", true, orderedHandValues["ThreeOfAKind"]},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, calculateValueOfHand(test.input, test.jokersAreWild), fmt.Sprintf("%s - %v", test.input, test.jokersAreWild))
	}
}

func TestHighCardNewWins(t *testing.T) {
	tests := []struct {
		inputA        string
		inputB        string
		jokersAreWild bool
		expected      bool
	}{
		{"AJQK9", "AJQK9", false, false},
		{"AJQK9", "9KQJA", false, false},
		{"9KQJA", "AJQK9", false, true},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, doesSecondBeatFirstInHighCard(test.inputA, test.inputB, false))
	}
}

func TestConvertRuneToScale(t *testing.T) {
	tests := []struct {
		input         rune
		jokersAreWild bool
		expected      int
	}{
		{'A', false, 14},
		{'K', false, 13},
		{'Q', false, 12},
		{'J', false, 11},
		{'T', false, 10},
		{'9', false, 9},
		{'8', false, 8},
		{'7', false, 7},
		{'6', false, 6},
		{'5', false, 5},
		{'4', false, 4},
		{'3', false, 3},
		{'2', false, 2},
		{'J', true, 1},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, convertCardToValue(test.input, test.jokersAreWild))
	}
}
