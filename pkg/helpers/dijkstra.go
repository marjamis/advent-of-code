package helpers

import (
	"log"
	"strconv"
	"strings"
)

// Edge contains data about the connection of two Nodes
type Edge struct {
	Weight      int
	Source      Node
	Destination Node
	Directed    bool
}

// Node contains details about a specific Node in the graph
type Node struct {
	Name  string
	Edges []Edge
}

// Nodes a map of Nodes in the graph
type Nodes map[string]Node

func (nodes Nodes) createNode(nodeName string) {
	_, ok := nodes[nodeName]
	if !ok {
		nodes[nodeName] = Node{
			Name: nodeName,
		}
	}
}

func (nodes Nodes) createEdge(source, destination string, directed bool) {
	_, ok := nodes[source]
	if !ok {
		nodes.createNode(source)
	}

	_, ok = nodes[destination]
	if !ok {
		nodes.createNode(destination)
	}

	// Creates the Edge and adds it to the source Edge list
	tmpNode := nodes[source]
	edge := Edge{
		Source:      nodes[source],
		Destination: nodes[destination],
		Directed:    directed,
	}
	tmpNode.Edges = append(tmpNode.Edges, edge)
	nodes[source] = tmpNode
}

// createDirectedEdge creates an Edge object and assigns it to the source node
func (nodes Nodes) createDirectedEdge(source, destination string) {
	nodes.createEdge(source, destination, true)
}

// createDirectedEdge creates an Edge object and assigns it to the source and destination node
// as the traversal can go in both directions
func (nodes Nodes) createUndirectedEdge(source, destination string) {
	nodes.createEdge(source, destination, false)
	nodes.createEdge(destination, source, false)
}

// LoadNodes takes raw data for a graph and converts it into the Nodes map
// Format of a rawData string separated by the delimiter is "<From><To>[Weight][Directed]"
// Last two are optional but Directed requires the Weight
func LoadNodes(rawData []string, delimiter string) (nodes Nodes) {
	nodes = make(map[string]Node)

	for _, nodeEdges := range rawData {
		input := strings.Split(nodeEdges, delimiter)
		sourceNode := input[0]
		destinationNode := input[1]

		// If the optional fourth field is provided it's used to determine if it's a directed
		// graph or not. If this fourth field is not provided then the default is an undirected graph
		directed := false
		if len(input) == 4 {
			var err error
			directed, err = strconv.ParseBool(input[3])
			if err != nil {
				log.Fatal("Error connverting string to int...")
			}

			if directed {
				nodes.createDirectedEdge(sourceNode, destinationNode)
			} else {
				nodes.createUndirectedEdge(sourceNode, destinationNode)
			}
		} else {
			// If nodes don't already exist they will also be created
			nodes.createUndirectedEdge(sourceNode, destinationNode)
		}

		// If the optional third field is provided this is used as the weight for the edge for both
		// directed or undirected graphs
		if len(input) >= 3 {
			weight, err := strconv.Atoi(input[2])
			if err != nil {
				log.Fatal("Error connverting string to int...")
			}

			nodes[sourceNode].Edges[len(nodes[sourceNode].Edges)-1].Weight = weight

			// If the edge isn't directed then the weight is added for the back direction as well
			if !directed {
				nodes[destinationNode].Edges[len(nodes[destinationNode].Edges)-1].Weight = weight
			}
		}
	}

	return
}
