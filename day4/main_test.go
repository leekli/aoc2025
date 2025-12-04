package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringTo2DArray_ProducesCorrect2DNumArray(test *testing.T) {
	input := `..@@
@@@.
@@@.
@.@.`

	output := StringTo2DArray(input)

	assert.Equal(test, ".", output[0][0], "Expected: '.', Received: %s", output[0][0])
	assert.Equal(test, "@", output[1][0], "Expected: '@', Received: %s", output[1][0])
	assert.Equal(test, "@", output[2][0], "Expected: '@', Received: %s", output[2][0])
	assert.Equal(test, "@", output[0][3], "Expected: '@', Received: %s", output[0][3])
	assert.Equal(test, "@", output[2][2], "Expected: '@', Received: %s", output[2][2])
}

func TestIsPaperRollFound_ReturnsFalseForRollNotFound(test *testing.T) {
	input := "."

	output := IsPaperRollFound(input)

	assert.Equal(test, false, output, "Expected: false, Received: %v", output)
}

func TestIsPaperRollFound_ReturnsTrueForRollNotFound(test *testing.T) {
	input := "@"

	output := IsPaperRollFound(input)

	assert.Equal(test, true, output, "Expected: false, Received: %v", output)
}

func TestPart1_ReturnsRollsTotalWhichCanBeAccessed(test *testing.T) {
	input := `....
....
.@..
....`

	output := Part1(input)

	assert.Equal(test, 1, output, "Expected: 0, Received: %v", output)

	input = `.@@.
@...
.@@.
...@`

	output = Part1(input)

	assert.Equal(test, 6, output, "Expected: 6, Received: %v", output)

	input = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	output = Part1(input)

	assert.Equal(test, 13, output, "Expected: 13, Received: %v", output)
}