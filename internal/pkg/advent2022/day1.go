package advent2022

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func calculateCalories(input []string) []int {
	calories := []int{}

	current := 0
	for _, line := range input {
		if strings.Compare(line, "") == 0 {
			calories = append(calories, current)
			current = 0
			continue
		}

		currentItemCalories, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		current += currentItemCalories
	}
	calories = append(calories, current)

	sort.Sort(sort.IntSlice(calories))

	return calories
}

// Day1Part1 returns the total amount of calories from the one elf
func Day1Part1(input []string) int {
	allCalories := calculateCalories(input)

	return allCalories[len(allCalories)-1]
}

// Day1Part2 returns the total amount of calories from the top three elfs
func Day1Part2(input []string) int {
	allCalories := calculateCalories(input)

	return allCalories[len(allCalories)-1] + allCalories[len(allCalories)-2] + allCalories[len(allCalories)-3]
}
