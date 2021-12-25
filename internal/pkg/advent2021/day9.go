package advent2021

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/marjamis/advent-of-code/pkg/helpers"
)

type heights [][]int

type point struct {
	row, col int
	height   int
}

func createHeightMap(heightMapInput []string) (heightMap heights) {
	for _, row := range heightMapInput {
		var outputRow = make([]int, len(row))
		for col := range row {
			t, err := strconv.Atoi(string(row[col]))
			if err != nil {
				log.Fatal(err)
			}
			outputRow[col] = t
		}
		heightMap = append(heightMap, outputRow)
	}

	return
}

func (hm heights) print() {
	for _, row := range hm {
		for col := range row {
			fmt.Printf("%+v", row[col])
		}
		fmt.Println()
	}
}

func (hm heights) isLowPoint(row, col int) bool {
	current := hm[row][col]
	// Above
	if row >= 1 && hm[row-1][col] <= current {
		return false
	}
	// Below
	if row < (len(hm)-1) && hm[row+1][col] <= current {
		return false
	}

	// Left
	if col >= 1 && hm[row][col-1] <= current {
		return false
	}

	// Right
	if col < (len(hm[row])-1) && hm[row][col+1] <= current {
		return false
	}

	return true
}

func (hm heights) findLowPoints() (lowPoints []point) {
	for rowIndex, row := range hm {
		for colIndex := range row {
			if hm.isLowPoint(rowIndex, colIndex) {
				lowPoints = append(lowPoints, point{rowIndex, colIndex, hm[rowIndex][colIndex]})
			}
		}
	}

	return
}

func (hm heights) isPartOfBasin(row, col, height int) bool {
	if row > len(hm)-1 || row < 0 || col > len(hm[0])-1 || col < 0 {
		return false
	}

	difference := helpers.Abs(hm[row][col] - height)
	if difference == 0 || difference == 1 {
		return true
	}

	return false
}

func (hm heights) findBasinSize(row, col int) (size int) {
	validPoints := []point{}
	currentHeight := hm[row][col]

	if hm.isPartOfBasin(row-1, col, currentHeight) {
		validPoints = append(validPoints, point{row - 1, col, 0})
	}
	if hm.isPartOfBasin(row+1, col, currentHeight) {
		validPoints = append(validPoints, point{row + 1, col, 0})
	}
	if hm.isPartOfBasin(row, col-1, currentHeight) {
		validPoints = append(validPoints, point{row, col - 1, 0})
	}
	if hm.isPartOfBasin(row, col+1, currentHeight) {
		validPoints = append(validPoints, point{row, col + 1, 0})
	}

	subs := 0
	for _, p := range validPoints {
		subs += hm.findBasinSize(p.row, p.col)
	}

	return len(validPoints) + subs
}

// Day9Part1 returns the risk level based on the low points
func Day9Part1(heightMapInput []string) (riskLevel int) {
	heightMap := createHeightMap(heightMapInput)

	lowPoints := heightMap.findLowPoints()
	for _, height := range lowPoints {
		riskLevel += 1 + heightMap[height.row][height.col]
	}

	return
}

// Day9Part2 returns the risk based on all the basins sizes
func Day9Part2(heightMapInput []string) (riskLevel int) {
	heightMap := createHeightMap(heightMapInput)

	basinSizes := []int{}
	lowPoints := heightMap.findLowPoints()
	for _, height := range lowPoints {
		basinSizes = append(basinSizes, heightMap.findBasinSize(height.row, height.col))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}
