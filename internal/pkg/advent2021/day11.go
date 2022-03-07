package advent2021

import "fmt"

// OctopusEnergyMap is the 2d positioning of all dumbo octopuses
type OctopusEnergyMap [][]int

// PointX is a row/col position
type PointX struct {
	row int
	col int
}

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

func (oem OctopusEnergyMap) flash(row int, col int) (numberOfFlashes int) {
	oem[row][col] = 0
	numberOfFlashes++

	for _, cell := range oem.listValidCells(row, col) {
		if oem[cell.row][cell.col] != 0 {
			oem[cell.row][cell.col]++
			if oem[cell.row][cell.col] == 9 {
				numberOfFlashes += oem.flash(cell.row, cell.col)
			}
		}
	}

	return
}
func (oem OctopusEnergyMap) step() (numberOfFlashes int) {
	for row := 0; row < len(oem); row++ {
		for col := 0; col < len(oem[row]); col++ {
			oem[row][col]++
			if oem[row][col] >= 9 {
				numberOfFlashes += oem.flash(row, col)
			}
		}
	}

	return
}

// Day11Part1 returns the number of flashes from the dumbo octopuses
func Day11Part1(initialEnergyLevels [][]int) (numberOfFlashes int) {
	oem := OctopusEnergyMap(initialEnergyLevels)

	oem.print()

	for i := 0; i < 2; i++ {
		numberOfFlashes += oem.step()
		oem.print()
	}

	return
}
