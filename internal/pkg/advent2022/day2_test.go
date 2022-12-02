package advent2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day2ProvidedTestInput = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestDay2Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day2ProvidedTestInput,
			15,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day2Part1(test.input))
	}
}

func TestDay2Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day2ProvidedTestInput,
			12,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day2Part2(test.input))
	}
}
