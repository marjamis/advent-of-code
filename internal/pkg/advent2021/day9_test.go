package advent2021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay9Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"2199943210",
				"3987894921",
				"9856789892",
				"8767896789",
				"9899965678",
			},
			15,
		},
		{
			[]string{
				"999",
				"999",
				"999",
			},
			0,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day9Part1(test.input))
	}
}

func TestFindBasinSize(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"952",
				"101",
				"234",
			},
			8,
		},
	}

	for _, test := range tests {
		heightMap := createHeightMap(test.input)

		basinSize, _ := heightMap.findBasinSize([]point{}, 1, 1)
		assert.Equal(t, test.expected, basinSize)
	}
}

func TestDay9Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"2199943210",
				"3987894921",
				"9856789892",
				"8767896789",
				"9899965678",
			},
			1134,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day9Part2(test.input))
	}
}
