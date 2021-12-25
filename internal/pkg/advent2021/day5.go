package advent2021

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// VentMap is the 2d int array that represents the location of the heat vents
type VentMap [1000][1000]int

// Note: Generally the above should be dynamic and have more error checking on exceeding the limits

// coordinates represents the X,Y coordinates on the VentMap
type coordinates struct {
	startRow, startCol, endRow, endCol int
}

// getVentLineCoordinates takes in a line of input and generates the relevant coordinates
func getVentLineCoordinates(ventLine string) coordinates {
	v := strings.Split(ventLine, " -> ")
	fromCoordinates := strings.Split(v[0], ",")
	toCoordinates := strings.Split(v[1], ",")

	startCol, err := strconv.Atoi(fromCoordinates[0])
	if err != nil {
		log.Fatal(err)
	}

	startRow, err := strconv.Atoi(fromCoordinates[1])
	if err != nil {
		log.Fatal(err)
	}

	endRow, err := strconv.Atoi(toCoordinates[1])
	if err != nil {
		log.Fatal(err)
	}

	endCol, err := strconv.Atoi(toCoordinates[0])
	if err != nil {
		log.Fatal(err)
	}

	return coordinates{
		startRow: startRow,
		startCol: startCol,
		endRow:   endRow,
		endCol:   endCol,
	}
}

func (c *coordinates) printCoordinates() {
	fmt.Printf("%d %d %d %d\n", c.startRow, c.endRow, c.startCol, c.endCol)
}

// drawVerticalLine takes coordinates of a vertical vent line and draws it on the map
func (vm *VentMap) drawVerticalLine(coords coordinates) {
	// Note: this does overwrite settings for coordinates but for now this is OK (need to think and investigate)
	if coords.startRow > coords.endRow {
		tmpX := coords.startRow
		coords.startRow = coords.endRow
		coords.endRow = tmpX
	}

	for row := coords.startRow; row <= coords.endRow; row++ {
		vm[row][coords.endCol]++
	}
}

// drawHorizontalLine takes coordinates of a horizontal vent line and draws it on the map
func (vm *VentMap) drawHorizontalLine(coords coordinates) {
	// Note: this does overwrite settings for coordinates but for now this is OK (need to think and investigate)
	if coords.startCol > coords.endCol {
		tmpY := coords.startCol
		coords.startCol = coords.endCol
		coords.endCol = tmpY
	}

	for col := coords.startCol; col <= coords.endCol; col++ {
		vm[coords.startRow][col]++
	}
}

// drawDiagonalLine takes coordinates of a diagonal vent line and draws it on the map
func (vm *VentMap) drawDiagonalLine(coords coordinates) {
	if (coords.startRow > coords.endRow) && (coords.startCol < coords.endCol) {
		// Used if the row is decreasing and the column is increasing in direction
		for row, col := coords.startRow, coords.startCol; row >= coords.endRow && col <= coords.endCol; row, col = row-1, col+1 {
			vm[row][col]++
		}
	} else if (coords.startCol > coords.endCol) && (coords.startRow < coords.endRow) {
		// Used if the row is increasing and the column is decreasing in direction
		for row, col := coords.startRow, coords.startCol; row <= coords.endRow && col >= coords.endCol; row, col = row+1, col-1 {
			vm[row][col]++
		}
	} else {
		//Used if the row and column are both increasing or both decreasing in the same direction

		// If both coordinates are decreasing the coordinates are flipped to make them increase for easier calculations
		if coords.startRow > coords.endRow && coords.startCol > coords.endCol {
			tmpRow := coords.startRow
			tmpCol := coords.startCol

			coords.startRow = coords.endRow
			coords.startCol = coords.endCol

			coords.endRow = tmpRow
			coords.endCol = tmpCol
		}

		for row, col := coords.startRow, coords.startCol; row <= coords.endRow && col <= coords.endCol; row, col = row+1, col+1 {
			vm[row][col]++
		}
	}
}

// drawLine takes coordinates and draws the required vent line on the vent map
func (vm *VentMap) drawLine(c coordinates, useDiagonal bool) {
	if (c.startRow-c.endRow) != 0 && (c.startCol-c.endCol) != 0 {
		// Note: This is used to determine if the drawing considering diagonal line drawings or skips them, as per the puzzle
		if !useDiagonal {
			return
		}
		vm.drawDiagonalLine(c)
	} else if c.startRow != c.endRow {
		vm.drawVerticalLine(c)
	} else if c.startCol != c.endCol {
		vm.drawHorizontalLine(c)
	}

	return
}

func (vm *VentMap) printMap() {
	for row := 0; row < len(vm); row++ {
		for col := 0; col < len(vm[0]); col++ {
			fmt.Printf(" %d", vm[row][col])
		}
		fmt.Println()
	}
}

// countMoreThanTwoVents will loop through the vent map and reutrn the number of locations with 2 or more vents
func (vm *VentMap) countMoreThanTwoVents() (count int) {
	for row := 0; row < len(vm); row++ {
		for col := 0; col < len(vm[0]); col++ {
			if vm[row][col] >= 2 {
				count++
			}
		}
	}

	return
}

// Day5Part1 takes the input of vent lines (horizontal and vertical) and returns the count of locations with 2 or more vents
func Day5Part1(ventLines []string) int {
	var ventMap VentMap

	for _, ventLine := range ventLines {
		coords := getVentLineCoordinates(ventLine)
		ventMap.drawLine(coords, false)
	}

	return ventMap.countMoreThanTwoVents()
}

// Day5Part2 takes the input of vent lines (horizontal, vertical, and diagonal) and returns the count of locations with 2 or more vents
func Day5Part2(ventLines []string) int {
	var ventMap VentMap

	for _, ventLine := range ventLines {
		coords := getVentLineCoordinates(ventLine)
		ventMap.drawLine(coords, true)
	}

	return ventMap.countMoreThanTwoVents()
}
