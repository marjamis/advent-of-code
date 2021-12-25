package advent2021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeastFuelCheck(t *testing.T) {
	tests := []struct {
		currentLeastFuel int
		fuelToTest       int
		expected         int
	}{
		{
			10,
			11,
			10,
		},
		{
			11,
			10,
			10,
		},
		{
			10,
			10,
			10,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, leastFuelCheck(test.currentLeastFuel, test.fuelToTest))
	}
}
func TestInitialFuelCalculation(t *testing.T) {
	tests := []struct {
		input      []int
		toPosition int
		expected   int
	}{
		{
			[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			2,
			37,
		},
		{
			[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			1,
			41,
		},
		{
			[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			3,
			39,
		},
		{
			[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			10,
			71,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, initialFuelCalculation(test.toPosition, test.input))
	}
}

func TestUpdatedFuelCalculation(t *testing.T) {
	tests := []struct {
		input      []int
		toPosition int
		expected   int
	}{
		{
			[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			5,
			168,
		},
		{
			[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			2,
			206,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, updatedFuelCalculation(test.toPosition, test.input))
	}
}
