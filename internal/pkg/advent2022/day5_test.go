package advent2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day5ProvidedTestInput = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestDay5Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{
			day5ProvidedTestInput,
			"CMZ",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day5Part1(test.input))
	}
}

func TestDay5Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{
			day5ProvidedTestInput,
			"MCD",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day5Part2(test.input))
	}
}
