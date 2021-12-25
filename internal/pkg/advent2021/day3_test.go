package advent2021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			198,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day3Part1(test.input))
	}
}

func TestDay3Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			230,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day3Part2(test.input))
	}
}
