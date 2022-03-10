package advent2021

import (
	"strings"

	"github.com/marjamis/advent-of-code/pkg/helpers"
)

// Cave is an individual Node of the graph
type Cave helpers.Node

// Caves is a map of all the caves in the puzzle
type Caves helpers.Nodes

// Filters is a struct to pass around functions signatures for the filters to allow for a
// generic implementation while calling specific filters
type Filters struct {
	start      func(string) bool
	smallCaves func([]string, string) bool
}

func filterStart(caveName string) bool {
	if caveName == "start" {
		return true
	}
	return false
}

func filterSmallCavesFromPath(pathSoFar []string, currentCave string) bool {
	// If it's a big cave then it remains a valid path
	if strings.ToUpper(currentCave) == currentCave {
		return false
	}

	for _, path := range pathSoFar {
		if strings.Compare(currentCave, path) == 0 {
			return true
		}
	}

	return false
}

func filterUpdatedSmallCavesFromPath(pathSoFar []string, currentCave string) bool {
	// If it's a big cave then it remains a valid path
	if strings.ToUpper(currentCave) == currentCave {
		return false
	}

	// Allow one small cave to have two visits but all others one. Filter any that try for a second

	return false
}

func cavesToVisit(filters Filters, caves Caves, pathSoFar []string) (validNextCaves []string) {
	// Get a list of all the current caves (last element in pathSoFar) connections
	// for the beginnings of next cave to explore
	for _, i := range caves[pathSoFar[len(pathSoFar)-1]].Edges {
		validNextCaves = append(validNextCaves, i.Destination.Name)
	}

	// Exclude the start cave, if it's a possibility
	// TODO make it a filter
	remove := []int{}
	for i, possibility := range validNextCaves {
		if filters.start(possibility) {
			remove = append(remove, i)
		}
	}
	validNextCaves = helpers.RemoveItemsAtIndexes(validNextCaves, remove)

	// Filter out small caves that are already visited, i.e. already in the pathSoFar
	// TODO make it a filter
	remove = []int{}
	for i, possibility := range validNextCaves {
		if filters.smallCaves(pathSoFar, possibility) {
			remove = append(remove, i)
		}
	}
	validNextCaves = helpers.RemoveItemsAtIndexes(validNextCaves, remove)

	return validNextCaves
}

func traverse(caves Caves, nextCave string, pathSoFar []string, filters Filters) (totalPaths int) {
	pathSoFar = append(pathSoFar, nextCave)
	if nextCave == "end" {
		// If this next cave is the end then it's a valid path to end on
		totalPaths++
	} else {
		// Find all caves that can be visited and then continually traverse these until they fail
		// or they find the end cave
		for _, caveToVisit := range cavesToVisit(filters, caves, pathSoFar) {
			totalPaths += traverse(caves, caveToVisit, pathSoFar, filters)
		}
	}

	return totalPaths
}

// Day12Part1 returns the number of paths from the start to the end based on the rules
func Day12Part1(rawData []string) int {
	caves := Caves(helpers.LoadNodes(rawData, "-"))

	return traverse(caves, "start", []string{}, Filters{
		start:      filterStart,
		smallCaves: filterSmallCavesFromPath,
	})
}

// Day12Part2 returns the number of paths from the start to the end based on the updated rules
func Day12Part2(rawData []string) int {
	caves := Caves(helpers.LoadNodes(rawData, "-"))

	return traverse(caves, "start", []string{}, Filters{
		start:      filterStart,
		smallCaves: filterUpdatedSmallCavesFromPath,
	})
}
