package advent2021

import (
	log "github.com/sirupsen/logrus"
)

// Lanternfish contains the attributes of each inidividual Lanternfish
type Lanternfish struct {
	timer int
}

// School contains all the Lanternfish that are born
type School []Lanternfish

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
	timeLeftDistribution := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, age := range initialTimers {
		timeLeftDistribution[age]++
	}

	log.Debugf("Initial state: %v", timeLeftDistribution)

	for day := 1; day <= daysToCalculate; day++ {
		numberOfFishBorn := timeLeftDistribution[0]
		// Added to position 7 as it's in the same day and with the next day theyll be in position 6 as expected
		timeLeftDistribution[7] += numberOfFishBorn
		timeLeftDistribution = append(timeLeftDistribution[1:], numberOfFishBorn)

		log.Debugf("After %d day: %v", day, timeLeftDistribution)
	}

	for _, n := range timeLeftDistribution {
		schoolSize += n
	}

	return schoolSize
}
