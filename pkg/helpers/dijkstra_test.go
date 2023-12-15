package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNode(t *testing.T) {
	nodes := Nodes{}
	testNodeName := "testNode"

	// Doesn't exist in the map yet
	_, ok := nodes[testNodeName]
	assert.False(t, ok)

	nodes.createNode(testNodeName)

	// Node should now exist
	_, ok = nodes[testNodeName]
	assert.True(t, ok)
}

func TestCreateEdge(t *testing.T) {
	t.Run("Undirected Edge", func(t *testing.T) {
		nodes := Nodes{}
		testEdgeSource := "start"
		testEdgeDestination := "end"

		// Nodes don't exist
		_, ok := nodes[testEdgeSource]
		assert.False(t, ok)
		_, ok = nodes[testEdgeDestination]
		assert.False(t, ok)

		nodes.CreateUndirectedEdge(testEdgeSource, testEdgeDestination)

		// SourceNode and Edge details should be created
		sourceNode, ok := nodes[testEdgeSource]
		assert.True(t, ok)
		assert.Len(t, sourceNode.Edges, 1)
		assert.Equal(t, testEdgeSource, sourceNode.Edges[0].Source.Name)
		assert.Equal(t, testEdgeDestination, sourceNode.Edges[0].Destination.Name)

		// As it's Undirected DestinationNode and a backwards edge should be created
		destinationNode, ok := nodes[testEdgeDestination]
		assert.True(t, ok)
		assert.Len(t, destinationNode.Edges, 1)
		assert.Equal(t, testEdgeDestination, destinationNode.Edges[0].Source.Name)
		assert.Equal(t, testEdgeSource, destinationNode.Edges[0].Destination.Name)
	})

	t.Run("Directed Edge", func(t *testing.T) {
		nodes := Nodes{}
		testEdgeSource := "start"
		testEdgeDestination := "end"

		// Nodes don't exist
		_, ok := nodes[testEdgeSource]
		assert.False(t, ok)
		_, ok = nodes[testEdgeDestination]
		assert.False(t, ok)

		nodes.CreateDirectedEdge(testEdgeSource, testEdgeDestination)

		// SourceNode and Edge details should be created
		sourceNode, ok := nodes[testEdgeSource]
		assert.True(t, ok)
		assert.Len(t, sourceNode.Edges, 1)
		assert.Equal(t, testEdgeSource, sourceNode.Edges[0].Source.Name)
		assert.Equal(t, testEdgeDestination, sourceNode.Edges[0].Destination.Name)

		// As it's Directed the DestinationNode should have no edges
		destinationNode, ok := nodes[testEdgeDestination]
		assert.True(t, ok)
		assert.Len(t, destinationNode.Edges, 0)
	})
}

