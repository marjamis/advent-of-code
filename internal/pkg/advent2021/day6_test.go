package advent2021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6Part1(t *testing.T) {
	tests := []struct {
		input          []int
		daysToSimulate int
		expected       int
	}{
		{
			[]int{3, 4, 3, 1, 2},
			18,
			26,
		},
		{
			[]int{3, 4, 3, 1, 2},
			80,
			5934,
		},
		{
			[]int{3},
			28,
			12,
		},
		{
			[]int{4},
			5,
			2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day6Part1(test.input, test.daysToSimulate))
	}
}

func TestDay6Part2(t *testing.T) {
	tests := []struct {
		input           []int
		daysToCalculate int
		expected        int
	}{
		{
			[]int{3, 4, 3, 1, 2},
			18,
			26,
		},
		{
			[]int{3, 4, 3, 1, 2},
			80,
			5934,
		},
		{
			[]int{3},
			28,
			12,
		},
		{
			[]int{3, 4, 3, 1, 2},
			256,
			26984457539,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day6Part2(test.input, test.daysToCalculate))
	}
}
