package advent2022

import (
	"fmt"

	"github.com/marjamis/advent-of-code/pkg/helpers"
	"github.com/marjamis/advent-of-code/pkg/structures"
)

func getValidAdjacentEdges(heightMap [][]rune, currentPosition structures.MatrixPathQueueCoordinates) (adj []structures.MatrixPathQueueCoordinates) {
	surroundingLocations := []structures.MatrixPathQueueCoordinates{
		{
			Col: currentPosition.Col,
			Row: currentPosition.Row - 1,
		},
		{
			Col: currentPosition.Col,
			Row: currentPosition.Row + 1,
		},
		{
			Col: currentPosition.Col + 1,
			Row: currentPosition.Row,
		},
		{
			Col: currentPosition.Col - 1,
			Row: currentPosition.Row,
		},
	}

	currentHeight := int(heightMap[currentPosition.Row][currentPosition.Col])
	// HACK If S, converts to a
	if currentHeight == 83 {
		currentHeight = 97
	}

	for _, location := range surroundingLocations {
		if helpers.IsLocationValid(heightMap, location.Col, location.Row) {
			newHeight := int(heightMap[location.Row][location.Col])
			// HACK If E, converts to z
			if newHeight == 69 {
				newHeight = 122
			}

			if (newHeight - currentHeight) <= 1 {
				adj = append(adj, location)
			}
		}
	}

	return adj
}

func createSetKey(pos structures.MatrixPathQueueCoordinates) string {
	return fmt.Sprintf("%d|%d", pos.Row, pos.Col)
}

func breadthFirstSearch(heightMap [][]rune, startingPosition structures.MatrixPathQueueCoordinates, goal rune) structures.MatrixPathQueueCoordinates {
	// This likely could be generalised in the future for generic BFS implementation
	queue := structures.NewQueue()
	explored := map[string]struct{}{}
	explored[createSetKey(startingPosition)] = struct{}{}

	queue.Enqueue(structures.MatrixPathQueueCoordinates{
		Col: startingPosition.Col,
		Row: startingPosition.Row,
	})

	for !queue.IsEmpty() {
		currentNode := queue.Dequeue()

		if heightMap[currentNode.Row][currentNode.Col] == goal {
			return currentNode
		}

		for _, node := range getValidAdjacentEdges(heightMap, currentNode) {
			key := createSetKey(node)

			if _, found := explored[key]; !found {
				explored[key] = struct{}{}
				node.Parent = &currentNode
				queue.Enqueue(node)
			}
		}
	}

	return structures.MatrixPathQueueCoordinates{}
}

func findPathLength(fromNode structures.MatrixPathQueueCoordinates) (pathLength int) {
	for fromNode.Parent != nil {
		pathLength++
		fromNode = *fromNode.Parent
	}

	return pathLength
}

func findStartingPositions(heightMap [][]rune, startingValue rune) (positions []structures.MatrixPathQueueCoordinates) {
	for row := range heightMap {
		for col := range heightMap[row] {
			if heightMap[row][col] == startingValue {
				positions = append(positions, structures.MatrixPathQueueCoordinates{
					Row: row,
					Col: col,
				})
			}
		}
	}

	return
}

// Day12Part1 returns the shortest from to E from S
func Day12Part1(heightMap [][]rune, goal rune) int {
	startingPosition := findStartingPositions(heightMap, 'S')[0]

	return findPathLength(breadthFirstSearch(heightMap, startingPosition, goal))
}

// Day12Part2 returns the shortest path to E from any a height, this includes S, position
func Day12Part2(heightMap [][]rune, goal rune) int {
	startingPositions := findStartingPositions(heightMap, 'S')
	startingPositions = append(startingPositions, findStartingPositions(heightMap, 'a')...)

	shortestPath := 99999
	for _, startingPosition := range startingPositions {
		pathLength := findPathLength(breadthFirstSearch(heightMap, startingPosition, goal))

		if pathLength < shortestPath && pathLength != 0 {
			shortestPath = pathLength
		}
	}

	return shortestPath
}
