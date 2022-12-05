package advent2022

import (
	"strconv"
	"strings"

	"github.com/marjamis/advent-of-code/pkg/structures"
)

func moveCrates(input []string, moveCratesFunction func([]structures.Stack, int, int, int)) string {
	breakBetweenStackMappingAndInstructions := 0

	for index, value := range input {
		if strings.Compare(value, "") == 0 {
			breakBetweenStackMappingAndInstructions = index
			break
		}
	}

	originalStackMapping := input[0:breakBetweenStackMappingAndInstructions]
	stackCountListing := strings.Split(originalStackMapping[len(originalStackMapping)-1], " ")
	numberOfStacks, err := strconv.Atoi(stackCountListing[len(stackCountListing)-2])
	if err != nil {
		return ""
	}

	crateToStackMapping := make([]structures.Stack, numberOfStacks, numberOfStacks)

	for i := len(originalStackMapping) - 2; i >= 0; i-- {
		// Starts at 1 and increments by 4 as thats the spacing from the input of the crates positions
		for j := 1; j <= len(originalStackMapping[i])-1; j += 4 {
			// Does a cautionary check to make sure it's a crate
			if originalStackMapping[i][j] >= 65 && originalStackMapping[i][j] <= 90 {
				// The stack that a crate is on can be derived with the below
				owningStack := (j - 1) / 4

				// Maps the crate to the appropriate Stack from the original input
				crateToStackMapping[owningStack].Push(originalStackMapping[i][j])
			}
		}
	}

	// Gets all the instructions from the stack input
	instructions := input[breakBetweenStackMappingAndInstructions+1:]
	for _, instruction := range instructions {
		instructionLine := strings.Split(instruction, " ")

		countOfCratesToMove, err := strconv.Atoi(instructionLine[1])
		if err != nil {
			return ""
		}

		fromStack, err := strconv.Atoi(instructionLine[3])
		if err != nil {
			return ""
		}

		toStack, _ := strconv.Atoi(instructionLine[5])
		if err != nil {
			return ""
		}

		// Converts the stack number to the index in the array
		fromStack = fromStack - 1
		toStack = toStack - 1

		// Executes the custom move function that is provided
		moveCratesFunction(crateToStackMapping, fromStack, toStack, countOfCratesToMove)
	}

	// Creates the output string of the crate at the top of each stack
	output := ""
	for index := range crateToStackMapping {
		value := crateToStackMapping[index].Pop()
		if value != nil {
			output += string(value.(uint8))
		}
	}

	return output
}

// Day5Part1 returns the crates at the top of each stack using the original CrateMover 9000 capabilities
func Day5Part1(input []string) (cratesAtTopOfStack string) {
	return moveCrates(input, func(crateToStackMapping []structures.Stack, from int, to int, count int) {
		for i := 0; i < count; i++ {
			crateToStackMapping[to].Push(crateToStackMapping[from].Pop())
		}
	})
}

// Day5Part2 returns the crates at the top of each stack using the new CrateMover 9001 capabilities
func Day5Part2(input []string) (cratesAtTopOfStacks9001 string) {
	return moveCrates(input, func(crateToStackMapping []structures.Stack, from int, to int, count int) {
		subStack := &structures.Stack{}
		for i := count; i > 0; i-- {
			subStack.Push(crateToStackMapping[from].Pop())
		}

		for i := 0; i < count; i++ {
			crateToStackMapping[to].Push(subStack.Pop())
		}
	})
}
