package advent2019

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6Part1(t *testing.T) {
	var day6Part1Tests = []struct {
		input    []string
		expected int
	}{
		//Provided Test
		{
			[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"},
			42,
		},
		//Personal Tests
		{
			[]string{"E)J", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "J)K", "COM)B", "B)C", "K)L"},
			42,
		},
	}

	for _, test := range day6Part1Tests {
		assert.Equal(t, test.expected, Day6Part1(test.input))
	}
}

func TestDay6Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		//Provided Test
		{
			[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"},
			4,
		},
		//Personal Tests
		{
			[]string{"E)J", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "J)K", "COM)B", "B)C", "K)L", "K)YOU", "I)SAN"},
			4,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day6Part2(test.input))
	}
}
