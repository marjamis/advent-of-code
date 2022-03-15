package advent2021

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type paperPoint struct {
	row, col int
}

type paperPoints []paperPoint

// Paper is the representation of the paper after each fold
type Paper [][]string

func getCoordinates(coordinates []string) []paperPoint {
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
			row: row,
			col: col,
		}
	}

	return paperPoints
}

func getPaperSize(paperPoints []paperPoint) (colSize, rowSize int) {
	for _, paperPoint := range paperPoints {
		if paperPoint.col > colSize {
			colSize = paperPoint.col
		}

		if paperPoint.row > rowSize {
			rowSize = paperPoint.row
		}
	}

	// The +1's is because the largest point found doesn't take into account the 0 index
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

func loadPaper(stringCoordinates []string) Paper {
	coordinates := getCoordinates(stringCoordinates)
	colSize, rowSize := getPaperSize(coordinates)
	paper := createBlankPaper(colSize, rowSize)

	for _, coordinate := range coordinates {
		paper[coordinate.row][coordinate.col] = "#"
	}

	return paper
}

func (paper Paper) foldOnHorizontalLine(foldAtLine int) Paper {
	newPaper := createBlankPaper(len(paper[0]), foldAtLine)

	for _, col := range paper[foldAtLine] {
		if col == "#" {
			log.Fatal(fmt.Sprintf("Horizontal fold at line issue: %d\n", foldAtLine))
		}
	}

	for inner, outer := 0, len(paper)-1; inner < foldAtLine; inner, outer = inner+1, outer-1 {
		for col := 0; col < len(paper[0]); col++ {
			if paper[inner][col] == "#" || paper[outer][col] == "#" {
				newPaper[inner][col] = "#"
			}
		}
	}

	return newPaper
}

func (paper Paper) foldOnVerticalLine(foldAtLine int) Paper {
	newPaper := createBlankPaper(foldAtLine, len(paper))

	for row := range paper {
		if paper[row][foldAtLine] == "#" {
			log.Fatal(fmt.Sprintf("Vertical fold at line issue: %d", foldAtLine))
		}
	}

	for row := 0; row < len(paper); row++ {
		for inner, outer := 0, len(paper[row])-1; inner < foldAtLine; inner, outer = inner+1, outer-1 {
			if paper[row][inner] == "#" || paper[row][outer] == "#" {
				newPaper[row][inner] = "#"
			}
		}
	}

	return newPaper
}

func splitCoordinatesAndInstructions(rawData []string) (instructions []string, coordinates []string) {
	for _, row := range rawData {
		if strings.Contains(row, "fold along") {
			instructions = append(instructions, row)
		} else if strings.Contains(row, ",") {
			coordinates = append(coordinates, row)
		}
	}

	return
}

func (paper Paper) fold(instruction string) Paper {
	lineString := strings.Split(instruction, "=")[1]
	line, err := strconv.Atoi(lineString)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(instruction, "x") {
		paper = paper.foldOnVerticalLine(line)
	} else {
		paper = paper.foldOnHorizontalLine(line)
	}

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

func (paper Paper) print() {
	for _, row := range paper {
		fmt.Println(row)
	}
	fmt.Println()
}

func (paper Paper) sprintf() string {
	formatted := ""
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
	formatted += "\n"

	return formatted
}

func (points paperPoints) loadPaperPoints(colSize, rowSize int) Paper {
	paper := createBlankPaper(colSize, rowSize)

	for _, coordinate := range points {
		paper[coordinate.row][coordinate.col] = "#"
	}

	return paper
}

func calculatedPaper(instructions []string, coordinates []string) Paper {
	coords := paperPoints(getCoordinates(coordinates))
	colSize, rowSize := getPaperSize(coords)
	for _, instruction := range instructions {
		lineString := strings.Split(instruction, "=")[1]
		line, err := strconv.Atoi(lineString)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(instruction, "x") {
			for i, coordinate := range coords {
				if coordinate.col > line {
					coords[i].col = colSize - coordinate.col - 1
				} else if coordinate.col == line {
					coords[i].col = 0
					coords[i].row = 0
				}
			}

			colSize = (colSize) / 2
		} else {
			for i, coordinate := range coords {
				if coordinate.row > line {
					coords[i].row = rowSize - coordinate.row - 1
				} else if coordinate.row == line {
					coords[i].row = 0
					coords[i].col = 0
				}
			}

			rowSize = (rowSize) / 2
		}
	}

	paper := coords.loadPaperPoints(colSize, rowSize)

	return paper
}

// Day13Part1Simulated returns the simulated version for the number of dots visible after the first fold
func Day13Part1Simulated(rawData []string) int {
	instructions, coordinates := splitCoordinatesAndInstructions(rawData)

	paper := loadPaper(coordinates)
	paper = paper.fold(instructions[0])

	return paper.countDots()
}

// Day13Part1Calculated returns the calculated version for the number of dots visible after the first fold
func Day13Part1Calculated(rawData []string) int {
	instructions, coordinates := splitCoordinatesAndInstructions(rawData)

	paper := calculatedPaper([]string{
		instructions[0],
	}, coordinates)

	return paper.countDots()
}

// Day13Part2Simulated displays the simulated version of the 8 letters (in ASCII picture format) which is the code for the thermal camera
func Day13Part2Simulated(rawData []string) string {
	instructions, coordinates := splitCoordinatesAndInstructions(rawData)

	paper := loadPaper(coordinates)
	for _, instruction := range instructions {
		paper = paper.fold(instruction)
	}

	return paper.sprintf()
}

// Day13Part2Calculated displays the calculated version of the 8 letters (in ASCII picture format) which is the code for the thermal camera
func Day13Part2Calculated(rawData []string) string {
	instructions, coordinates := splitCoordinatesAndInstructions(rawData)

	paper := calculatedPaper(instructions, coordinates)

	return paper.sprintf()
}
