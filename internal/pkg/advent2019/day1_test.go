package advent2019

import (
	"bufio"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	assert := assert.New(t)

	fdata, _ := os.Open("../../../test/advent2019/Day1Data.txt")
	defer fdata.Close()
	scanner := bufio.NewScanner(fdata)
	scanner.Split(bufio.ScanLines)
	var ints []int
	for scanner.Scan() {
		d, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, d)
	}

	//Part 1 - https://adventofcode.com/2019/day/1
	assert.Equal(2, Day1(false, 12))
	assert.Equal(2, Day1(false, 14))
	assert.Equal(654, Day1(false, 1969))
	assert.Equal(33583, Day1(false, 100756))

	//Part 2 - https://adventofcode.com/2019/day/1#part2
	assert.Equal(2, Day1(true, 14))
	assert.Equal(966, Day1(true, 1969))
	assert.Equal(50346, Day1(true, 100756))
}
