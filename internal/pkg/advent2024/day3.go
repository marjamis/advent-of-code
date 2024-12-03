package advent2024

import (
	"regexp"
	"strconv"
	"strings"
)

func findAllMultis(memorySection []byte) (sectionValue int) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for _, mul := range re.FindAll(memorySection, -1) {
		// The size in used string it to remove the wrapping mul( and )
		instruction := strings.Split(string(mul[4:len(mul)-1]), ",")

		first, err := strconv.Atoi(instruction[0])
		if err != nil {
			panic("couldn't convert string to int")
		}

		second, err := strconv.Atoi(instruction[1])
		if err != nil {
			panic("couldn't convert string to int")
		}

		sectionValue += first * second
	}

	return
}

// Day3Part1 return the multiplication of the non-corrupted instructions
func Day3Part1(corruptedMemory string) (multiplicationResults int) {
	return findAllMultis([]byte(corruptedMemory))
}

// Day3Part2 return the multiplication of the non-corrupted instructions
func Day3Part2(corruptedMemory string) (multiplicationResults int) {
	// The (?s) ensure the regex treats newline characters as valid to the .
	// as the input also has multiple lines
	re := regexp.MustCompile(`(?s)don\'t\(\).*?(do\(\)|$)`)
	validInstructions := re.ReplaceAll([]byte(corruptedMemory), []byte(""))

	return findAllMultis([]byte(validInstructions))
}
