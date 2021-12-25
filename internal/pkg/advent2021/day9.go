package advent2021

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/marjamis/advent-of-code/pkg/helpers"
)

type HeightMap [][]int

type Point struct {
	row, col int
	height   int
}

func createHeightMap(heightMapInput []string) (heightMap HeightMap) {
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

func (hm HeightMap) print() {
	for _, row := range hm {
		for col := range row {
			fmt.Printf("%+v", row[col])
		}
		fmt.Println()
	}
}

func (hm HeightMap) isLowPoint(row, col int) bool {
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

func (hm HeightMap) findLowPoints() (lowPoints []Point) {
	for rowIndex, row := range hm {
		for colIndex := range row {
			if hm.isLowPoint(rowIndex, colIndex) {
				lowPoints = append(lowPoints, Point{rowIndex, colIndex, hm[rowIndex][colIndex]})
			}
		}
	}

	return
}

func (hm HeightMap) isPartOfBasin(row, col, height int) bool {
	if row > len(hm)-1 || row < 0 || col > len(hm[0])-1 || col < 0 {
		return false
	}

	difference := helpers.Abs(hm[row][col] - height)
	if difference == 0 || difference == 1 {
		return true
	}

	return false
}

func (hm HeightMap) findBasinSize(row, col int) (size int) {
	validPoints := []Point{}
	currentHeight := hm[row][col]

	if hm.isPartOfBasin(row-1, col, currentHeight) {
		validPoints = append(validPoints, Point{row - 1, col, 0})
	}
	if hm.isPartOfBasin(row+1, col, currentHeight) {
		validPoints = append(validPoints, Point{row + 1, col, 0})
	}
	if hm.isPartOfBasin(row, col-1, currentHeight) {
		validPoints = append(validPoints, Point{row, col - 1, 0})
	}
	if hm.isPartOfBasin(row, col+1, currentHeight) {
		validPoints = append(validPoints, Point{row, col + 1, 0})
	}

	subs := 0
	for _, p := range validPoints {
		subs += hm.findBasinSize(p.row, p.col)
	}

	return len(validPoints) + subs
}

func Day9Part1(heightMapInput []string) (riskLevel int) {
	heightMap := createHeightMap(heightMapInput)

	lowPoints := heightMap.findLowPoints()
	for _, height := range lowPoints {
		riskLevel += 1 + heightMap[height.row][height.col]
	}

	return
}

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
