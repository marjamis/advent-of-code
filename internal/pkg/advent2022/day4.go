package advent2022

import (
	"strconv"
	"strings"
)

func countPairsWithValidationFunction(sectionAssignments []string, validationFunction func(int, int, int, int) bool) (countOfPairs int) {
	for _, pair := range sectionAssignments {
		pairSplit := strings.Split(pair, ",")

		listOne := pairSplit[0]
		listTwo := pairSplit[1]

		oneSplit := strings.Split(listOne, "-")
		twoSplit := strings.Split(listTwo, "-")

		oneLower, err := strconv.Atoi(oneSplit[0])
		if err != nil {
			return -1
		}

		oneUpper, err := strconv.Atoi(oneSplit[1])
		if err != nil {
			return -1
		}

		twoLower, err := strconv.Atoi(twoSplit[0])
		if err != nil {
			return -1
		}

		twoUpper, err := strconv.Atoi(twoSplit[1])
		if err != nil {
			return -1
		}

		if validationFunction(oneLower, oneUpper, twoLower, twoUpper) {
			countOfPairs++
		}
	}

	return
}

// Day4Part1 returns a count of pairs that have sections that are fully contained
func Day4Part1(sectionAssignments []string) (countOfFullyContainedPairs int) {
	return countPairsWithValidationFunction(sectionAssignments, func(oneLower, oneUpper, twoLower, twoUpper int) bool {
		if ((oneLower <= twoLower) && (oneUpper >= twoUpper)) || ((twoLower <= oneLower) && (twoUpper >= oneUpper)) {
			return true
		}

		return false
	})
}

// Day4Part2 returns count of all pairs with any overlap
func Day4Part2(sectionAssignments []string) (countOfAnyOverlappedPairs int) {
	return countPairsWithValidationFunction(sectionAssignments, func(oneLower, oneUpper, twoLower, twoUpper int) bool {
		if ((oneLower <= twoUpper) && (oneLower >= twoLower)) || ((twoLower <= oneUpper) && (twoLower >= oneLower)) {
			return true
		}

		return false
	})
}
