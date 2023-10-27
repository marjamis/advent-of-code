package advent2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day12ProvidedTestInput = [][]rune{
	{'S', 'a', 'b', 'q', 'p', 'o', 'n', 'm'},
	{'a', 'b', 'c', 'r', 'y', 'x', 'x', 'l'},
	{'a', 'c', 'c', 's', 'z', 'E', 'x', 'k'},
	{'a', 'c', 'c', 't', 'u', 'v', 'w', 'j'},
	{'a', 'b', 'd', 'e', 'f', 'g', 'h', 'i'},
}

func TestDay12Part1(t *testing.T) {
	tests := []struct {
		input    [][]rune
		goal     rune
		expected int
	}{
		{
			day12ProvidedTestInput,
			'E',
			31,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day12Part1(test.input, test.goal))
	}
}

func TestDay12Part2(t *testing.T) {
	tests := []struct {
		input    [][]rune
		goal     rune
		expected int
	}{
		{
			day12ProvidedTestInput,
			'E',
			29,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day12Part2(test.input, test.goal))
	}
}
