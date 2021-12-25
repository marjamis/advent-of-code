package advent2021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Part1(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			7,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day1Part1(test.input...))
	}
}

func TestDay1Part2(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			5,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day1Part2(test.input...))
	}
}
