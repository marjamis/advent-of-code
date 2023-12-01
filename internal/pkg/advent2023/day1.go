package advent2023

import (
	"strings"
)

func isNumber(row []rune, index int) bool {
	return row[index] >= 48 && row[index] <= 57
}

// Day1Part1 return the total of the calibration values for the trebuchet
func Day1Part1(calculations [][]rune) (totalCalibrationValue int) {
	for _, row := range calculations {
		var numbers []int

		for index := range row {
			if isNumber(row, index) {
				numbers = append(numbers, int(row[index]-48))
			}
		}

		totalCalibrationValue += (numbers[0] * 10) + numbers[len(numbers)-1]
	}

	return
}

var numbers []string = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func isDigitString(row []rune, num string, startNum int, endNum int) bool {
	if startNum < 0 || endNum > len(row) {
		return false
	}

	if strings.Compare(num, string(row[startNum:endNum])) == 0 {
		return true

	}

	return false
}

// Day1Part2 return the total of the calibration values for the trebuchet with text digits allowed
func Day1Part2(calculations [][]rune) (totalCalibrationValue int) {
	for _, row := range calculations {
		var first, last int

	out:
		for i := 0; i < len(row); i++ {
			if isNumber(row, i) {
				first = int(row[i] - 48)
				break out
			}

			for numberValue, number := range numbers {
				startNum := i
				endNum := i + len(number)

				if isDigitString(row, number, startNum, endNum) {
					first = numberValue
					break out
				}
			}
		}

	out2:
		for i := len(row) - 1; i >= 0; i-- {
			if isNumber(row, i) {
				last = int(row[i] - 48)
				break out2
			}

			for numberValue, number := range numbers {
				endNum := i + 1
				startNum := endNum - len(number)

				if isDigitString(row, number, startNum, endNum) {
					last = numberValue
					break out2
				}
			}
		}

		totalCalibrationValue += (first * 10) + last
	}

	return
}
