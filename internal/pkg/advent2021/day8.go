package advent2021

import (
	"strings"

	"github.com/marjamis/advent-of-code/pkg/helpers"
)

// Segments provides a mapping of the original positions to the new positions
type Segments map[rune]rune

// parseDisplayInput will take a given line of input and provided a list of strings for the inputs and outputs for processing
func parseDisplayOutput(linesOfOutputs string) (inputs []string, outputs []string) {
	i := strings.Split(linesOfOutputs, " | ")
	inputs = strings.Split(i[0], " ")
	outputs = strings.Split(i[1], " ")

	return
}

// isValueInMap returns if the provided search values is a value of any key in the Segments map
func isValueInMap(segments Segments, searchForValue rune) bool {
	for _, segmentValue := range segments {
		if segmentValue == searchForValue {
			return true
		}
	}

	return false
}

// isLetterInString returns if the search letter is in the provided string
func isLetterInString(searchString string, searchForLetter rune) bool {
	for _, letter := range searchString {
		if searchForLetter == letter {
			return true
		}
	}

	return false
}

func getSegmentA(inputs []string) (foundSegment rune) {
	var one string
	var seven string
	for _, input := range inputs {
		if len(input) == 2 {
			one = input
		} else if len(input) == 3 {
			seven = input
		}
	}

	for _, letter := range seven {
		if !isLetterInString(one, rune(letter)) {
			foundSegment = rune(letter)
			break
		}
	}

	return
}

func getSegmentB(inputs []string, segments Segments) (foundSegment rune) {
outerLoop:
	for input := range inputs {
		if len(inputs[input]) == 4 {
			for letter := range inputs[input] {
				testingRune := rune(inputs[input][letter])
				if !isValueInMap(segments, testingRune) {
					foundSegment = testingRune
					break outerLoop
				}
			}
		}
	}
	return
}

func getSegmentCandF(inputs []string) (c, f rune) {
	var sixes []string
	var two string

	for _, input := range inputs {
		if len(input) == 6 {
			sixes = append(sixes, input)
		}
	}

	for _, input := range inputs {
		if len(input) == 2 {
			two = input
			break
		}
	}

	for _, input := range inputs {
		if len(input) == 6 {
			for letter := range two {
				if !isLetterInString(input, rune(two[letter])) {
					if letter == 0 {
						c = rune(two[0])
						f = rune(two[1])
					} else {
						c = rune(two[1])
						f = rune(two[0])
					}
				}
			}
		}
	}

	return
}

func getSegmentD(inputs []string, segments Segments) (foundSegment rune) {
	foursMissingValues := []rune{}
	for input := range inputs {
		if len(inputs[input]) == 4 {
			for letter := range inputs[input] {
				testingRune := rune(inputs[input][letter])
				if !isValueInMap(segments, testingRune) {
					foursMissingValues = append(foursMissingValues, testingRune)
				}
			}
		}
	}

	perms := []rune{}
	for input := range inputs {
		if len(inputs[input]) == 5 {
			threesMissingValues := []rune{}
			for letter := range inputs[input] {
				testingRune := rune(inputs[input][letter])
				if !isValueInMap(segments, testingRune) {
					threesMissingValues = append(threesMissingValues, testingRune)
				}
			}
			if len(threesMissingValues) == 2 {
				perms = threesMissingValues
			}
		}
	}

outerLoop:
	for _, foursV := range foursMissingValues {
		for _, threesV := range perms {
			if threesV == foursV {
				foundSegment = foursV
				break outerLoop
			}
		}
	}

	return
}

func getSegmentE(inputs []string, segments Segments) (foundSegment rune) {
outerLoop:
	for input := range inputs {
		if len(inputs[input]) == 7 {
			for letter := range inputs[input] {
				testingRune := rune(inputs[input][letter])
				if !isValueInMap(segments, testingRune) {
					foundSegment = testingRune
					break outerLoop
				}
			}
		}
	}

	return
}

