package advent2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day3ProvidedTestInput = [][]rune{
	{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
	{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '3', '5', '.', '.', '6', '3', '3', '.'},
	{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
	{'6', '1', '7', '*', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '+', '.', '5', '8', '.'},
	{'.', '.', '5', '9', '2', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '7', '5', '5', '.'},
	{'.', '.', '.', '$', '.', '*', '.', '.', '.', '.'},
	{'.', '6', '6', '4', '.', '5', '9', '8', '.', '.'},
}

func TestDay3Part1(t *testing.T) {
	tests := []struct {
		input    [][]rune
		expected int
	}{
		{
			day3ProvidedTestInput,
			4361,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day3Part1(test.input))
	}
}

func TestDay3Part2(t *testing.T) {
	tests := []struct {
		input    [][]rune
		expected int
	}{
		{
			day3ProvidedTestInput,
			467835,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day3Part2(test.input))
	}
}
