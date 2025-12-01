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

	newLocation, _ := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 8, newLocation, "New location incorrect")
}

func TestMakeMove_ShouldMoveRight_NoMinNumCycleNeeded(t *testing.T) {
	direction := "R"
	numOfMoves := 4
	currentPosition := 6
	minNum := 0
	maxNum := 99

	newLocation, _ := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 10, newLocation, "New location incorrect")
}

func TestMakeMove_ShouldMoveLeft_HitsMinNum_ShouldMoveToMaxNumCycleRound(t *testing.T) {
	direction := "L"
	numOfMoves := 3
	currentPosition := 2
	minNum := 0
	maxNum := 99

	newLocation, _ := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 99, newLocation, "New location incorrect")
}

func TestMakeMove_ShouldMoveLeft_HitsMinNum_ShouldMoveToMaxNum_AndContinueWithLargerNum(t *testing.T) {
	direction := "L"
	numOfMoves := 5
	currentPosition := 2
	minNum := 0
	maxNum := 99

	newLocation, _ := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 97, newLocation, "New location incorrect")
}

func TestMakeMove_ShouldMoveRight_HitsMaxNum_ShouldMoveToMinNumCycleRound(t *testing.T) {
	direction := "R"
	numOfMoves := 3
	currentPosition := 96
	minNum := 0
	maxNum := 99

	newLocation, _ := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 99, newLocation, "New location incorrect")
}

func TestMakeMove_ShouldMoveRight_HitsMaxNum_ShouldMoveToMinNum_AndContinueWithLargerNum(t *testing.T) {
	direction := "R"
	numOfMoves := 5
	currentPosition := 97
	minNum := 0
	maxNum := 99

	newLocation, _ := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 2, newLocation, "New location incorrect")
}

func TestMakeMove_ShouldReturnNumOfZeroClicks_NoneWhenMovingRight(t *testing.T) {
	direction := "R"
	numOfMoves := 1
	currentPosition := 97
	minNum := 0
	maxNum := 99

	_, zeroClicksCount := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 0, zeroClicksCount, "Incorrect zero count")
}

func TestMakeMove_ShouldReturnNumOfZeroClicks_OneWhenMovingRight(t *testing.T) {
	direction := "R"
	numOfMoves := 3
	currentPosition := 97
	minNum := 0
	maxNum := 99

	_, zeroClicksCount := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 1, zeroClicksCount, "Incorrect zero count")
}

func TestMakeMove_ShouldReturnNumOfZeroClicks_NoneWhenMovingLeft(t *testing.T) {
	direction := "L"
	numOfMoves := 1
	currentPosition := 3
	minNum := 0
	maxNum := 99

	_, zeroClicksCount := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 0, zeroClicksCount, "Incorrect zero count")
}

func TestMakeMove_ShouldReturnNumOfZeroClicks_OneWhenMovingLeft(t *testing.T) {
	direction := "L"
	numOfMoves := 3
	currentPosition := 1
	minNum := 0
	maxNum := 99

	_, zeroClicksCount := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 1, zeroClicksCount, "Incorrect zero count")
}

func TestMakeMove_ShouldReturnNumOfZeroClicks_CountMultipleTimes(t *testing.T) {
	direction := "R"
	numOfMoves := 1000
	currentPosition := 50
	minNum := 0
	maxNum := 99

	_, zeroClicksCount := MakeMove(direction, numOfMoves, currentPosition, minNum, maxNum)

	assert.Equal(t, 10, zeroClicksCount, "Incorrect zero count")
}

func TestPart1_ShouldRunInstructions_SingleInstruction_NoZerosHit(t *testing.T) {
	input := "L68"
	minNum, maxNum, startingPosition := 0, 99, 50

	output := Part1(input, minNum, maxNum, startingPosition)

	assert.Equal(t, 0, output, "Incorrect password output")
}

func TestPart1_ShouldRunInstructions_MultipleInstructions_ZeroHitOnce(t *testing.T) {
	input := `L68
L30
R48`
	minNum, maxNum, startingPosition := 0, 99, 50

	output := Part1(input, minNum, maxNum, startingPosition)

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
	minNum, maxNum, startingPosition := 0, 99, 50

	output := Part1(input, minNum, maxNum, startingPosition)

	assert.Equal(t, 3, output, "Incorrect password output")
}

func TestPart2_ShouldRunInstructions_SingleInstruction_NoZerosHit(t *testing.T) {
	input := "L1"
	minNum, maxNum, startingPosition := 0, 99, 50

	output := Part2(input, minNum, maxNum, startingPosition)

	assert.Equal(t, 0, output, "Incorrect zero hit count")
}

func TestPart2_ShouldRunInstructions_SingleInstruction_OneZeroHit(t *testing.T) {
	input := "L68"
	minNum, maxNum, startingPosition := 0, 99, 50

	output := Part2(input, minNum, maxNum, startingPosition)

	assert.Equal(t, 1, output, "Incorrect zero hit count")
}

func TestPart2_ShouldRunInstructions_MultipleInstructions_MultipleZerosHit(t *testing.T) {
	input := `L68
L30
R48`
	minNum, maxNum, startingPosition := 0, 99, 50

	output := Part2(input, minNum, maxNum, startingPosition)

	assert.Equal(t, 2, output, "Incorrect zero hit count")
}

func TestPart2_ShouldRunInstructions_WorksWithPuzzleTestFullInput(t *testing.T) {
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
	minNum, maxNum, startingPosition := 0, 99, 50

	output := Part2(input, minNum, maxNum, startingPosition)

	assert.Equal(t, 6, output, "Incorrect zero hit count")
}