package advent2021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCoordinates(t *testing.T) {
	tests := []struct {
		input    []string
		expected []paperPoint
	}{
		{
			[]string{
				"6,10",
				"0,14",
				"9,10",
				"0,3",
				"10,4",
			},
			[]paperPoint{
				{row: 10, col: 6},
				{row: 14, col: 0},
				{row: 10, col: 9},
				{row: 3, col: 0},
				{row: 4, col: 10},
			},
		},
	}

	for _, test := range tests {
		assert.ElementsMatch(t, test.expected, getCoordinates(test.input))
	}
}

func TestGetPaperSize(t *testing.T) {
	tests := []struct {
		input           []paperPoint
		expectedColSize int
		expectedRowSize int
	}{
		{
			[]paperPoint{
				{col: 6, row: 10},
				{col: 0, row: 14},
				{col: 9, row: 10},
				{col: 0, row: 3},
				{col: 10, row: 4},
			},
			11,
			15,
		},
	}

	for _, test := range tests {
		colSize, rowSize := getPaperLengths(test.input)
		assert.Equal(t, test.expectedColSize, colSize)
		assert.Equal(t, test.expectedRowSize, rowSize)
	}
}

func TestCreateBlankPaper(t *testing.T) {
	tests := []struct {
		colSize int
		rowSize int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{121, 50},
		{50, 121},
	}

	for _, test := range tests {
		paper := createBlankPaper(test.colSize, test.rowSize)

		assert.Equal(t, test.rowSize, len(paper))
		for _, row := range paper {
			assert.Equal(t, test.colSize, len(row))
		}
	}
}

func TestSplitCoordinatesAndInstructions(t *testing.T) {
	data := []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
		"",
		"fold along y=7",
		"fold along x=5",
	}
	expectedCoordinates := []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
	}
	expectedInstructions := []string{
		"fold along y=7",
		"fold along x=5",
	}
	inst, coord := splitInstructionsAndCoordinates(data)
	assert.ElementsMatch(t, expectedInstructions, inst)
	assert.ElementsMatch(t, expectedCoordinates, coord)
}

func TestPositionAfterFold(t *testing.T) {
	tests := []struct {
		index      int
		foldAtLine int
		maxIndex   int
		expected   int
	}{
		{0, 1, 2, 0},
		// {1, 1, 2, -1},
		{2, 1, 2, 0},
		{10, 7, 14, 4},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, indexAfterFold(test.index, test.foldAtLine, test.maxIndex))
	}
}

func TestProvidedExample(t *testing.T) {
	data := []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
		"",
		"fold along y=7",
		"fold along x=5",
	}
	expected := Paper{
		{"#", "#", "#", "#", "#"},
		{"#", ".", ".", ".", "#"},
		{"#", ".", ".", ".", "#"},
		{"#", ".", ".", ".", "#"},
		{"#", "#", "#", "#", "#"},
		{".", ".", ".", ".", "."},
		{".", ".", ".", ".", "."},
	}

	t.Run("Calculated", func(t *testing.T) {
		instructions, coordinates := splitInstructionsAndCoordinates(data)
		finalCoordinates := createFoldedPaper(instructions, coordinates)

		assert.ElementsMatch(t, expected, finalCoordinates)
	})
}
