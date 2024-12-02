package advent2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day1Input = []string{
	"3   4",
	"4   3",
	"2   5",
	"1   3",
	"3   9",
	"3   3",
}

func TestDay1Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day1Input,
			11,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day1Part1(test.input))
	}
}

func TestDay1Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day1Input,
			31,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day1Part2(test.input))
	}
}
