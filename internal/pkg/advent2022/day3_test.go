package advent2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day3ProvidedTestInput = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestDay3Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day3ProvidedTestInput,
			157,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day3Part1(test.input))
	}
}

func TestDay3Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day3ProvidedTestInput,
			70,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day3Part2(test.input))
	}
}
