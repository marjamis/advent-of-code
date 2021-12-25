package advent2019

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(9, Day4("10-100", Day4Rules1))
	assert.Equal(8, Day4("10-200", Day4Rules2))
}
