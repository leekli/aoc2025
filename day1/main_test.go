package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertInstructionsToList_ReturnsEmptySliceForEmptyInstructions(t *testing.T) {
	input := ""

	output := ConvertInstructionsToList(input)
	
	outputLength := len(output)

	assert.Equal(t, 1, outputLength, "Output length should be 1")
}

func TestConvertInstructionsToList_ReturnsSingleSliceForSingleInstruction(t *testing.T) {
	input := "L68"

	output := ConvertInstructionsToList(input)
	outputLength := len(output)
	firstExpected := "L68"

	assert.Equal(t, 1, outputLength, "Output length should be 1")
	assert.Equal(t, firstExpected, output[0], "Index 1 is incorrect")
}

func TestConvertInstructionsToList_ReturnsSliceForMultiLineInstructions(t *testing.T) {
	input := `L68
L30
R48`

	output := ConvertInstructionsToList(input)
	outputLength := len(output)
	firstExpected := "L68"
	secondExpected := "L30"
	thirExpected := "R48"

	assert.Equal(t, 3, outputLength, "Output length should be 3")
	assert.Equal(t, firstExpected, output[0], "Index 1 is incorrect")
	assert.Equal(t, secondExpected, output[1], "Index 2 is incorrect")
	assert.Equal(t, thirExpected, output[2], "Index 3 is incorrect")
}

func TestMakeMove_ShouldMoveLeft_NoMinNumCycleNeeded(t *testing.T) {
	direction := "L"
	numOfMoves := 2
	currentPosition := 10
	minNum := 0
	maxNum := 99

	newLocation := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 8, newLocation, "New location incorrect")
}

func TestMakeMove_ShouldMoveRight_NoMinNumCycleNeeded(t *testing.T) {
	direction := "R"
	numOfMoves := 4
	currentPosition := 6
	minNum := 0
	maxNum := 99

	newLocation := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 10, newLocation, "New location incorrect")
}

func TestMakeMove_ShouldMoveLeft_HitsMinNum_ShouldMoveToMaxNumCycleRound(t *testing.T) {
	direction := "L"
	numOfMoves := 3
	currentPosition := 2
	minNum := 0
	maxNum := 99

	newLocation := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 99, newLocation, "New location incorrect")
}

func TestMakeMove_ShouldMoveLeft_HitsMinNum_ShouldMoveToMaxNum_AndContinueWithLargerNum(t *testing.T) {
	direction := "L"
	numOfMoves := 5
	currentPosition := 2
	minNum := 0
	maxNum := 99

	newLocation := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 97, newLocation, "New location incorrect")
}

func TestMakeMove_ShouldMoveRight_HitsMaxNum_ShouldMoveToMinNumCycleRound(t *testing.T) {
	direction := "R"
	numOfMoves := 3
	currentPosition := 96
	minNum := 0
	maxNum := 99

	newLocation := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 99, newLocation, "New location incorrect")
}

func TestMakeMove_ShouldMoveRight_HitsMaxNum_ShouldMoveToMinNum_AndContinueWithLargerNum(t *testing.T) {
	direction := "R"
	numOfMoves := 5
	currentPosition := 97
	minNum := 0
	maxNum := 99

	newLocation := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 2, newLocation, "New location incorrect")
}

func TestPart1_ShouldRunInstructions_SingleInstruction_NoZerosHit(t *testing.T) {
	input := "L68"

	output := Part1(input)

	assert.Equal(t, 0, output, "Incorrect password output")
}

func TestPart1_ShouldRunInstructions_MultipleInstructions_ZeroHitOnce(t *testing.T) {
	input := `L68
L30
R48`

	output := Part1(input)

	assert.Equal(t, 1, output, "Incorrect password output")
}

func TestPart1_ShouldRunInstructions_WorksWithPuzzleTestFullInput(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	output := Part1(input)

	assert.Equal(t, 3, output, "Incorrect password output")
}