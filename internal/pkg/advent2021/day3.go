package advent2021

import (
	"bytes"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Day3Part1 returns the power consumption of the submarine based off of the diagnostic report intput
func Day3Part1(diagnosticReport []string) (powerConsumption int) {
	var gammaRate []byte
	var epsilonRate []byte

	for col := 0; col < len(diagnosticReport[0]); col++ {
		var countOfZeroes int

		for row := 0; row < len(diagnosticReport); row++ {
			// Note: 48 is the ASCII number for a zero
			if 48 == diagnosticReport[row][col] {
				countOfZeroes++
			}
		}

		if countOfZeroes > len(diagnosticReport)-countOfZeroes {
			// Note: 48 is the ASCII number for a zero
			gammaRate = append(gammaRate, 48)
			// Note: 49 is the ASCII number for a one
			epsilonRate = append(epsilonRate, 49)
		} else {
			// Note: 49 is the ASCII number for a one
			gammaRate = append(gammaRate, 49)
			// Note: 48 is the ASCII number for a zero
			epsilonRate = append(epsilonRate, 48)
		}
	}

	// Converts the string to binary
	gammaRateNumber, err := strconv.ParseInt(bytes.NewBuffer(gammaRate).String(), 2, 64)
	if err != nil {
		log.Error(err)
	}

	// Converts the string to binary
	epsilonRateNumber, err := strconv.ParseInt(string(epsilonRate), 2, 64)
	if err != nil {
		log.Error(err)
	}

	return int(gammaRateNumber * epsilonRateNumber)
}

// getRelevantDiagnosticLines will return a subset of diagnostic report
func getRelevantDiagnosticLines(diagnosticReportLines []string, position int, useHighCount bool) (relevantDiagnosticReportLines []string) {
	var zeroList, oneList []string

	for row := 0; row < len(diagnosticReportLines); row++ {
		if 48 == diagnosticReportLines[row][position] {
			zeroList = append(zeroList, diagnosticReportLines[row])
		} else {
			oneList = append(oneList, diagnosticReportLines[row])
		}
	}

	if useHighCount {
		if len(zeroList) > len(oneList) {
			relevantDiagnosticReportLines = zeroList
		} else {
			relevantDiagnosticReportLines = oneList
		}
	} else {
		if len(zeroList) > len(oneList) {
			relevantDiagnosticReportLines = oneList
		} else {
			relevantDiagnosticReportLines = zeroList
		}
	}

	return
}

// Day3Part2 returns the life support rating of the submarine based off of oxygen generator rating and the CO2 scrubber rating
func Day3Part2(diagnosticReport []string) (lifeSupportRating int) {
	findRating := func(useHighCount bool) int64 {
		// Makes a copy of the diagnosticReport to be whittled down via this process
		list := make([]string, len(diagnosticReport))
		copy(list, diagnosticReport)

		// Loops through diagnosticReport to find the relevant lines of the rating type being looked for
		for position := 0; position < len(diagnosticReport[0]); position++ {
			list = getRelevantDiagnosticLines(list, position, useHighCount)
			if len(list) == 1 {
				break
			}
		}

		// Converts the remaining diagnosticReport line, which is the rating for the current type, and converts it to an integer
		number, err := strconv.ParseInt(list[0], 2, 64)
		if err != nil {
			fmt.Println(err)
		}

		return number
	}

	oxygenGeneratorRating := findRating(true)
	co2scrubberRating := findRating(false)

	return int(oxygenGeneratorRating * co2scrubberRating)
}
