package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertRangesToList_ReturnsEmptySliceForEmptyInput(t *testing.T) {
	input := ""

	output := ConvertRangesToList(input)
	
	outputLength := len(output)

	assert.Equal(t, 1, outputLength, "Output length should be 1")
}

func TestConvertRangesToList_ReturnsSliceForSingleInput(t *testing.T) {
	input := "987654321111111"

	output := ConvertRangesToList(input)
	
	outputLength := len(output)

	assert.Equal(t, 1, outputLength, "Output length should be 1")
}

func TestConvertRangesToList_ReturnsSliceForMultipleInputs(t *testing.T) {
	input := `987654321111111
811111111111119`

	output := ConvertRangesToList(input)
	
	outputLength := len(output)

	assert.Equal(t, 2, outputLength, "Output length should be 2")

	input = `987654321111111
811111111111119
234234234234278
818181911112111`

	output = ConvertRangesToList(input)
	
	outputLength = len(output)

	assert.Equal(t, 4, outputLength, "Output length should be 4")
}

func TestConvertBankSliceToInts_ReturnsEmptySliceForEmptyInput(t *testing.T) {
	input := ""

	output := ConvertBankSliceToInts(input)
	
	outputLength := len(output)

	assert.Equal(t, 0, outputLength, "Output length should be 0")
}

func TestConvertBankSliceToInts_ReturnsSliceForSingleInput(t *testing.T) {
	input := "1"

	output := ConvertBankSliceToInts(input)
	outputLength := len(output)

	assert.Equal(t, 1, output[0], "Output[0] should be 1")
	assert.Equal(t, 1, outputLength, "Output length should be 1")

	input = "0"

	output = ConvertBankSliceToInts(input)
	outputLength = len(output)

	assert.Equal(t, 0, output[0], "Output[0] should be 0")
	assert.Equal(t, 1, outputLength, "Output length should be 1")
}

func TestConvertBankSliceToInts_ReturnsSliceForMultipleInputs(t *testing.T) {
	input := "123"

	output := ConvertBankSliceToInts(input)
	outputLength := len(output)

	assert.Equal(t, 1, output[0], "Output[0] should be 1")
	assert.Equal(t, 2, output[1], "Output[1] should be 2")
	assert.Equal(t, 3, output[2], "Output[2] should be 3")
	assert.Equal(t, 3, outputLength, "Output length should be 3")

	input = "987654321111111"

	output = ConvertBankSliceToInts(input)
	outputLength = len(output)

	assert.Equal(t, 9, output[0], "Output[0] should be 9")
	assert.Equal(t, 5, output[4], "Output[4] should be 5")
	assert.Equal(t, 1, output[14], "Output[14] should be 1")
	assert.Equal(t, 15, outputLength, "Output length should be 15")
}

func TestConvertBankSliceToStrs_ReturnsEmptySliceForEmptyInput(t *testing.T) {
	input := []int{}

	output := ConvertBankSliceToStrs(input)
	
	outputLength := len(output)

	assert.Equal(t, 0, outputLength, "Output length should be 0")
}

func TestConvertBankSliceToStrs_ReturnsSliceForSingleInput(t *testing.T) {
	input := []int{ 1 }

	output := ConvertBankSliceToStrs(input)
	outputLength := len(output)

	assert.Equal(t, "1", output[0], "Output[0] should be '1'")
	assert.Equal(t, 1, outputLength, "Output length should be 1")

	input = []int{ 0 }

	output = ConvertBankSliceToStrs(input)
	outputLength = len(output)

	assert.Equal(t, "0", output[0], "Output[0] should be '0'")
	assert.Equal(t, 1, outputLength, "Output length should be 1")
}

func TestConvertBankSliceToStrs_ReturnsSliceForMultipleInputs(t *testing.T) {
	input := []int{ 1, 2, 3 }

	output := ConvertBankSliceToStrs(input)
	outputLength := len(output)

	assert.Equal(t, "1", output[0], "Output[0] should be '1'")
	assert.Equal(t, "2", output[1], "Output[1] should be '2'")
	assert.Equal(t, "3", output[2], "Output[2] should be '3'")
	assert.Equal(t, 3, outputLength, "Output length should be 3")

	input = []int{ 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1 }

	output = ConvertBankSliceToStrs(input)
	outputLength = len(output)

	assert.Equal(t, "9", output[0], "Output[0] should be '9'")
	assert.Equal(t, "5", output[4], "Output[4] should be '5'")
	assert.Equal(t, "1", output[14], "Output[14] should be '1'")
	assert.Equal(t, 15, outputLength, "Output length should be 15")
}

func TestFindHighestNumber_ReturnsNegativeValuesForEmptySlice(t *testing.T) {
	input := []int{}

	outputNum, outputIndex := FindHighestNumber(input)

	assert.Equal(t, -1, outputNum, "Output should be -1")
	assert.Equal(t, -1, outputIndex, "Output should be -1")
}

func TestFindHighestNumber_ReturnsHighestNumAndIndexForVariousIntSlices(t *testing.T) {
	input := []int{ 1 }

	outputNum, outputIndex := FindHighestNumber(input)

	assert.Equal(t, 1, outputNum, "Output should be 1")
	assert.Equal(t, 0, outputIndex, "Output should be 0")

	input = []int{ 1, 0 }

	outputNum, outputIndex = FindHighestNumber(input)

	assert.Equal(t, 1, outputNum, "Output should be 1")
	assert.Equal(t, 0, outputIndex, "Output should be 0")

	input = []int{ 4, 6, 8, 1, 2 }

	outputNum, outputIndex = FindHighestNumber(input)

	assert.Equal(t, 8, outputNum, "Output should be 8")
	assert.Equal(t, 2, outputIndex, "Output should be 2")

	input = []int{ 9,8,7,6,5,4,3,2,1,1,1,1,1,1,1 }

	outputNum, outputIndex = FindHighestNumber(input)

	assert.Equal(t, 9, outputNum, "Output should be 9")
	assert.Equal(t, 0, outputIndex, "Output should be 0")

	input = []int{ 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9 }

	outputNum, outputIndex = FindHighestNumber(input)

	assert.Equal(t, 9, outputNum, "Output should be 9")
	assert.Equal(t, 14, outputIndex, "Output should be 14")

	input = []int{ 2,3,4,2,3,4,2,3,4,2,3,4,2,7,8 }

	outputNum, outputIndex = FindHighestNumber(input)

	assert.Equal(t, 8, outputNum, "Output should be 8")
	assert.Equal(t, 14, outputIndex, "Output should be 14")
}

func TestFindSecondBattery_ReturnsSecondHighestNum(t *testing.T) {
	input := "98"

	num := FindSecondBattery(input, 9, 0)

	assert.Equal(t, 98, num, "Output should be 98")

	input = "987654321111111"

	num = FindSecondBattery(input, 9, 0)

	assert.Equal(t, 98, num, "Output should be 98")

	input = "811111111111119"

	num = FindSecondBattery(input, 9, 14)

	assert.Equal(t, 89, num, "Output should be 89")

	input = "234234234234278"

	num = FindSecondBattery(input, 8, 14)

	assert.Equal(t, 78, num, "Output should be 78")

	input = "818181911112111"

	num = FindSecondBattery(input, 9, 6)

	assert.Equal(t, 92, num, "Output should be 92")
}

func TestPart1_TestInput(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`

	output := Part1(input)

	assert.Equal(t, 357, output, "Output should be 357")
}