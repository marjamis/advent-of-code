package advent2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// var day8ProvidedTestInput =

func TestDay8Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"RL",
				"",
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			2,
		},
		{
			[]string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			6,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day8Part1(test.input))
	}
}

func TestDay8Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"LR",
				"",
				"11A = (11B, XXX)",
				"11B = (XXX, 11Z)",
				"11Z = (11B, XXX)",
				"22A = (22B, XXX)",
				"22B = (22C, 22C)",
				"22C = (22Z, 22Z)",
				"22Z = (22B, 22B)",
				"XXX = (XXX, XXX)",
			},
			6,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day8Part2(test.input))
	}
}
