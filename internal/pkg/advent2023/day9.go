package advent2023

import (
	"strconv"
	"strings"
)

func extrapolateValue(reading string, right bool) (extrapolatedValue int) {
	stringValues := strings.Fields((reading))
	intValues := make([]int, len(stringValues))
	for index, v := range stringValues {
		var err error
		intValues[index], err = strconv.Atoi(v)
		if err != nil {
			panic("Failed to convert to a string")
		}
	}

	result := extrapolate(intValues, right)
	if right {
		extrapolatedValue = result[len(result)-1]
	} else {
		extrapolatedValue = result[0]
	}

	return
}

func extrapolate(input []int, right bool) []int {
	newValues := make([]int, len(input)-1)

	for index := range newValues {
		newValues[index] = input[index+1] - input[index]
	}

	allEqualZero := true
	for _, value := range newValues {
		if value != 0 {
			allEqualZero = false
			break
		}
	}

	if !allEqualZero {
		newValues = extrapolate(newValues, right)
	}

	if right {
		input = append(input, input[len(input)-1]+newValues[len(newValues)-1])
	} else {
		input = append([]int{input[0] - newValues[0]}, input...)
	}

	return input
}

// Day9Part1 returns sum of right extrapolated values
func Day9Part1(readings []string) (sumOfExtrapolatedValues int) {
	for _, reading := range readings {
		sumOfExtrapolatedValues += extrapolateValue(reading, true)
	}

	return
}

// Day9Part2 returns sum of left extrapolated values
func Day9Part2(readings []string) (sumOfExtrapolatedValues int) {
	for _, reading := range readings {
		sumOfExtrapolatedValues += extrapolateValue(reading, false)
	}

	return
}
