package advent2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day6ProvidedTestInput = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestDay6Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day6ProvidedTestInput,
			288,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day6Part1(test.input))
	}
}

func TestDay6Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day6ProvidedTestInput,
			71503,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day6Part2(test.input))
	}
}
