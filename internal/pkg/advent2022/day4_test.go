package advent2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day4ProvidedTestInput = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestDay4Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day4ProvidedTestInput,
			2,
		},
		{
			[]string{
				"11-72,2-5",
			},
			0,
		},
		{
			[]string{
				"2-5,11-72",
			},
			0,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day4Part1(test.input))
	}
}

func TestDay4Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day4ProvidedTestInput,
			4,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day4Part2(test.input))
	}
}
