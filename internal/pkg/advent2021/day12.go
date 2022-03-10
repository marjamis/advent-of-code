package advent2021

import (
	"strings"

	"github.com/marjamis/advent-of-code/pkg/helpers"
)

// Cave is an individual Node of the graph
type Cave helpers.Node

// Caves is a map of all the caves in the puzzle
type Caves helpers.Nodes

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

func (caves Caves) cavesToVisit(pathSoFar []string) (validNextCaves []string) {
	// Get a list of all the current caves (last element in pathSoFar) connections
	// for the beginnings of next cave to explore
	for _, i := range caves[pathSoFar[len(pathSoFar)-1]].Edges {
		validNextCaves = append(validNextCaves, i.Destination.Name)
	}

	// Exclude the start cave, if it's a possibility
	// TODO make it a filter
	remove := []int{}
	for i, possibility := range validNextCaves {
		if filterStart(possibility) {
			remove = append(remove, i)
		}
	}
	validNextCaves = helpers.RemoveItemsAtIndexes(validNextCaves, remove)

	// Filter out small caves that are already visited, i.e. already in the pathSoFar
	// TODO make it a filter
	remove = []int{}
	for i, possibility := range validNextCaves {
		if filterSmallCavesFromPath(pathSoFar, possibility) {
			remove = append(remove, i)
		}
	}
	validNextCaves = helpers.RemoveItemsAtIndexes(validNextCaves, remove)

	return validNextCaves
}

func (caves Caves) traverse(nextCave string, pathSoFar []string) (totalPaths int) {
	pathSoFar = append(pathSoFar, nextCave)
	if nextCave == "end" {
		// If this next cave is the end then it's a valid path to end on
		totalPaths++
	} else {
		// Find all caves that can be visited and then continually traverse these until they fail
		// or they find the end cave
		for _, caveToVisit := range caves.cavesToVisit(pathSoFar) {
			totalPaths += caves.traverse(caveToVisit, pathSoFar)
		}
	}

	return totalPaths
}

// Day12Part1 returns the number of paths from the start to the end based on the rules
func Day12Part1(rawData []string) int {
	caves := Caves(helpers.LoadNodes(rawData, "-"))

	return caves.traverse("start", []string{})
}
