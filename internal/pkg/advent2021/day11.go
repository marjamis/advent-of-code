package advent2021

import "fmt"

// OctopusEnergyMap is the 2d positioning of all dumbo octopuses
type OctopusEnergyMap [][]int

// PointX is a row/col position
type PointX struct {
	row int
	col int
}

var flashed = []PointX{}

func (oem OctopusEnergyMap) print() {
	for _, row := range oem {
		for _, cell := range row {
			fmt.Printf("%d ", cell)
		}
		fmt.Println()
	}

	fmt.Println()
}

func (oem OctopusEnergyMap) listValidCells(row, col int) (validCells []PointX) {
	// TODO make this more generic
	// Above
	if row >= 1 {
		validCells = append(validCells, PointX{row - 1, col})
	}
	// Below
	if row < (len(oem) - 1) {
		validCells = append(validCells, PointX{row + 1, col})
	}

	// Left
	if col >= 1 {
		validCells = append(validCells, PointX{row, col - 1})
	}

	// Right
	if col < (len(oem[row]) - 1) {
		validCells = append(validCells, PointX{row, col + 1})
	}

	// Diagonal Above/Left
	if row >= 1 && col >= 1 {
		validCells = append(validCells, PointX{row - 1, col - 1})
	}

	// Diagonal Above/Right
	if row >= 1 && col < (len(oem[row])-1) {
		validCells = append(validCells, PointX{row - 1, col + 1})
	}

	// Diagonal Below/Right
	if row < (len(oem)-1) && col < (len(oem[row])-1) {
		validCells = append(validCells, PointX{row + 1, col + 1})
	}

	// Diagonal Below/Left
	if row < (len(oem)-1) && col >= 1 {
		validCells = append(validCells, PointX{row + 1, col - 1})
	}

	return
}

func (oem OctopusEnergyMap) flash() (numberOfFlashes int) {
	increasingPoints := []PointX{}

	alreadyFlashed := func(row, col int) bool {
		for _, f1 := range flashed {
			if f1.row == row && f1.col == col {
				return true
			}
		}
		return false
	}

	for row := 0; row < len(oem); row++ {
		for col := 0; col < len(oem[row]); col++ {
			if oem[row][col] > 9 && !alreadyFlashed(row, col) {
				newPoint := &PointX{
					row: row,
					col: col,
				}
				increasingPoints = append(increasingPoints, *newPoint)
				flashed = append(flashed, *newPoint)
				numberOfFlashes++
			}
		}
	}

	for _, increasingPoint := range increasingPoints {
		validCells := oem.listValidCells(increasingPoint.row, increasingPoint.col)
		for _, validCell := range validCells {
			oem[validCell.row][validCell.col] = increase(oem[validCell.row][validCell.col])
			if oem[validCell.row][validCell.col] > 9 {
				numberOfFlashes += oem.flash()
			}
		}
	}

	return
}

func (oem OctopusEnergyMap) mapping(function func(value int) int) {
	for row := 0; row < len(oem); row++ {
		for col := 0; col < len(oem[row]); col++ {
			// TODO should the assigning happen here or in the function call?
			oem[row][col] = function(oem[row][col])
		}
	}
}

func increase(value int) int {
	return value + 1
}

func setToZero(value int) int {
	if value > 9 {
		return 0
	}

	return value
}

func (oem OctopusEnergyMap) step() (numberOfFlashes int) {
	// Each octopus is increased by one
	oem.mapping(increase)
	// Each one greater than 9 flashes, including surrounding continue until no additional changes
	numberOfFlashes = oem.flash()
	// Reseting the flashed list between steps as each step can allow flashing again
	flashed = []PointX{}
	// Any flashed octopus is set to zero
	oem.mapping(setToZero)

	oem.print()

	return
}

func (oem OctopusEnergyMap) isSynchonised() bool {
	for row := 0; row < len(oem); row++ {
		for col := 0; col < len(oem[row]); col++ {
			if oem[row][col] != 0 {
				return false
			}
		}
	}
	return true
}

// Day11Part1 returns the number of flashes from the dumbo octopuses
func Day11Part1(initialEnergyLevels [][]int) (numberOfFlashes int) {
	oem := OctopusEnergyMap(initialEnergyLevels)

	// oem.print()

	for i := 0; i < 100; i++ {
		numberOfFlashes += oem.step()
	}

	return
}

// Day11Part2 returns the first step in which all dumbo octopuses flash at the same time
func Day11Part2(initialEnergyLevels [][]int) (synchronisedFlashStep int) {
	oem := OctopusEnergyMap(initialEnergyLevels)

	oem.print()

	// Starting a 1 as there isn't a step 0 in this calculation
	for synchronisedFlashStep = 1; ; synchronisedFlashStep++ {
		oem.step()
		if oem.isSynchonised() {
			break
		}
	}

	return
}
