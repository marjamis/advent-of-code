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

	for _, cave := range pathSoFar {
		if strings.Compare(currentCave, cave) == 0 {
			return true
		}
	}

	return false
}

func hasSmallCaveDoubleVisit(pathSoFar []string) bool {
	caveVisits := map[string]int{}
	for _, cave := range pathSoFar {
		// Only applies to small caves (hence lower case)
		if cave == strings.ToLower(cave) {
			_, ok := caveVisits[cave]
			if !ok {
				caveVisits[cave] = 1
			} else {
				// If the key is found this means its already been visited and there is a double visit
				return true
			}
		}
	}

	return false
}

func filterUpdatedSmallCavesFromPath(pathSoFar []string, currentCave string) bool {
	// If it's a big cave then it remains a valid path
	if strings.ToUpper(currentCave) == currentCave {
		return false
	}

	// Find if there are already any double visits for small caves
	doubles := hasSmallCaveDoubleVisit(pathSoFar)

	// If there is a double already then this cave needs checking if it would also be a double visit
	// if there aren't any doubles then this currentCave is valid already
	// can be visited without further checks
	if doubles {
		// Loops through the path for the currentCave and if a match is found this means its a invalid
		// path as it will be a second double
		for _, cave := range pathSoFar {
			if strings.Compare(currentCave, cave) == 0 {
				return true
			}
		}
	}

	return false
}

func cavesToVisit(caves Caves, filters Filters, pathSoFar []string) (validNextCaves []string) {
	// Get a list of all the current caves (last element in pathSoFar) connections
	// for the beginnings of next cave to explore
	for _, edge := range caves[pathSoFar[len(pathSoFar)-1]].Edges {
		validNextCaves = append(validNextCaves, edge.Destination.Name)
	}

	// Exclude the start cave, if it's a possibility
	// TODO make it a filter
	remove := []int{}
	for i, nextCave := range validNextCaves {
		if filters.start(nextCave) {
			remove = append(remove, i)
		}
	}
	validNextCaves = helpers.RemoveItemsAtIndexes(validNextCaves, remove)

	// Filter out small caves that are already visited, i.e. already in the pathSoFar
	// TODO make it a filter
	remove = []int{}
	for i, nextCave := range validNextCaves {
		if filters.smallCaves(pathSoFar, nextCave) {
			remove = append(remove, i)
		}
	}
	validNextCaves = helpers.RemoveItemsAtIndexes(validNextCaves, remove)

	return validNextCaves
}

func traverse(caves Caves, filters Filters, nextCave string, pathSoFar []string) (totalPaths int) {
	pathSoFar = append(pathSoFar, nextCave)
	if nextCave == "end" {
		// If this next cave is the end then it's a valid path to end on
		totalPaths++
	} else {
		// Find all caves that can be visited and then continually traverse these until they fail
		// or they find the end cave
		for _, caveToVisit := range cavesToVisit(caves, filters, pathSoFar) {
			totalPaths += traverse(caves, filters, caveToVisit, pathSoFar)
		}
	}

	return totalPaths
}

// Day12Part1 returns the number of paths from the start to the end based on the rules
func Day12Part1(rawData []string) int {
	caves := Caves(helpers.LoadNodes(rawData, "-"))

	return traverse(caves, Filters{
		start:      filterStart,
		smallCaves: filterSmallCavesFromPath,
	}, "start", []string{})
}

// Day12Part2 returns the number of paths from the start to the end based on the updated rules
func Day12Part2(rawData []string) int {
	caves := Caves(helpers.LoadNodes(rawData, "-"))

	return traverse(caves, Filters{
		start:      filterStart,
		smallCaves: filterUpdatedSmallCavesFromPath,
	}, "start", []string{})
}
