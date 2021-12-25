package advent2019

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3Part1(t *testing.T) {
	//Provided Tests
	assert.Equal(t, 6, Day3([]string{"R8", "U5", "L5", "D3"}, []string{"U7", "R6", "D4", "L4"}, Day3Manhattan))
	assert.Equal(t, 159, Day3([]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}, Day3Manhattan))
	assert.Equal(t, 135, Day3([]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}, Day3Manhattan))
}

func TestDay3Part2(t *testing.T) {
	// Provided Tests
	assert.Equal(t, 30, Day3([]string{"R8", "U5", "L5", "D3"}, []string{"U7", "R6", "D4", "L4"}, Day3Steps))
	assert.Equal(t, 610, Day3([]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}, Day3Steps))
	assert.Equal(t, 410, Day3([]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}, Day3Steps))
}
