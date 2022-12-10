package advent2022

import (
	"fmt"
	"strconv"
	"strings"
)

func generateCyclesList(instructions []string) (cycles []int) {
	cycles = []int{1}

	for _, instruction := range instructions {
		inst := strings.Split(instruction, " ")
		switch inst[0] {
		case "noop":
			// No operation
			cycles = append(cycles, cycles[len(cycles)-1])
		case "addx":
			// No operation
			previousValue := cycles[len(cycles)-1]
			cycles = append(cycles, previousValue)

			// Gets and sets the new X register for this cycle
			val, err := strconv.Atoi(inst[1])
			if err != nil {
				return nil
			}
			registerValue := previousValue + val
			cycles = append(cycles, registerValue)
		}
	}
	return
}

// DisplayFrame displays the rendered frame to stdout
func DisplayFrame(crtDisplay [][]rune) {
	for _, line := range crtDisplay {
		fmt.Println(string(line))
	}
}

// Day10Part1 returns the signal strength as per the calculation
func Day10Part1(instructions []string) (signalStrength int) {
	cycles := generateCyclesList(instructions)
	cyclesToDetermineStrength := []int{
		20, 60, 100, 140, 180, 220,
	}

	for _, cycle := range cyclesToDetermineStrength {
		// There is a -1 to ensure it takes the value before the cycles completion rather than after it's completion. Side effect of calculating the cycles first
		signalStrength += (cycles[cycle-1] * cycle)
	}

	return
}

// Day10Part2 returns the rune map (i.e. pixel display) of the CRT display
func Day10Part2(instructions []string) (crtDisplay [][]rune) {
	displayResolution := struct {
		rows int
		cols int
	}{
		rows: 6,
		cols: 40,
	}

	// With the cycles it's calculated with one long line, for ease, that's split for the 40 * 6 CRT display end results
	displayLine := make([]rune, displayResolution.cols*displayResolution.rows)
	// Initialises the data with .'s as per the puzzle
	for index := range displayLine {
		displayLine[index] = '.'
	}

	// Pre-generates the instructions for ease. Bit less performant but shows what's happening a bit nicer
	cycles := generateCyclesList(instructions)
	// It starts at cycle 1 but as the generateCyclesList stores the value  after the cycle, not during, we have to go one step backwards. Side effect of calculating the cycles first
	for cycle := 0; cycle < len(cycles); cycle++ {
		spriteCenterPosition := cycles[cycle]
		pixelWritePosition := cycle % displayResolution.cols
		// Calculates if the pixel write location is in the 3 sprite length, based off of the sprites center position. If yes, write the pixel to the display
		if spriteCenterPosition-1 == pixelWritePosition || spriteCenterPosition == pixelWritePosition || spriteCenterPosition+1 == pixelWritePosition {
			displayLine[cycle] = '#'
		}
	}

	// With the displayLine pixels written split it to 40 pixels per line for the 40 * 6 CRT display resolution
	crtDisplay = make([][]rune, displayResolution.rows)
	for rows := 0; rows < displayResolution.rows; rows++ {
		crtDisplay[rows] = displayLine[rows*displayResolution.cols : rows*displayResolution.cols+displayResolution.cols]
	}

	return
}
