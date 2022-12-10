package advent2022

import (
	"fmt"
	"strconv"
	"strings"
)

// RopePosition contains the row and col of where the rope is
type RopePosition struct {
	row int
	col int
}

func generateMapKey(position *RopePosition) string {
	return fmt.Sprintf("%d|%d", position.row, position.col)
}

func simulateRopeMovement(ropeInstructions []string, numberOfKnots int) int {
	// Create initial rope positions for the number of knots
	knots := make([]*RopePosition, numberOfKnots)
	for i := 0; i < numberOfKnots; i++ {
		knots[i] = &RopePosition{
			row: 1000,
			col: 1000,
		}
	}

	// Tracks all unique positions the tail, i.e. the last knot, has been in
	trackingPositions := map[string]int{
		generateMapKey(knots[len(knots)-1]): 1,
	}

	for _, instruction := range ropeInstructions {
		split := strings.Split(instruction, " ")
		direction := split[0]
		count, err := strconv.Atoi(split[1])
		if err != nil {
			return -1
		}

		for i := 0; i < count; i++ {
			// Moves the head knot based on instructions
			switch direction {
			case "L":
				knots[0].col--
			case "R":
				knots[0].col++
			case "U":
				knots[0].row--
			case "D":
				knots[0].row++
			}

			// For each knot move it based on the knot in front
			for knot := 1; knot < len(knots); knot++ {
				knotToPreviousKnotRowDifference := knots[knot-1].row - knots[knot].row
				knotToPreviousKnotColDifference := knots[knot-1].col - knots[knot].col

				/*
					Example positions map to explain where each knot moves based on where the previous knot is:
						1	2	3	4	5
						6	7	8	9	10
						11	12	X	14	15
						16	17	18	19	20
						21	22	23	24	25

					X is the current knot.

					NOTE: Likely can be condensed but I like the verbosity of positions
				*/
				if knotToPreviousKnotRowDifference == -2 && knotToPreviousKnotColDifference == -2 {
					// Previous knot position: 1 and current knot moves to: 7
					knots[knot].col--
					knots[knot].row--
				} else if knotToPreviousKnotRowDifference == -2 && knotToPreviousKnotColDifference == 2 {
					// Previous knot position: 5 and current knot moves to: 9
					knots[knot].col++
					knots[knot].row--
				} else if knotToPreviousKnotRowDifference == 2 && knotToPreviousKnotColDifference == 2 {
					// Previous knot position: 25 and current knot moves to: 19
					knots[knot].col++
					knots[knot].row++
				} else if knotToPreviousKnotRowDifference == 2 && knotToPreviousKnotColDifference == -2 {
					// Previous knot position: 21 and current knot moves to: 17
					knots[knot].col--
					knots[knot].row++
				} else if knotToPreviousKnotRowDifference == -1 && knotToPreviousKnotColDifference == 2 {
					// Previous knot position: 10 and current knot moves to: 9
					knots[knot].col++
					knots[knot].row--
				} else if knotToPreviousKnotRowDifference == 1 && knotToPreviousKnotColDifference == -2 {
					// Previous knot position: 16 and current knot moves to: 17
					knots[knot].col--
					knots[knot].row++
				} else if knotToPreviousKnotRowDifference == 1 && knotToPreviousKnotColDifference == 2 {
					// Previous knot position: 20 and current knot moves to: 19
					knots[knot].col++
					knots[knot].row++
				} else if knotToPreviousKnotRowDifference == -1 && knotToPreviousKnotColDifference == -2 {
					// Previous knot position: 6 and current knot moves to: 7
					knots[knot].col--
					knots[knot].row--
				} else if knotToPreviousKnotRowDifference == -2 && knotToPreviousKnotColDifference == 1 {
					// Previous knot position: 4 and current knot moves to: 9
					knots[knot].col++
					knots[knot].row--
				} else if knotToPreviousKnotRowDifference == 2 && knotToPreviousKnotColDifference == -1 {
					// Previous knot position: 22 and current knot moves to: 17
					knots[knot].col--
					knots[knot].row++
				} else if knotToPreviousKnotRowDifference == 2 && knotToPreviousKnotColDifference == 1 {
					// Previous knot position: 24 and current knot moves to: 19
					knots[knot].col++
					knots[knot].row++
				} else if knotToPreviousKnotRowDifference == -2 && knotToPreviousKnotColDifference == -1 {
					// Previous knot position: 2 and current knot moves to: 7
					knots[knot].col--
					knots[knot].row--
				} else if knotToPreviousKnotRowDifference == 2 {
					// Previous knot position: 23 and current knot moves to: 18
					knots[knot].row++
				} else if knotToPreviousKnotRowDifference == -2 {
					// Previous knot position: 3 and current knot moves to: 8
					knots[knot].row--
				} else if knotToPreviousKnotColDifference == 2 {
					// Previous knot position: 15 and current knot moves to: 14
					knots[knot].col++
				} else if knotToPreviousKnotColDifference == -2 {
					// Previous knot position: 11 and current knot moves to: 12
					knots[knot].col--
				}
			}
			trackingPositions[generateMapKey(knots[len(knots)-1])]++
		}
	}

	return len(trackingPositions)
}

// Day9Part1 returns the unique positions for a 2 knot'ed rope
func Day9Part1(ropeInstructions []string) int {
	return simulateRopeMovement(ropeInstructions, 2)
}

// Day9Part2 returns the unique positions for a 10 knot'ed rope
func Day9Part2(ropeInstructions []string) int {
	return simulateRopeMovement(ropeInstructions, 10)
}
