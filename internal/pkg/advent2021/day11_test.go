package advent2021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrease(t *testing.T) {
	t.Run("Basic test", func(t *testing.T) {
		assert.Equal(t, 0, increase(-1))
		assert.Equal(t, 1, increase(0))
		assert.Equal(t, 2, increase(1))
	})

	t.Run("Mapping test", func(t *testing.T) {
		oem := OctopusEnergyMap{
			{-2, -1, 0},
			{1, 2, 3},
			{4, 5, 6},
		}
		expected := OctopusEnergyMap{
			{-1, 0, 1},
			{2, 3, 4},
			{5, 6, 7},
		}
		oem.mapping(increase)
		assert.ElementsMatch(t, expected, oem)
	})
}

func TestSetToZero(t *testing.T) {
	t.Run("Basic test", func(t *testing.T) {
		assert.Equal(t, 8, setToZero(8))
		assert.Equal(t, 9, setToZero(9))
		assert.Equal(t, 0, setToZero(10))
	})

	t.Run("Mapping test", func(t *testing.T) {
		oem := OctopusEnergyMap{
			{-2, -1, 0},
			{9, 9, 9},
			{10, 11, 12},
		}
		expected := OctopusEnergyMap{
			{-2, -1, 0},
			{9, 9, 9},
			{0, 0, 0},
		}
		oem.mapping(setToZero)
		assert.ElementsMatch(t, expected, oem)
	})
}

func TestStepOutput(t *testing.T) {
	tests := []struct {
		input       OctopusEnergyMap
		expectedOEM OctopusEnergyMap
		expected    int
	}{
		{
			[][]int{
				{9, 8},
				{7, 3},
			},
			[][]int{
				{0, 0},
				{0, 7},
			},
			3,
		},
		{
			[][]int{
				{1, 1, 1, 1, 1},
				{1, 9, 9, 9, 1},
				{1, 9, 1, 9, 1},
				{1, 9, 9, 9, 1},
				{1, 1, 1, 1, 1},
			},
			[][]int{
				{3, 4, 5, 4, 3},
				{4, 0, 0, 0, 4},
				{5, 0, 0, 0, 5},
				{4, 0, 0, 0, 4},
				{3, 4, 5, 4, 3},
			},
			9,
		},
		{
			[][]int{
				{3, 4, 5, 4, 3},
				{4, 0, 0, 0, 4},
				{5, 0, 0, 0, 5},
				{4, 0, 0, 0, 4},
				{3, 4, 5, 4, 3},
			},
			[][]int{
				{4, 5, 6, 5, 4},
				{5, 1, 1, 1, 5},
				{6, 1, 1, 1, 6},
				{5, 1, 1, 1, 5},
				{4, 5, 6, 5, 4},
			},
			0,
		},
	}

	for _, test := range tests {
		oem := test.input
		flashCount := oem.step()
		assert.Equal(t, test.expectedOEM, oem)
		assert.Equal(t, test.expected, flashCount)
	}
}

func TestDay11Part1(t *testing.T) {
	tests := []struct {
		input    OctopusEnergyMap
		expected int
	}{
		{
			[][]int{
				{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
				{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
				{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
				{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
				{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
				{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
				{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
				{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
				{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
				{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
			},
			1656,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day11Part1(test.input))
	}
}

func TestIsSynchronised(t *testing.T) {
	falseSynchornised := OctopusEnergyMap{
		{-2, -1, 0},
		{1, 2, 3},
		{4, 5, 6},
	}
	assert.False(t, falseSynchornised.isSynchonised())

	trueSynchornised := OctopusEnergyMap{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	assert.True(t, trueSynchornised.isSynchonised())
}

func TestDay11Part2(t *testing.T) {
	tests := []struct {
		input    OctopusEnergyMap
		expected int
	}{
		{
			[][]int{
				{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
				{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
				{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
				{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
				{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
				{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
				{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
				{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
				{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
				{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
			},
			195,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day11Part2(test.input))
	}
}
