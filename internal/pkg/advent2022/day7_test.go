package advent2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day7ProvidedTestInput = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

func TestDay7Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day7ProvidedTestInput,
			95437,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day7Part1(test.input))
	}
}

func TestDay7Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			day7ProvidedTestInput,
			24933642,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day7Part2(test.input))
	}
}
