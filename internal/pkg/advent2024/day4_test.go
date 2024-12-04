package advent2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day4Input = [][]rune{
	{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
	{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
	{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
	{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
	{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
	{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
	{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
	{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
	{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
	{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
}

func TestDay4Part1(t *testing.T) {
	tests := []struct {
		input    [][]rune
		expected int
	}{
		{
			day4Input,
			18,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day4Part1(test.input))
	}
}

func TestDay4Part2(t *testing.T) {
	tests := []struct {
		input    [][]rune
		expected int
	}{
		{
			day4Input,
			9,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day4Part2(test.input))
	}
}
