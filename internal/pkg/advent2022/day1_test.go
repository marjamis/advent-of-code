package advent2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day1ProvidedTestInput = []string{
	"1000",
	"2000",
	"3000",
	"",
	"4000",
	"",
	"5000",
	"6000",
	"",
	"7000",
	"8000",
	"9000",
	"",
	"10000",
}

func TestDay1Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day1ProvidedTestInput,
			24000,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day1Part1(test.input))
	}
}

func TestDay1Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day1ProvidedTestInput,
			45000,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day1Part2(test.input))
	}
}
