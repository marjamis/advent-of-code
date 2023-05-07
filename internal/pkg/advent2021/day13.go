package advent2021

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type paperPoint struct {
	col, row int
}

type paperPoints []paperPoint

// Paper is the representation of the paper after each fold
type Paper [][]string

func splitInstructionsAndCoordinates(rawData []string) (instructions []string, coordinates []string) {
	for _, row := range rawData {
		if strings.Contains(row, "fold along") {
			instructions = append(instructions, row)
		} else if strings.Contains(row, ",") {
			coordinates = append(coordinates, row)
		}
	}

	return
}

func getCoordinates(coordinates []string) paperPoints {
	paperPoints := make([]paperPoint, len(coordinates))

	for index, coordinate := range coordinates {
		c := strings.Split(coordinate, ",")
		col, err := strconv.Atoi(c[0])
		if err != nil {
			log.Fatal(err)
		}

		row, err := strconv.Atoi(c[1])
		if err != nil {
			log.Fatal(err)
		}

		paperPoints[index] = paperPoint{
			col: col,
			row: row,
		}
	}

	return paperPoints
}

func getPaperLengths(paperPoints []paperPoint) (colSize, rowSize int) {
	for _, paperPoint := range paperPoints {
		if paperPoint.col > colSize {
			colSize = paperPoint.col
		}

		if paperPoint.row > rowSize {
			rowSize = paperPoint.row
		}
	}

	// The +1's is because the lengths are one larger than the biggest value
	return colSize + 1, rowSize + 1
}

func createBlankPaper(colSize, rowSize int) Paper {
	paper := make([][]string, rowSize)
	for row := 0; row < rowSize; row++ {
		paper[row] = make([]string, colSize)
		for col := 0; col < colSize; col++ {
			paper[row][col] = "."
		}
	}

	return paper
}

func (points paperPoints) addPointsToPaper(colSize, rowSize int) Paper {
	paper := createBlankPaper(colSize, rowSize)

	for _, coordinate := range points {
		paper[coordinate.row][coordinate.col] = "#"
	}

	return paper
}

func indexAfterFold(initialIndex, foldAtLine, maxIndex int) int {
	if initialIndex == foldAtLine {
		log.Fatal("error: existing index on fold line")
	}

	if initialIndex < foldAtLine {
		return initialIndex
	}

	calculatedIndex := maxIndex - initialIndex
	if calculatedIndex < 0 || calculatedIndex >= foldAtLine {
		log.Fatal("error: calculated index outside the expected bounds")
	}

	return calculatedIndex
}

func createFoldedPaper(instructions []string, coordinates []string) Paper {
	coords := getCoordinates(coordinates)
	// Initial "guessimate" of paper length if not overwritten below
	colLength, rowLength := getPaperLengths(coords)

	for _, instruction := range instructions {
		lineString := strings.Split(instruction, "=")[1]
		foldAtIndex, err := strconv.Atoi(lineString)

		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(instruction, "x") {
			for i := range coords {
				coords[i].col = indexAfterFold(coords[i].col, foldAtIndex, (foldAtIndex * 2))
			}

			colLength = foldAtIndex
		} else {
			for i := range coords {
				coords[i].row = indexAfterFold(coords[i].row, foldAtIndex, (foldAtIndex * 2))
			}

			rowLength = foldAtIndex
		}
	}

	paper := coords.addPointsToPaper(colLength, rowLength)

	return paper
}

func (paper Paper) countDots() (count int) {
	for _, row := range paper {
		for _, col := range row {
			if col == "#" {
				count++
			}
		}
	}

	return
}

func (paper Paper) display() string {
	formatted := "----\n"
	for _, row := range paper {
		for _, col := range row {
			if col == "." {
				formatted += fmt.Sprintf(" ")
			} else {
				formatted += fmt.Sprintf(col)
			}
		}
		formatted += "\n"
	}
	formatted += "----\n"

	return formatted
}

// Day13Part1 returns the calculated version for the number of dots visible after the first fold
func Day13Part1(rawData []string) int {
	instructions, coordinates := splitInstructionsAndCoordinates(rawData)

	paper := createFoldedPaper([]string{
		instructions[0],
	}, coordinates)

	return paper.countDots()
}

// Day13Part2 displays the calculated version of the 8 letters (in ASCII picture format) which is the code for the thermal camera
func Day13Part2(rawData []string) string {
	instructions, coordinates := splitInstructionsAndCoordinates(rawData)

	paper := createFoldedPaper(instructions, coordinates)

	return paper.display()
}
