package advent2021

import "fmt"

// OctopusEnergyMap is the 2d positioning of all dumbo octopuses
type OctopusEnergyMap [][]int

// Coordinates is a row/col position on a 2D matrix
type Coordinates struct {
	row int
	col int
}

var flashedThisStep = []Coordinates{}

func (oem OctopusEnergyMap) print() {
	for _, row := range oem {
		for _, cell := range row {
			fmt.Printf("%d ", cell)
		}
		fmt.Println()
	}

	fmt.Println()
}

func (oem OctopusEnergyMap) getSurroundingOctopuses(row, col int) (surroundingOctopuses []Coordinates) {
	// Above
	if row >= 1 {
		surroundingOctopuses = append(surroundingOctopuses, Coordinates{row - 1, col})
	}
	// Below
	if row < (len(oem) - 1) {
		surroundingOctopuses = append(surroundingOctopuses, Coordinates{row + 1, col})
	}

	// Left
	if col >= 1 {
		surroundingOctopuses = append(surroundingOctopuses, Coordinates{row, col - 1})
	}

	// Right
	if col < (len(oem[row]) - 1) {
		surroundingOctopuses = append(surroundingOctopuses, Coordinates{row, col + 1})
	}

	// Diagonal Above/Left
	if row >= 1 && col >= 1 {
		surroundingOctopuses = append(surroundingOctopuses, Coordinates{row - 1, col - 1})
	}

	// Diagonal Above/Right
	if row >= 1 && col < (len(oem[row])-1) {
		surroundingOctopuses = append(surroundingOctopuses, Coordinates{row - 1, col + 1})
	}

	// Diagonal Below/Right
	if row < (len(oem)-1) && col < (len(oem[row])-1) {
		surroundingOctopuses = append(surroundingOctopuses, Coordinates{row + 1, col + 1})
	}

	// Diagonal Below/Left
	if row < (len(oem)-1) && col >= 1 {
		surroundingOctopuses = append(surroundingOctopuses, Coordinates{row + 1, col - 1})
	}

	return
}

func (oem OctopusEnergyMap) flash() (numberOfFlashes int) {
	flashingOctopuses := []Coordinates{}

	hasFlashed := func(row, col int) bool {
		for _, f1 := range flashedThisStep {
			if f1.row == row && f1.col == col {
				return true
			}
		}
		return false
	}

	for row := 0; row < len(oem); row++ {
		for col := 0; col < len(oem[row]); col++ {
			if oem[row][col] > 9 && !hasFlashed(row, col) {
				newPoint := &Coordinates{
					row: row,
					col: col,
				}
				flashingOctopuses = append(flashingOctopuses, *newPoint)
				flashedThisStep = append(flashedThisStep, *newPoint)
				numberOfFlashes++
			}
		}
	}

	for _, flashingOctopus := range flashingOctopuses {
		surroundingOctopuses := oem.getSurroundingOctopuses(flashingOctopus.row, flashingOctopus.col)
		for _, octopus := range surroundingOctopuses {
			oem[octopus.row][octopus.col] = increase(oem[octopus.row][octopus.col])
			if oem[octopus.row][octopus.col] > 9 {
				numberOfFlashes += oem.flash()
			}
		}
	}

	return
}

func (oem OctopusEnergyMap) mapping(function func(value int) int) {
	for row := 0; row < len(oem); row++ {
		for col := 0; col < len(oem[row]); col++ {
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
	flashedThisStep = []Coordinates{}
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
