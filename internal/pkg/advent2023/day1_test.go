package advent2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12Part1(t *testing.T) {
	tests := []struct {
		input    [][]rune
		expected int
	}{
		{
			[][]rune{
				{'1', 'a', 'b', 'c', '2'},
				{'p', 'q', 'r', '3', 's', 't', 'u', '8', 'v', 'w', 'x'},
				{'a', '1', 'b', '2', 'c', '3', 'd', '4', 'e', '5', 'f'},
				{'t', 'r', 'e', 'b', '7', 'u', 'c', 'h', 'e', 't'},
			},
			142,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day1Part1(test.input))
	}
}

func TestDay12Part2(t *testing.T) {
	tests := []struct {
		input    [][]rune
		expected int
	}{
		{
			[][]rune{
				{'t', 'w', 'o', '1', 'n', 'i', 'n', 'e'},
				{'e', 'i', 'g', 'h', 't', 'w', 'o', 't', 'h', 'r', 'e', 'e'},
				{'a', 'b', 'c', 'o', 'n', 'e', '2', 't', 'h', 'r', 'e', 'e', 'x', 'y', 'z'},
				{'x', 't', 'w', 'o', 'n', 'e', '3', 'f', 'o', 'u', 'r'},
				{'4', 'n', 'i', 'n', 'e', 'e', 'i', 'g', 'h', 't', 's', 'e', 'v', 'e', 'n', '2'},
				{'z', 'o', 'n', 'e', 'i', 'g', 'h', 't', '2', '3', '4'},
				{'7', 'p', 'q', 'r', 's', 't', 's', 'i', 'x', 't', 'e', 'e', 'n'},
			},
			281,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day1Part2(test.input))
	}
}