func TestLoadDataToNodes(t *testing.T) {
	t.Run("Basic Edge with different delimiters", func(t *testing.T) {
		tests := []struct {
			rawData           []string
			delimiter         string
			expectedNodeNames []string
		}{
			{
				[]string{
					"start-A",
					"start-b",
					"A-c",
					"A-b",
					"b-d",
					"A-end",
					"b-end",
				},
				"-",
				[]string{
					"start",
					"A",
					"b",
					"c",
					"d",
					"end",
				},
			},
			{
				[]string{
					"start,A",
					"start,b",
					"A,c",
					"A,b",
					"b,d",
					"A,end",
					"b,end",
				},
				",",
				[]string{
					"start",
					"A",
					"b",
					"c",
					"d",
					"end",
				},
			},
			{
				[]string{
					"start->A",
					"start->b",
					"A->c",
					"A->b",
					"b->d",
					"A->end",
					"b->end",
				},
				"->",
				[]string{
					"start",
					"A",
					"b",
					"c",
					"d",
					"end",
				},
			},
		}

		for _, test := range tests {
			nodes := LoadNodes(test.rawData, test.delimiter)

			// Check the expected Nodes exist
			for _, nodeName := range test.expectedNodeNames {
				_, ok := nodes[nodeName]
				assert.True(t, ok)
			}
		}
	})

	t.Run("Adding a Weight", func(t *testing.T) {
		tests := []struct {
			rawData       []string
			delimiter     string
			expectedEdges []struct {
				Source      string
				Destination string
				Weight      int
			}
		}{
			{
				[]string{
					"start->A->1",
					"start->b->2",
					"A->c->3",
					"A->b->4",
					"b->d->9",
					"A->end->0",
					"b->end->10",
				},
				"->",
				[]struct {
					Source      string
					Destination string
					Weight      int
				}{
					// Main data with weight
					{Source: "start", Destination: "A", Weight: 1},
					{Source: "start", Destination: "b", Weight: 2},
					{Source: "A", Destination: "c", Weight: 3},
					{Source: "A", Destination: "b", Weight: 4},
					{Source: "b", Destination: "d", Weight: 9},
					{Source: "A", Destination: "end", Weight: 0},
					{Source: "b", Destination: "end", Weight: 10},

					// As the default is undirected, ensure the edges are created in the reverse direction
					{Source: "A", Destination: "start", Weight: 1},
					{Source: "b", Destination: "start", Weight: 2},
					{Source: "c", Destination: "A", Weight: 3},
					{Source: "b", Destination: "A", Weight: 4},
					{Source: "d", Destination: "b", Weight: 9},
					{Source: "end", Destination: "A", Weight: 0},
					{Source: "end", Destination: "b", Weight: 10},
				},
			},
		}

		for _, test := range tests {
			nodes := LoadNodes(test.rawData, test.delimiter)

			// Check the expected Edges are created with the correct weight
			for _, testEdge := range test.expectedEdges {
				for _, sourceEdge := range nodes[testEdge.Source].Edges {
					// Require this check to only test the right destination edge
					if sourceEdge.Destination.Name == testEdge.Destination {
						assert.Equal(t, testEdge.Weight, sourceEdge.Weight)
					}
				}
			}
		}
	})

	t.Run("Adding a Weight and Directed", func(t *testing.T) {
		tests := []struct {
			rawData       []string
			delimiter     string
			expectedEdges []struct {
				Source      string
				Destination string
				Weight      int
				Directed    bool
			}
			unexpectedEdges []struct {
				Destination string
				Source      string
			}
		}{
			{
				[]string{
					"start->A->1->true",
					"start->b->2->false",
					"A->c->3->false",
					"A->b->4->true",
					"b->d->9->false",
					"A->end->0->true",
					"b->end->10->true",
				},
				"->",
				[]struct {
					Source      string
					Destination string
					Weight      int
					Directed    bool
				}{
					// Main data specifying direction
					{Source: "start", Destination: "A", Weight: 1, Directed: true},
					{Source: "start", Destination: "b", Weight: 2, Directed: false},
					{Source: "A", Destination: "c", Weight: 3, Directed: false},
					{Source: "A", Destination: "b", Weight: 4, Directed: true},
					{Source: "b", Destination: "d", Weight: 9, Directed: false},
					{Source: "A", Destination: "end", Weight: 0, Directed: true},
					{Source: "b", Destination: "end", Weight: 10, Directed: true},

					// Checking to ensure the backwards direction is added for undirected edges
					{Source: "b", Destination: "start", Weight: 2, Directed: false},
					{Source: "c", Destination: "A", Weight: 3, Directed: false},
					{Source: "d", Destination: "b", Weight: 9, Directed: false},
				},
				[]struct {
					Destination string
					Source      string
				}{
					{Source: "start", Destination: "A"},
					{Source: "A", Destination: "b"},
					{Source: "A", Destination: "end"},
					{Source: "b", Destination: "end"},
				},
			},
		}

		for _, test := range tests {
			nodes := LoadNodes(test.rawData, test.delimiter)

			// Check the expected Edges are created with the correct direction bool
			for _, testEdge := range test.expectedEdges {
				for _, e := range nodes[testEdge.Source].Edges {
					// Require this check to only test the right destination edge
					if e.Destination.Name == testEdge.Destination {
						assert.Equal(t, testEdge.Directed, e.Directed)
					}
				}
			}

			for _, testEdge := range test.unexpectedEdges {
				for _, destinationEdge := range nodes[testEdge.Destination].Edges {
					// No Edge in the destination Node should be back to the source from this list
					// as these are directed edges
					if destinationEdge.Destination.Name == testEdge.Source {
						t.Fail()
					}
				}
			}
		}
	})
}
