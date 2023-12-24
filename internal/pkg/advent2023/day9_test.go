package advent2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day9ProvidedTestInput = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}

func TestDay9Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day9ProvidedTestInput,
			114,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day9Part1(test.input))
	}
}

func TestExtrapolatedRight(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			day9ProvidedTestInput[0],
			18,
		},
		{
			day9ProvidedTestInput[1],
			28,
		},
		{
			day9ProvidedTestInput[2],
			68,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, extrapolateValue(test.input, true))
	}
}

func TestDay9Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day9ProvidedTestInput,
			2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day9Part2(test.input))
	}
}

func TestExtrapolatedLeft(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			day9ProvidedTestInput[0],
			-3,
		},
		{
			day9ProvidedTestInput[1],
			0,
		},
		{
			day9ProvidedTestInput[2],
			5,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, extrapolateValue(test.input, false))
	}
}
