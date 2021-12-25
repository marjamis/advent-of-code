package advent2019

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntcodeComputer(t *testing.T) {
	assert := assert.New(t)

	//Provided Tests
	assert.ElementsMatch([]int{2, 0, 0, 0, 99}, IntcodeComputer(1, 0, 0, 0, 99))
	assert.ElementsMatch([]int{2, 3, 0, 6, 99}, IntcodeComputer(2, 3, 0, 3, 99))
	assert.ElementsMatch([]int{2, 4, 4, 5, 99, 9801}, IntcodeComputer(2, 4, 4, 5, 99, 0))
	assert.ElementsMatch([]int{30, 1, 1, 4, 2, 5, 6, 0, 99}, IntcodeComputer(1, 1, 1, 4, 99, 5, 6, 0, 99))
	assert.ElementsMatch([]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, IntcodeComputer(1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50))

	//My tests to build it up
	assert.ElementsMatch([]int{2, 5, 6, 28, 99, 4, 7}, IntcodeComputer(2, 5, 6, 3, 99, 4, 7))
	assert.ElementsMatch([]int{1, 5, 6, 11, 99, 4, 7}, IntcodeComputer(1, 5, 6, 3, 99, 4, 7))
	assert.ElementsMatch([]int{1, 9, 10, 28, 2, 9, 10, 3, 99, 4, 7}, IntcodeComputer(1, 9, 10, 3, 2, 9, 10, 3, 99, 4, 7))
}
