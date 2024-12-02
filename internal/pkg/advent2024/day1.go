package advent2024

import (
	"sort"
	"strconv"
	"strings"

	"github.com/marjamis/advent-of-code/pkg/helpers"
)

// Day1Part1 returns the difference between the sorted location lists
func Day1Part1(lists []string) (difference int) {
	length := len(lists)
	firstList := make([]int, length)
	secondList := make([]int, length)

	for i := 0; i < length; i++ {
		s := strings.Split(lists[i], "   ")
		var err error
		firstList[i], err = strconv.Atoi(s[0])
		if err != nil {
			panic("Failed to convert to an int")

		}
		secondList[i], err = strconv.Atoi(s[1])
		if err != nil {
			panic("Failed to convert to an int")

		}
	}

	sort.Ints(firstList)
	sort.Ints(secondList)

	for i := 0; i < length; i++ {
		difference += helpers.Abs(firstList[i] - secondList[i])
	}

	return difference
}

// Day1Part2 returns the similarity score
func Day1Part2(lists []string) (score int) {
	length := len(lists)
	firstList := make([]int, length)
	secondListMap := map[int]int{}

	for i := 0; i < length; i++ {
		s := strings.Split(lists[i], "   ")
		var err error
		firstList[i], err = strconv.Atoi(s[0])
		if err != nil {
			panic("Failed to convert to an int")

		}
		secondVal, err := strconv.Atoi(s[1])
		if err != nil {
			panic("Failed to convert to an int")

		}
		secondListMap[secondVal]++
	}

	for i := 0; i < length; i++ {
		score += firstList[i] * secondListMap[firstList[i]]
	}

	return score
}
