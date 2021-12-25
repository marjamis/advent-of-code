package advent2021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2",
			},
			150,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day2Part1(test.input...))
	}
}

func TestDay2Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2",
			},
			900,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day2Part2(test.input...))
	}
}
