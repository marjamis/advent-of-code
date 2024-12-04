package advent2024

import (
	log "github.com/sirupsen/logrus"

	"github.com/marjamis/advent-of-code/pkg/helpers"
)

type arrayPositionStep struct {
	row int
	col int
}

type arrayPositionSteps []arrayPositionStep

// TODO alias these directions so can be referenced rather than manual?
func findXmas(input [][]rune, steps arrayPositionSteps) bool {
	if helpers.IsArrayLocationValid(input, steps[0].row, steps[0].col) && input[steps[0].row][steps[0].col] == 'M' {
		if helpers.IsArrayLocationValid(input, steps[1].row, steps[1].col) && input[steps[1].row][steps[1].col] == 'A' {
			if helpers.IsArrayLocationValid(input, steps[2].row, steps[2].col) && input[steps[2].row][steps[2].col] == 'S' {
				return true
			}
		}
	}

	return false
}

// Day4Part1 return number of XMAS's in the input
func Day4Part1(input [][]rune) (count int) {
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] == 'X' {
				log.Debugf("Col %d : Row %d\n", col, row)

				directions := []arrayPositionSteps{
					arrayPositionSteps{
						{row: row - 1, col: col},
						{row: row - 2, col: col},
						{row: row - 3, col: col},
					},
					arrayPositionSteps{
						{row: row + 1, col: col},
						{row: row + 2, col: col},
						{row: row + 3, col: col},
					},
					arrayPositionSteps{
						{row: row, col: col - 1},
						{row: row, col: col - 2},
						{row: row, col: col - 3},
					},
					arrayPositionSteps{
						{row: row, col: col + 1},
						{row: row, col: col + 2},
						{row: row, col: col + 3},
					},
					arrayPositionSteps{
						{row: row - 1, col: col - 1},
						{row: row - 2, col: col - 2},
						{row: row - 3, col: col - 3},
					},
					arrayPositionSteps{
						{row: row - 1, col: col + 1},
						{row: row - 2, col: col + 2},
						{row: row - 3, col: col + 3},
					},
					arrayPositionSteps{
						{row: row + 1, col: col - 1},
						{row: row + 2, col: col - 2},
						{row: row + 3, col: col - 3},
					},
					arrayPositionSteps{
						{row: row + 1, col: col + 1},
						{row: row + 2, col: col + 2},
						{row: row + 3, col: col + 3},
					},
				}

				for _, direction := range directions {
					if findXmas(input, direction) {
						count++
					}
				}
			}

		}
	}

	return count
}

// Day4Part2 return the count of X-MAS' in the input
func Day4Part2(input [][]rune) (count int) {
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] == 'A' {
				log.Debugf("Col %d : Row %d\n", col, row)

				var upLeft, upRight, downLeft, downRight rune
				if helpers.IsArrayLocationValid(input, row-1, col-1) {
					upLeft = input[row-1][col-1]
				}
				if helpers.IsArrayLocationValid(input, row-1, col+1) {
					upRight = input[row-1][col+1]
				}
				if helpers.IsArrayLocationValid(input, row+1, col-1) {
					downLeft = input[row+1][col-1]
				}
				if helpers.IsArrayLocationValid(input, row+1, col+1) {
					downRight = input[row+1][col+1]
				}

				if ((upLeft == 'M' && downRight == 'S') || (upLeft == 'S' && downRight == 'M')) &&
					((downLeft == 'M' && upRight == 'S') || (downLeft == 'S' && upRight == 'M')) {
					count++
				}
			}
		}
	}

	return
}
