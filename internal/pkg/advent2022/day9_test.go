package advent2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day9ProvidedTestInput = []string{
	"R 4",
	"U 4",
	"L 3",
	"D 1",
	"R 4",
	"D 1",
	"L 5",
	"R 2",
}

func TestDay9Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day9ProvidedTestInput,
			13,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day9Part1(test.input))
	}
}

func TestDay9Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day9ProvidedTestInput,
			1,
		},
		{
			[]string{
				"R 5",
				"U 8",
				"L 8",
				"D 3",
				"R 17",
				"D 10",
				"L 25",
				"U 20",
			},
			36,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day9Part2(test.input))
	}
}
