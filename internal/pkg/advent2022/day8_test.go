package advent2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day8ProvidedTestInput = [][]int{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func TestDay8Part1(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{
			day8ProvidedTestInput,
			21,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day8Part1(test.input))
	}
}

func TestDay8Part2(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{
			day8ProvidedTestInput,
			8,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day8Part2(test.input))
	}
}
