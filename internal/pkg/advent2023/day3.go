package advent2023

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/marjamis/advent-of-code/pkg/helpers"
)

func findSurroundingNumbers(schematicMap [][]rune, currentPosition helpers.MatrixCoordinates) (foundPartNumbers []int) {
	surroundingLocations := helpers.FindSurroundingCoordinates(currentPosition, true)

	numberLocations := []helpers.MatrixCoordinates{}
	for _, location := range surroundingLocations {
		if helpers.IsLocationValid(schematicMap, location.Row, location.Col) {
			if unicode.IsDigit(schematicMap[location.Row][location.Col]) {
				numberLocations = append(numberLocations, helpers.MatrixCoordinates{
					Col: location.Col,
					Row: location.Row,
				})
			}
		}
	}

	validLocations := map[int]int{}
	for _, location := range numberLocations {
		start := location.Col
		end := location.Col + 1

		for i := location.Col - 1; i >= 0; i-- {
			if unicode.IsDigit(schematicMap[location.Row][i]) {
				start = i
			} else {
				break
			}
		}

		for i := location.Col + 1; i < len(schematicMap[location.Row]); i++ {
			if unicode.IsDigit(schematicMap[location.Row][i]) {
				end = i + 1
			} else {
				break
			}
		}

		number, err := strconv.Atoi(string(schematicMap[location.Row][start:end]))
		if err != nil {
			fmt.Println("Failed")
			return nil
		}
		validLocations[number] = 0
	}

	// Convert map keys to an array of those keys
	for location := range validLocations {
		foundPartNumbers = append(foundPartNumbers, location)
	}

	return foundPartNumbers
}

// Day3Part1 returns a sum of all the found parts
func Day3Part1(parts [][]rune) (sumOfParts int) {
	for rowIndex, row := range parts {
		for colIndex, char := range row {
			if (char >= 33 && char <= 47 && char != '.') || (char >= 58 && char <= 64) || (char >= 91 && char <= 96) {
				surroundingNumbers := findSurroundingNumbers(parts, helpers.MatrixCoordinates{
					Col: colIndex,
					Row: rowIndex,
				})

				for _, number := range surroundingNumbers {
					sumOfParts += number
				}
			}
		}
	}

	return
}

// Day3Part2 returns a sum of all the gears
func Day3Part2(parts [][]rune) (sumOfGears int) {
	for rowIndex, row := range parts {
		for colIndex, char := range row {
			if char == '*' {
				surroundingNumbers := findSurroundingNumbers(parts, helpers.MatrixCoordinates{
					Col: colIndex,
					Row: rowIndex,
				})

				if len(surroundingNumbers) == 2 {
					sumOfGears += (surroundingNumbers[0] * surroundingNumbers[1])
				}
			}
		}
	}

	return
}
