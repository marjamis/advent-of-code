package advent2023

import (
	"strings"

	"github.com/marjamis/advent-of-code/pkg/helpers"
)

func createNodeMap(nodeData []string) helpers.Nodes {
	nodes := helpers.CreateNodeMap()

	for _, n := range nodeData {
		s := strings.Split(n, " ")
		source := s[0]
		left := s[2][1:4]
		right := s[3][:3]

		nodes.CreateDirectedEdge(source, left)
		nodes.CreateDirectedEdge(source, right)
	}

	return nodes

}

// Day8Part1 returns the steps taken from the start to end node
func Day8Part1(graphData []string) (stepsTaken int) {
	instructions := graphData[0]
	nodeData := graphData[2:]
	nodes := createNodeMap(nodeData)

	startingNode := "AAA"
	for startingNode != "ZZZ" {
		for _, instruction := range instructions {
			if instruction == 'L' {
				startingNode = nodes[startingNode].Edges[0].Destination.Name
			} else {
				startingNode = nodes[startingNode].Edges[1].Destination.Name
			}
			stepsTaken++

			if startingNode == "ZZZ" {
				break
			}
		}
	}

	return
}

func filterOnlyNodeNamesEndingWithA(nodes helpers.Nodes) map[string]int {
	filteredNodes := map[string]int{}

	for nodeName := range nodes {
		if nodeName[2] == 'A' {
			filteredNodes[nodeName] = 1
		}
	}

	return filteredNodes
}

func nodeNameEndsWith(nodeName string, endsWith rune) bool {
	return nodeName[2] == byte(endsWith)
}

// Day8Part2 returns the steps taken from all start nodes to all end nodes at the same time
func Day8Part2(graphData []string) int {
	instructions := graphData[0]
	nodeData := graphData[2:]
	nodes := createNodeMap(nodeData)
	nodesEndingWithA := filterOnlyNodeNamesEndingWithA(nodes)

	for startingNode := range nodesEndingWithA {
		currentNode := startingNode
		stepsTaken := 0
		for !nodeNameEndsWith(currentNode, 'Z') {
			for _, instruction := range instructions {
				if instruction == 'L' {
					currentNode = nodes[currentNode].Edges[0].Destination.Name
				} else {
					currentNode = nodes[currentNode].Edges[1].Destination.Name
				}
				stepsTaken++

				if nodeNameEndsWith(currentNode, 'Z') {
					nodesEndingWithA[startingNode] = stepsTaken
					break
				}
			}
		}
	}

	nodesSteps := make([]int, len(nodesEndingWithA))
	j := 0
	for _, v := range nodesEndingWithA {
		nodesSteps[j] = v
		j++
	}

	return helpers.LCM(nodesSteps[0], nodesSteps[1], nodesSteps[2:]...)
}
