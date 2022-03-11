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

func TestFilterUpdatedSmallCavesFromPath(t *testing.T) {
	assert.False(t, filterUpdatedSmallCavesFromPath([]string{"A", "c"}, "c"))
	assert.False(t, filterUpdatedSmallCavesFromPath([]string{"c", "A"}, "A"))
	assert.True(t, filterUpdatedSmallCavesFromPath([]string{
		"start", "A", "c", "A", "c", "A", "b", "A"}, "b"))
}

func TestHasSmallCaveDoubleVisit(t *testing.T) {
	assert.False(t, hasSmallCaveDoubleVisit([]string{"A", "c"}))
	assert.True(t, hasSmallCaveDoubleVisit([]string{"A", "c", "c"}))
	assert.False(t, hasSmallCaveDoubleVisit([]string{"A", "A", "A"}))
}

func TestCavesToVisitP1(t *testing.T) {
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

	assert.ElementsMatch(t, []string{"c", "end"}, cavesToVisit(testCaves, Filters{
		start:      filterStart,
		smallCaves: filterSmallCavesFromPath,
	}, []string{
		"start", "A", "b", "A",
	}))

	assert.ElementsMatch(t, []string{"A"}, cavesToVisit(testCaves, Filters{
		start:      filterStart,
		smallCaves: filterSmallCavesFromPath,
	}, []string{
		"start", "A", "b", "A", "c",
	}))

	assert.ElementsMatch(t, []string{"end"}, cavesToVisit(testCaves, Filters{
		start:      filterStart,
		smallCaves: filterSmallCavesFromPath,
	}, []string{
		"start", "A", "b", "A", "c", "A",
	}))
}

func TestTraverseP1(t *testing.T) {
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
		assert.Equal(t, test.expected, traverse(test.input, Filters{
			start:      filterStart,
			smallCaves: filterSmallCavesFromPath,
		}, "start", []string{}))
	}
}

func TestTraverseP2(t *testing.T) {
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
			36,
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
			103,
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
			3509,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, traverse(test.input, Filters{
			start:      filterStart,
			smallCaves: filterUpdatedSmallCavesFromPath,
		}, "start", []string{}))
	}
}
