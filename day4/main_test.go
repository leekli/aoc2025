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

func TestFindAllAdjacentRolls_ReturnsRollsTotalWhichCanBeAccessed(test *testing.T) {
	input := StringTo2DArray(`....
....
.@..
....`)

	output, _ := FindAndTrackAllAcessibleRolls(input)

	assert.Equal(test, 1, output, "Expected: 0, Received: %v", output)

	input = StringTo2DArray(`.@@.
@...
.@@.
...@`)

	output, _ = FindAndTrackAllAcessibleRolls(input)

	assert.Equal(test, 6, output, "Expected: 6, Received: %v", output)

	input = StringTo2DArray(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`)

	output, _ = FindAndTrackAllAcessibleRolls(input)

	assert.Equal(test, 13, output, "Expected: 13, Received: %v", output)
}

func TestFindAllAdjacentRolls_TrackingGridTracksLocationsCorrectly(test *testing.T) {
	input := StringTo2DArray(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`)

	_, trackingGrid := FindAndTrackAllAcessibleRolls(input)

	assert.Equal(test, false, trackingGrid[0][1], "Expected: false, Received: %v", trackingGrid[0][1])
	assert.Equal(test, true, trackingGrid[0][2], "Expected: true, Received: %v", trackingGrid[0][2])
	assert.Equal(test, true, trackingGrid[1][0], "Expected: true, Received: %v", trackingGrid[1][0])
	assert.Equal(test, true, trackingGrid[4][9], "Expected: true, Received: %v", trackingGrid[4][9])
	assert.Equal(test, false, trackingGrid[6][0], "Expected: false, Received: %v", trackingGrid[6][0])
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

func TestCreateTrackingGrid_CreatesTrackingGridAllSetToFalse(test *testing.T) {
	input := StringTo2DArray(`....
....`)

	outputGrid := CreateTrackingGrid(input)

	assert.Equal(test, 2, len(outputGrid), "Expected: 2, Received: %v", outputGrid)
	assert.Equal(test, 4, len(outputGrid[0]), "Expected: 4, Received: %v", outputGrid)
	assert.Equal(test, false, outputGrid[0][0], "Expected: false, Received: %v", outputGrid)
	assert.Equal(test, false, outputGrid[0][1], "Expected: false, Received: %v", outputGrid)
	assert.Equal(test, false, outputGrid[0][2], "Expected: false, Received: %v", outputGrid)
	assert.Equal(test, false, outputGrid[1][0], "Expected: false, Received: %v", outputGrid)
	assert.Equal(test, false, outputGrid[1][1], "Expected: false, Received: %v", outputGrid)
	assert.Equal(test, false, outputGrid[1][3], "Expected: false, Received: %v", outputGrid)
}

func TestPart2_ReturnsRollsTotalWhichCanBeAccessed(test *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	output := Part2(input)

	assert.Equal(test, 43, output, "Expected: 43, Received: %v", output)
}