package advent2021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay5Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"0,9 -> 5,9",
				"8,0 -> 0,8",
				"9,4 -> 3,4",
				"2,2 -> 2,1",
				"7,0 -> 7,4",
				"6,4 -> 2,0",
				"0,9 -> 2,9",
				"3,4 -> 1,4",
				"0,0 -> 8,8",
				"5,5 -> 8,2",
			},
			5,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day5Part1(test.input))
	}
}

func TestDay5Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		// Provided
		{
			[]string{
				"0,9 -> 5,9",
				"8,0 -> 0,8",
				"9,4 -> 3,4",
				"2,2 -> 2,1",
				"7,0 -> 7,4",
				"6,4 -> 2,0",
				"0,9 -> 2,9",
				"3,4 -> 1,4",
				"0,0 -> 8,8",
				"5,5 -> 8,2",
			},
			12,
		},
		// Checks diagnoal line draws in both directions
		{
			[]string{
				"1,1 -> 3,3",
				"9,7 -> 7,9",
			},
			0,
		},
		// Checks diagnoal line draws in both directions, with overlaps
		{
			[]string{
				"0,0 -> 3,3",
				"9,9 -> 3,3",
			},
			1,
		},
		// Checks diagnoal line draws in both directions, with all overlaps
		{
			[]string{
				"2,0 -> 0,2",
				"0,2 -> 2,0",
			},
			3,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day5Part2(test.input))
	}
}
