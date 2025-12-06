package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIngredientsData_ProducesCorrectFormat(test *testing.T) {
	input := ``

	IdRanges, availableIDs := GetIngredientsData(input)

	assert.Equal(test, 0, len(IdRanges), "Expected: 1, Received: %s", len(IdRanges))
	assert.Equal(test, 0, len(availableIDs), "Expected: 1, Received: %s", len(availableIDs))

	input = `12-18

1`

	IdRanges, availableIDs = GetIngredientsData(input)

	assert.Equal(test, 1, len(IdRanges), "Expected: 1, Received: %s", len(IdRanges))
	assert.Equal(test, 1, len(availableIDs), "Expected: 1, Received: %s", len(availableIDs))
	assert.Equal(test, "12-18", IdRanges[0], "Expected: '12-18', Received: %s", IdRanges[0])
	assert.Equal(test, 1, availableIDs[0], "Expected: 1, Received: %s", availableIDs[0])

	input = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	IdRanges, availableIDs = GetIngredientsData(input)

	assert.Equal(test, 4, len(IdRanges), "Expected: 4, Received: %s", len(IdRanges))
	assert.Equal(test, 6, len(availableIDs), "Expected: 6, Received: %s", len(availableIDs))
}

func TestIsNumberInRange_ReturnsFalseForNumNotInRange(test *testing.T) {
	testNum := 1
	testRange := "3-5"

	output := IsNumberInRange(testNum, testRange)

	assert.Equal(test, false, output, "Expected: false, Received: %s", output)

	testNum = 1
	testRange = "10-14"

	output = IsNumberInRange(testNum, testRange)

	assert.Equal(test, false, output, "Expected: false, Received: %s", output)

	testNum = 1
	testRange = "16-20"

	output = IsNumberInRange(testNum, testRange)

	assert.Equal(test, false, output, "Expected: false, Received: %s", output)

	testNum = 1
	testRange = "12-18"

	output = IsNumberInRange(testNum, testRange)

	assert.Equal(test, false, output, "Expected: false, Received: %s", output)
}

func TestIsNumberInRange_ReturnsTrueForNumNotInRange(test *testing.T) {
	testNum := 5
	testRange := "3-5"

	output := IsNumberInRange(testNum, testRange)

	assert.Equal(test, true, output, "Expected: true, Received: %s", output)

	testNum = 11
	testRange = "10-14"

	output = IsNumberInRange(testNum, testRange)

	assert.Equal(test, true, output, "Expected: true, Received: %s", output)

	testNum = 17
	testRange = "16-20"

	output = IsNumberInRange(testNum, testRange)

	assert.Equal(test, true, output, "Expected: true, Received: %s", output)

	testNum = 17
	testRange = "12-18"

	output = IsNumberInRange(testNum, testRange)

	assert.Equal(test, true, output, "Expected: true, Received: %s", output)
}

func TestGetTotalFreshAvailableIngredients_ReturnsTotal(test *testing.T) {
	idRanges, availableIDs := GetIngredientsData(`12-18

1`)

	output := GetTotalFreshAvailableIngredients(idRanges, availableIDs)

	assert.Equal(test, 0, output, "Expected: 0, Received: %s", output)

	idRanges, availableIDs = GetIngredientsData(`12-18

13`)

	output = GetTotalFreshAvailableIngredients(idRanges, availableIDs)

	assert.Equal(test, 1, output, "Expected: 1, Received: %s", output)

	idRanges, availableIDs = GetIngredientsData(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)

	output = GetTotalFreshAvailableIngredients(idRanges, availableIDs)

	assert.Equal(test, 3, output, "Expected: 3, Received: %s", output)
}

func TestPart1_ReturnsTotal_IncludingTestInput(test *testing.T) {
	input := `12-18

1`

	output := Part1(input)

	assert.Equal(test, 0, output, "Expected: 0, Received: %s", output)

	input = `12-18

13`

	output = Part1(input)

	assert.Equal(test, 1, output, "Expected: 1, Received: %s", output)

	input = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	output = Part1(input)

	assert.Equal(test, 3, output, "Expected: 3, Received: %s", output)
}

func TestGetTotalUniqueIDsInRanges_ReturnsUniqueIDsInArray(test *testing.T) {
	idRanges, _ := GetIngredientsData(`1-2

1`)

	output := GetTotalUniqueIDsInRanges(idRanges)

	assert.Equal(test, 2, output, "Expected: 2, Received: %s", output)

	idRanges, _ = GetIngredientsData(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)

	output = GetTotalUniqueIDsInRanges(idRanges)

	assert.Equal(test, 14, output, "Expected: 14, Received: %s", output)
}

func TestPart2_ReturnsTotal_IncludingTestInput(test *testing.T) {
	input := `12-18

1`

	output := Part2(input)

	assert.Equal(test, 7, output, "Expected: 7, Received: %s", output)

	input = `5-7

13`

	output = Part2(input)

	assert.Equal(test, 3, output, "Expected: 3, Received: %s", output)

	input = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	output = Part2(input)

	assert.Equal(test, 14, output, "Expected: 14, Received: %s", output)
}