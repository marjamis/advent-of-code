package advent2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3Part1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			161,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day3Part1(test.input))
	}
}

func TestDay3Part2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			48,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day3Part2(test.input))
	}
}
