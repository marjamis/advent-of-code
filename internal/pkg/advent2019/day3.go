package advent2019

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/marjamis/advent-of-code/pkg/helpers"
)

type coordinates struct {
	X int
	Y int
}

const (
	gridLength  = 30000
	startPointX = gridLength / 2
	startPointY = gridLength / 2
)

var (
	clashes, wire1CoordinatesPath, wire2CoordinatesPath []*coordinates
	clashAbsValues                                      []int
)

func day3AddWirePoint(flag int, grid *[gridLength][gridLength]int, x int, y int) {
	if grid[x][y] == 1 && flag == 2 {
		grid[x][y] = 3
		clashes = append(clashes, &coordinates{x, y})
		clashAbsValues = append(clashAbsValues, helpers.ManhattansDistance(x, y, startPointX, startPointY))

	} else {
		grid[x][y] = flag
	}

	if flag == 1 {
		wire1CoordinatesPath = append(wire1CoordinatesPath, &coordinates{x, y})
	}

	if flag == 2 {
		wire2CoordinatesPath = append(wire2CoordinatesPath, &coordinates{x, y})
	}
}

func day3Direction(flag int, instruction string, grid *[gridLength][gridLength]int, x int, y int) (int, int) {
	direction := string(instruction[0])
	steps, err := strconv.Atoi(instruction[1:])
	if err != nil {
		fmt.Println(err)
	}

	switch direction {
	case "U":
		// Up
		for i := 0; i < steps; i++ {
			//Note: This is -1 as it's go up the array which is minusing of rows
			y = y - 1
			day3AddWirePoint(flag, grid, x, y)
		}
	case "D":
		// Down
		for i := 0; i < steps; i++ {
			//Note: This is +1 as it's go down the array which is plusing of rows
			y = y + 1
			day3AddWirePoint(flag, grid, x, y)
		}
	case "L":
		// Left
		for i := 0; i < steps; i++ {
			x = x - 1
			day3AddWirePoint(flag, grid, x, y)
		}
	case "R":
		// Right
		for i := 0; i < steps; i++ {
			x = x + 1
			day3AddWirePoint(flag, grid, x, y)
		}
	}

	return x, y
}

// Day3Manhattan function
func Day3Manhattan() int {
	sort.Ints(clashAbsValues)

	return clashAbsValues[0]
}

// Day3Steps function
func Day3Steps() int {
	getSteps := func(wirePath []*coordinates, clash coordinates) (count int) {
		for _, coord := range wirePath {
			count++
			if coord.X == clash.X && coord.Y == clash.Y {
				break
			}
		}

		return
	}

	var clashStepCount []int
	for _, clash := range clashes {
		totalSteps := getSteps(wire1CoordinatesPath, *clash) + getSteps(wire2CoordinatesPath, *clash)
		clashStepCount = append(clashStepCount, totalSteps)
	}
	sort.Ints(clashStepCount)

	return clashStepCount[0]
}

// Day3 function
func Day3(wire1 []string, wire2 []string, processingFunction func() int) int {
	grid := [gridLength][gridLength]int{}
	grid[startPointX][startPointY] = -1
	// TODO Resetting this TODO maybe a better way
	clashes, wire1CoordinatesPath, wire2CoordinatesPath, clashAbsValues = nil, nil, nil, nil

	// Note: Used a closure as didnt need to be it's own function and looks cool
	placeWires := func(flag int, wire []string) {
		x := startPointX
		y := startPointY

		for wirePoint := range wire {
			x, y = day3Direction(flag, wire[wirePoint], &grid, x, y)
		}
	}

	placeWires(1, wire1)
	placeWires(2, wire2)

	return processingFunction()
}
