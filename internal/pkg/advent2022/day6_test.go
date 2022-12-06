package advent2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day6ProvidedTestInput = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	"nppdvjthqldpwncqszvftbrmjlhg",
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

func TestDay6Part1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			day6ProvidedTestInput[0],
			7,
		},
		{
			day6ProvidedTestInput[1],
			5,
		},
		{
			day6ProvidedTestInput[2],
			6,
		},
		{
			day6ProvidedTestInput[3],
			10,
		},
		{
			day6ProvidedTestInput[4],
			11,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day6Part1(test.input))
	}
}

func TestDay6Part2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			day6ProvidedTestInput[0],
			19,
		},
		{
			day6ProvidedTestInput[1],
			23,
		},
		{
			day6ProvidedTestInput[2],
			23,
		},
		{
			day6ProvidedTestInput[3],
			29,
		},
		{
			day6ProvidedTestInput[4],
			26,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day6Part2(test.input))
	}
}