func getSegmentG(inputs []string, segments Segments) (foundSegment rune) {
outerLoop:
	for input := range inputs {
		if len(inputs[input]) == 5 {
			missingValues := []rune{}
			for letter := range inputs[input] {
				testingRune := rune(inputs[input][letter])
				if !isValueInMap(segments, testingRune) {
					missingValues = append(missingValues, testingRune)
				}
			}
			if len(missingValues) == 1 {
				foundSegment = missingValues[0]
				break outerLoop
			}
		}
	}
	return
}

// findDisplayMappings returns a Segements map after following pre-determined rules for a mapping from the correct to corrupted segments
func findDisplayMappings(inputs []string) (segments Segments) {
	// The 7 here is for the number of segments being mapped, no more is needed in this implementation
	segments = make(Segments, 7)

	// The order is important as current rules are based on missing positions through finding of the previous segments, i.e. like Sudoku
	segments['a'] = getSegmentA(inputs)
	segments['c'], segments['f'] = getSegmentCandF(inputs)
	segments['d'] = getSegmentD(inputs, segments)
	segments['b'] = getSegmentB(inputs, segments)
	segments['g'] = getSegmentG(inputs, segments)
	segments['e'] = getSegmentE(inputs, segments)

	return segments
}

// getNumber returns the proper display number based on the corrupted input string and the generated segments mapping
func getNumber(segments Segments, corruptedInput string) (accurateDisplayNumber int) {
	lengthOfInputString := len(corruptedInput)

	switch lengthOfInputString {
	case 2:
		accurateDisplayNumber = 1
	case 3:
		accurateDisplayNumber = 7
	case 4:
		accurateDisplayNumber = 4
	case 7:
		accurateDisplayNumber = 8
	case 5:
		if !isLetterInString(corruptedInput, segments['b']) && !isLetterInString(corruptedInput, segments['f']) {
			accurateDisplayNumber = 2
		}
		if !isLetterInString(corruptedInput, segments['b']) && !isLetterInString(corruptedInput, segments['e']) {
			accurateDisplayNumber = 3
		}
		if !isLetterInString(corruptedInput, segments['c']) && !isLetterInString(corruptedInput, segments['e']) {
			accurateDisplayNumber = 5
		}
	case 6:
		if !isLetterInString(corruptedInput, segments['d']) {
			accurateDisplayNumber = 0
		}
		if !isLetterInString(corruptedInput, segments['c']) {
			accurateDisplayNumber = 6
		}
		if !isLetterInString(corruptedInput, segments['e']) {
			accurateDisplayNumber = 9
		}
	default:
		accurateDisplayNumber = -1
	}

	return
}

// Day8Part1 returns the number of ouputs which would display the numbers 1, 4, 7, or 8
func Day8Part1(linesOfOutputs []string) (count int) {
	var outputs []string
	for _, line := range linesOfOutputs {
		_, outputs = parseDisplayOutput(line)

		for _, output := range outputs {
			// Note: len 2 == num 1, len 3 == num 7, len 4 == num 4, len 7 == num 8
			if len(output) == 2 || len(output) == 3 || len(output) == 4 || len(output) == 7 {
				count++
			}
		}
	}

	return
}

// Day8Part2 returns the calculated total of all of the outputs (i.e. each output number is a position) after unscrambling the provided inputs of the display
func Day8Part2(linesOfOutputs []string) (total int) {
	for _, line := range linesOfOutputs {
		inputs, outputs := parseDisplayOutput(line)
		segments := findDisplayMappings(inputs)

		var subtotal int
		// This reads the output to be calculate backwards to build up the decimal positions of the numbers easily
		for i, j := len(outputs)-1, 0; i >= 0; i, j = i-1, j+1 {
			subtotal += helpers.DecimalPositionOf(j) * getNumber(segments, outputs[i])
		}
		total += subtotal
	}

	return
}
