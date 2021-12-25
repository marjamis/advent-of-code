package advent2021

// Day1Part1 returns the number of increases in depth compared to the previous reading
func Day1Part1(depthMeasurements ...int) (countOfIncreases int) {
	for i := 1; i < len(depthMeasurements); i++ {
		if depthMeasurements[i] > depthMeasurements[i-1] {
			countOfIncreases++
		}
	}

	return
}

//Day1Part2 returns the number of increases in depth compared to the previous window (the sum of three in order readings)
func Day1Part2(depthMeasurements ...int) (countOfIncreases int) {
	previousWindow := 0
	for i := 1; i < len(depthMeasurements)-1; i++ {
		currentWindow := depthMeasurements[i-1] + depthMeasurements[i] + depthMeasurements[i+1]
		if currentWindow > previousWindow && previousWindow != 0 {
			countOfIncreases++
		}
		previousWindow = currentWindow
	}

	return
}
