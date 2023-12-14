package advent2023

import (
	"fmt"
	"strconv"
	"strings"
)

func numberOfWinningOptions(raceTime int, recordDistance int) (possibleVictoriesCount int) {
	// Skipping 0 seconds as that means not holding the button and going nowhere
	// Skipping "time" seconds as that means holding the button forever and going nowhere
	for i := 1; i < raceTime; i++ {
		holdTime := i
		raceTime := (raceTime - holdTime)

		raceDistance := raceTime * holdTime

		if raceDistance > recordDistance {
			possibleVictoriesCount++
		}
	}

	return
}

// Day6Part1 returns the number of victories per race
func Day6Part1(data []string) (multipleOfVictories int) {
	multipleOfVictories = 1
	times := strings.Fields(data[0])
	distances := strings.Fields(data[1])

	for i := 1; i < len(times); i++ {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			panic("Failed to convert to a string")
		}

		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			panic("Failed to convert to a string")
		}

		fmt.Printf("Time: %d Distance: %d\n", time, distance)
		multipleOfVictories *= numberOfWinningOptions(time, distance)
	}

	return
}

// Day6Part2 returns the number of victories for the big race
func Day6Part2(data []string) (possibleVictoriesCount int) {
	possibleVictoriesCount = 1
	// Too much nesting but for now it's OK
	time, err := strconv.Atoi(strings.Replace(strings.Split(data[0], ":")[1], " ", "", -1))
	if err != nil {
		panic("Failed to convert to a string")
	}

	// Too much nesting but for now it's OK
	distance, err := strconv.Atoi(strings.Replace(strings.Split(data[1], ":")[1], " ", "", -1))
	if err != nil {
		panic("Failed to convert to a string")
	}

	fmt.Printf("Time: %d Distance: %d\n", time, distance)
	return numberOfWinningOptions(time, distance)
}
