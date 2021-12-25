package advent2021

import (
	"fmt"
)

// Lanternfish contains the attributes of each inidividual Lanternfish
type Lanternfish struct {
	timer int
}

// School contains all the Lanternfish that are born
type School []Lanternfish

func (s School) printState() {
	for _, fish := range s {
		fmt.Printf("%d, ", fish.timer)
	}
	fmt.Println()
}

// decreaseTimer will decrease the timer for the Laternfish to give birth and indicate if it has given birth
func (lf *Lanternfish) decreaseTimer() (spawnChildFish bool) {
	lf.timer--
	if lf.timer < 0 {
		lf.timer = 6
		return true
	}

	return false
}

// simulateSchoolGrowith returns the number of fish after the provided days of simulation based on the initial timers of the original fish
func simulateSchoolGrowth(school School, daysToSimulate int) int {
	// Simulates the population growth
	for i := 1; i <= daysToSimulate; i++ {
		for lanternfish := range school {
			spawnChildFish := school[lanternfish].decreaseTimer()
			if spawnChildFish {
				school = append(school, Lanternfish{
					timer: 8,
				})
			}
		}
	}

	return len(school)
}

// calculateSchoolGrowth simulates the number of fish after the provided number of days
func calculateSchoolGrowth(initialTimer, daysToCalculate int) int {
	return 0
}

// Day6Part1 returns the number of fish
func Day6Part1(initialTimers []int, daysToSimulate int) int {
	school := make(School, len(initialTimers))

	// Using initialTimers to create the start of the school
	for i, initialTimer := range initialTimers {
		school[i] = Lanternfish{
			timer: initialTimer,
		}
	}

	return simulateSchoolGrowth(school, daysToSimulate)
}

// Day6Part2 calculates the number of fish after the provided number of days
func Day6Part2(initialTimers []int, daysToCalculate int) (schoolSize int) {
	// Calculating the number of fish spawned based off of each possible initial timer based on the number of days
	initials := []int{
		0,
		calculateSchoolGrowth(1, daysToCalculate),
		calculateSchoolGrowth(2, daysToCalculate),
		calculateSchoolGrowth(3, daysToCalculate),
		calculateSchoolGrowth(4, daysToCalculate),
		calculateSchoolGrowth(5, daysToCalculate),
	}

	// With the initial timers calculated we can loop through each actual fish with their initial timers to get the total
	for _, laternfish := range initialTimers {
		schoolSize += initials[laternfish]
	}

	return schoolSize
}
