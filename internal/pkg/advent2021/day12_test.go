package advent2021

import (
	"testing"

	"github.com/marjamis/advent-of-code/pkg/helpers"
	"github.com/stretchr/testify/assert"
)

func TestFilterStart(t *testing.T) {
	assert.True(t, filterStart("start"))
	assert.False(t, filterStart("star"))
}

func TestFilterSmallCavesFromPath(t *testing.T) {
	assert.True(t, filterSmallCavesFromPath([]string{"A", "c"}, "c"))
	assert.False(t, filterSmallCavesFromPath([]string{"c", "A"}, "A"))
}

func TestCanVisit(t *testing.T) {
	testCaves := Caves(helpers.LoadNodes(
		[]string{
			"start-A",
			"start-b",
			"A-c",
			"A-b",
			"b-d",
			"A-end",
			"b-end",
		},
		"-"))

	assert.ElementsMatch(t, []string{"c", "end"}, testCaves.cavesToVisit([]string{
		"start", "A", "b", "A",
	}))

	assert.ElementsMatch(t, []string{"A"}, testCaves.cavesToVisit([]string{
		"start", "A", "b", "A", "c",
	}))

	assert.ElementsMatch(t, []string{"end"}, testCaves.cavesToVisit([]string{
		"start", "A", "b", "A", "c", "A",
	}))
}

func TestTraverse(t *testing.T) {
	tests := []struct {
		input    Caves
		expected int
	}{
		{
			Caves(helpers.LoadNodes(
				[]string{
					"start-A",
					"start-b",
					"A-c",
					"A-b",
					"b-d",
					"A-end",
					"b-end",
				},
				"-")),
			10,
		},
		{
			Caves(helpers.LoadNodes(
				[]string{
					"dc-end",
					"HN-start",
					"start-kj",
					"dc-start",
					"dc-HN",
					"LN-dc",
					"HN-end",
					"kj-sa",
					"kj-HN",
					"kj-dc",
				},
				"-")),
			19,
		},
		{
			Caves(helpers.LoadNodes(
				[]string{
					"fs-end",
					"he-DX",
					"fs-he",
					"start-DX",
					"pj-DX",
					"end-zg",
					"zg-sl",
					"zg-pj",
					"pj-he",
					"RW-he",
					"fs-DX",
					"pj-RW",
					"zg-RW",
					"start-pj",
					"he-WI",
					"zg-he",
					"pj-fs",
					"start-RW",
				},
				"-",
			)),
			226,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.input.traverse("start", []string{}))
	}
}
