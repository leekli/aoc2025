package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertInputToOperations_ProducesCorrectFormat(test *testing.T) {
	input := `123 
 45 
  6
*`

	output := ConvertInputToOperations(input)

	assert.Equal(test, 1, len(output), "Expected: 1, Received: %s", len(output))
	assert.Equal(test, 123, output[0].Numbers[0], "Expected: 123, Received: %s", output[0].Numbers[0])
	assert.Equal(test, 45, output[0].Numbers[1], "Expected: 45, Received: %s", output[0].Numbers[1])
	assert.Equal(test, 6, output[0].Numbers[2], "Expected: 6, Received: %s", output[0].Numbers[2])
	assert.Equal(test, "*", output[0].Op, "Expected: '*', Received: %s", output[0].Op)

	input = `123 328 
 45 64 
  6 98
*   + `

	output = ConvertInputToOperations(input)

	assert.Equal(test, 2, len(output), "Expected: 2, Received: %s", len(output))
	assert.Equal(test, 123, output[0].Numbers[0], "Expected: 123, Received: %s", output[0].Numbers[0])
	assert.Equal(test, 328, output[1].Numbers[0], "Expected: 328, Received: %s", output[1].Numbers[0])
	assert.Equal(test, "+", output[1].Op, "Expected: '+', Received: %s", output[1].Op)

	input = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	output = ConvertInputToOperations(input)

	assert.Equal(test, 4, len(output), "Expected: 4, Received: %s", len(output))
}

func TestGetTotalForCurrentOperation_ReturnsCorrectTotal(test *testing.T) {
	input := Operation{
		Numbers: []int{ 1, 2, 3 },
		Op: "+",
	}

	output := GetTotalForCurrentOperation(input)

	assert.Equal(test, 6, output, "Expected: 6, Received: %s", output)

	input = Operation{
		Numbers: []int{ 123, 45, 6 },
		Op: "*",
	}

	output = GetTotalForCurrentOperation(input)

	assert.Equal(test, 33210, output, "Expected: 33210, Received: %s", output)

	input = Operation{
		Numbers: []int{ 328, 64, 98 },
		Op: "+",
	}

	output = GetTotalForCurrentOperation(input)

	assert.Equal(test, 490, output, "Expected: 490, Received: %s", output)

	input = Operation{
		Numbers: []int{ 51, 387, 215 },
		Op: "*",
	}

	output = GetTotalForCurrentOperation(input)

	assert.Equal(test, 4243455, output, "Expected: 4243455, Received: %s", output)

	input = Operation{
		Numbers: []int{ 64, 23, 314 },
		Op: "+",
	}

	output = GetTotalForCurrentOperation(input)

	assert.Equal(test, 401, output, "Expected: 401, Received: %s", output)
}

func TestPart1_ReturnsCorrectTotal(test *testing.T) {
		input := `123 
 45 
  6
*`

	output := Part1(input)

	assert.Equal(test, 33210, output, "Expected: 33210, Received: %s", output)

	input = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	output = Part1(input)

	assert.Equal(test, 4277556, output, "Expected: 4277556, Received: %s", output)
}

func TestConvertInputRightToLeftToOperations_ProducesCorrectFormat(test *testing.T) {
	input := `123
 45
  6
*  `

	output := ConvertInputRightToLeftToOperations(input)

	assert.Equal(test, 1, len(output), "Expected: 1, Received: %s", len(output))

	input = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	output = ConvertInputRightToLeftToOperations(input)

	assert.Equal(test, 4, len(output), "Expected: 4, Received: %s", len(output))
}

func TestPart2_ReturnsCorrectTotal(test *testing.T) {
	input := `123
 45
  6
*  `

	output := Part2(input)

	assert.Equal(test, 8544, output, "Expected: 8544, Received: %s", output)

	input = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	output = Part2(input)

	assert.Equal(test, 3263827, output, "Expected: 3263827, Received: %s", output)
}