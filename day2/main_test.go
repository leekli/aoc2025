package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertRangesToList_ReturnsEmptySliceForEmptyInstructions(t *testing.T) {
	input := ""

	output := ConvertRangesToList(input)
	
	outputLength := len(output)

	assert.Equal(t, 1, outputLength, "Output length should be 1")
}

func TestConvertRangesToList_ReturnsSliceForSingleInstruction(t *testing.T) {
	input := "11-22"

	output := ConvertRangesToList(input)
	
	outputLength := len(output)

	assert.Equal(t, 1, outputLength, "Output length should be 1")
}

func TestConvertRangesToList_ReturnsSliceForMultipleInstruction(t *testing.T) {
	input := "11-22,95-115"

	output := ConvertRangesToList(input)
	
	outputLength := len(output)

	assert.Equal(t, 2, outputLength, "Output length should be 2")

	input = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`

	output = ConvertRangesToList(input)
	
	outputLength = len(output)

	assert.Equal(t, 11, outputLength, "Output length should be 11")
}

func TestIsNumberAnInvalidID_ReturnsFalseForNotAnInvalidID(t *testing.T) {
	input := 1

	output := IsNumberAnInvalidID(input)

	assert.Equal(t, false, output, "Result should be false")

	input = 123

	output = IsNumberAnInvalidID(input)

	assert.Equal(t, false, output, "Result should be false")

	input = 1012

	output = IsNumberAnInvalidID(input)

	assert.Equal(t, false, output, "Result should be false")

	input = 222220

	output = IsNumberAnInvalidID(input)

	assert.Equal(t, false, output, "Result should be false")
}

func TestIsNumberAnInvalidID_ReturnsTrueForAnInvalidID(t *testing.T) {
	input := 11

	output := IsNumberAnInvalidID(input)

	assert.Equal(t, true, output, "Result should be true")

	input = 22

	output = IsNumberAnInvalidID(input)

	assert.Equal(t, true, output, "Result should be true")

	input = 1010

	output = IsNumberAnInvalidID(input)

	assert.Equal(t, true, output, "Result should be true")

	input = 1188511885

	output = IsNumberAnInvalidID(input)

	assert.Equal(t, true, output, "Result should be true")

	input = 222222

	output = IsNumberAnInvalidID(input)

	assert.Equal(t, true, output, "Result should be true")

	input = 446446

	output = IsNumberAnInvalidID(input)

	assert.Equal(t, true, output, "Result should be true")

	input = 38593859

	output = IsNumberAnInvalidID(input)

	assert.Equal(t, true, output, "Result should be true")
}

func TestIsNumberAnInvalidID_ReturnsFalseIfInvalidIDHasLeadingZero(t *testing.T) {
	input := 0101

	output := IsNumberAnInvalidID(input)

	assert.Equal(t, false, output, "Result should be false")
}

func TestCheckRangeForInValidIDs_ReturnsEmptySliceForNoInvalidIDs(t *testing.T) {
	input := "1-2"

	output := CheckRangeForInValidIDs(input, 1)

	assert.Equal(t, 0, len(output), "Output should be 0")

	input = "12-21"

	output = CheckRangeForInValidIDs(input, 1)

	assert.Equal(t, 0, len(output), "Output should be 0")

	input = "1698522-1698528"

	output = CheckRangeForInValidIDs(input, 1)

	assert.Equal(t, 0, len(output), "Output should be 0")

	input = "2121212118-2121212124"

	output = CheckRangeForInValidIDs(input, 1)

	assert.Equal(t, 0, len(output), "Output should be 0")
}

func TestCheckRangeForInValidIDs_ReturnsSliceForNoInvalidIDs(t *testing.T) {
	input := "11-22"

	output := CheckRangeForInValidIDs(input, 1)

	assert.Equal(t, 2, len(output), "Output should be 2")

	input = "95-115"

	output = CheckRangeForInValidIDs(input, 1)

	assert.Equal(t, 1, len(output), "Output should be 1")

	input = "998-1012"

	output = CheckRangeForInValidIDs(input, 1)

	assert.Equal(t, 1, len(output), "Output should be 1")

	input = "1188511880-1188511890"

	output = CheckRangeForInValidIDs(input, 1)

	assert.Equal(t, 1, len(output), "Output should be 1")	

	input = "222220-222224"

	output = CheckRangeForInValidIDs(input, 1)

	assert.Equal(t, 1, len(output), "Output should be 1")	

	input = "446443-446449"

	output = CheckRangeForInValidIDs(input, 1)

	assert.Equal(t, 1, len(output), "Output should be 1")	

	input = "38593856-38593862"

	output = CheckRangeForInValidIDs(input, 1)

	assert.Equal(t, 1, len(output), "Output should be 1")	
}

func TestPart1_ReturnsTotalForAllInvalidIDsFound(t *testing.T) {
	input := "11-22"

	output := Part1(input)

	assert.Equal(t, 33, output, "Result should be 33")

	input = "11-22,95-115"

	output = Part1(input)

	assert.Equal(t, 132, output, "Result should be 132")
}

func TestPart1_ReturnsTotalForFullTestInput(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

	output := Part1(input)

	assert.Equal(t, 1227775554, output, "Result should be 1227775554")
}

func TestIsRepeatedAtLeastTwice_ReturnsFalseForNotAnInvalidID(t *testing.T) {
	input := 1

	output := IsRepeatedAtLeastTwice(input)

	assert.Equal(t, false, output, "Result should be false")

	input = 123

	output = IsRepeatedAtLeastTwice(input)

	assert.Equal(t, false, output, "Result should be false")

	input = 1698522
}

func TestIsRepeatedAtLeastTwice_ReturnsTrueForAnInvalidID(t *testing.T) {
	input := 999

	output := IsRepeatedAtLeastTwice(input)

	assert.Equal(t, true, output, "Result should be true")

	input = 1010

	output = IsRepeatedAtLeastTwice(input)

	assert.Equal(t, true, output, "Result should be true")

	input = 12341234

	output = IsRepeatedAtLeastTwice(input)

	assert.Equal(t, true, output, "Result should be true")

	input = 123123123

	output = IsRepeatedAtLeastTwice(input)

	assert.Equal(t, true, output, "Result should be true")

	input = 565656

	output = IsRepeatedAtLeastTwice(input)

	assert.Equal(t, true, output, "Result should be true")
}

func TestPart2_ReturnsTotalForFullTestInput(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

	output := Part2(input)

	assert.Equal(t, 4174379265, output, "Result should be 4174379265")
}