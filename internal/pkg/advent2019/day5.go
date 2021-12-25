package advent2019

import (
	"strconv"
)

// day5OpcodeBreak takes the operation details and breaks into it's relevant parts
func day5OpcodeBreak(data int) (int, int, int, int) {
	opcode := (data % 10) + ((data / 10 % 10) * 10)
	p1 := data / 100 % 10
	p2 := data / 1000 % 10
	p3 := data / 10000 % 10

	return opcode, p1, p2, p3
}

// Day5 returns the modified instruction set and the diagnostic code
func Day5(systemID string, providedInstructions ...int) ([]int, int) {
	// Makes a copy of the data as it's modified through the process
	instructions := make([]int, len(providedInstructions))
	copy(instructions, providedInstructions)

	position := 0
	output := 0

	modeSelector := func(position int, p1 int, p2 int) (x int, y int) {
		if p1 == 0 {
			x = instructions[instructions[position+1]]
		} else {
			x = instructions[position+1]
		}

		if p2 != -1 {
			if p2 == 0 {
				y = instructions[instructions[position+2]]
			} else {
				y = instructions[position+2]
			}
		}

		return
	}

	for true {
		opcode, p1, p2, _ := day5OpcodeBreak(instructions[position])
		switch opcode {
		case 1:
			// Addition
			x, y := modeSelector(position, p1, p2)
			instructions[instructions[position+3]] = x + y
			position += 4
		case 2:
			// Multiplication
			x, y := modeSelector(position, p1, p2)
			instructions[instructions[position+3]] = x * y
			position += 4
		case 3:
			// Input for a parameter
			systemIDI, _ := strconv.Atoi(systemID)
			instructions[instructions[position+1]] = systemIDI
			position += 2
		case 4:
			// Output value of parameter
			x, _ := modeSelector(position, p1, -1)
			output = x
			position += 2
		case 5:
			// Jump-if-true
			x, y := modeSelector(position, p1, p2)
			if x != 0 {
				position = y
			} else {
				position += 3
			}
		case 6:
			// Jump-if-false
			x, y := modeSelector(position, p1, p2)
			if x == 0 {
				position = y
			} else {
				position += 3
			}
		case 7:
			// Less than
			x, y := modeSelector(position, p1, p2)
			if x < y {
				instructions[instructions[position+3]] = 1
			} else {
				instructions[instructions[position+3]] = 0
			}
			position += 4
		case 8:
			// Equals
			x, y := modeSelector(position, p1, p2)
			if x == y {
				instructions[instructions[position+3]] = 1
			} else {
				instructions[instructions[position+3]] = 0
			}
			position += 4
		case 99:
			// End of app
			return instructions, output
		default:
			return nil, 0
		}
	}

	return nil, 0
}
