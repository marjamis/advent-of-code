package advent2021

import (
	"sort"

	"github.com/marjamis/advent-of-code/pkg/helpers"
)

func leastFuelCheck(currentLeastFuel, fuelToTest int) int {
	if fuelToTest < currentLeastFuel {
		return fuelToTest
	}

	return currentLeastFuel
}

func initialFuelCalculation(positionToTest int, initialPositions []int) (fuelUsed int) {
	for _, initialPosition := range initialPositions {
		fuelUsed += helpers.Abs(initialPosition - positionToTest)
	}

	return
}

func updatedFuelCalculation(positionToTest int, initialPositions []int) (fuelUsed int) {
	for _, initialPosition := range initialPositions {
		for i := 1; i <= helpers.Abs(initialPosition-positionToTest); i++ {
			fuelUsed += i
		}
	}

	return
}

// Day7Part1 returns the position that will take the least amount of fuel
func Day7Part1(initialPositions ...int) (leastFuel int) {
	sort.Sort(sort.IntSlice(initialPositions))

	leastFuel = initialFuelCalculation(0, initialPositions)
	for i := initialPositions[1]; i < initialPositions[len(initialPositions)-1]; i++ {
		leastFuel = leastFuelCheck(leastFuel, initialFuelCalculation(i, initialPositions))
	}

	return
}

// Day7Part2 returns the position that will take the least amount of fuel based on the revised fuel calculations
func Day7Part2(initialPositions ...int) (leastFuel int) {
	sort.Sort(sort.IntSlice(initialPositions))

	leastFuel = updatedFuelCalculation(0, initialPositions)
	for i := initialPositions[1]; i < initialPositions[len(initialPositions)-1]; i++ {
		leastFuel = leastFuelCheck(leastFuel, updatedFuelCalculation(i, initialPositions))
	}

	return
}
