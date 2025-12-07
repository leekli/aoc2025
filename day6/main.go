package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Data setup
	rawInput, err := readFileToString("input.txt")

	if err != nil {
		os.Exit(-1)
	}

	// Part 1
	p1Output := Part1(rawInput)

	fmt.Printf("Day 6, Part 1 Output: %d\n", p1Output)

	// Part 2
	p2Output := Part2(rawInput)

	fmt.Printf("Day 6, Part 2 Output: %d\n", p2Output)
}

func Part1(input string) int {
	total := 0

	operationsList := ConvertInputToOperations(input)

	for i := 0; i < len(operationsList); i++ {
		sum := GetTotalForCurrentOperation(operationsList[i])

		total += sum
	}

	return total
}

func Part2(input string) int {
	total := 0

	operationsList := ConvertInputRightToLeftToOperations(input)

	for i := 0; i < len(operationsList); i++ {
		sum := GetTotalForCurrentOperation(operationsList[i])

		total += sum
	}

	return total
}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}
	
	return string(content), nil
}

type Operation struct {
	Numbers []int
	Op 		string
}

func ConvertInputToOperations(input string) []Operation { 
	operationsList := []Operation{}

	lines := strings.Split(input, "\n")

	// Get rid of all whitespace apart from single spaces (replace), get rid of remaining space if any
	for i := 0; i < len(lines); i++ {
		space := regexp.MustCompile(`\s+`)

		lines[i] = space.ReplaceAllString(lines[i], " ")

		lines[i] = strings.TrimSpace(lines[i])
	}

	// Turn the `lines` above into seperate 2d array, each 2d array contains all numbers in the same column along with operation (* or +)
	allOperations := [][]string{}

	for i := 0; i < len(lines); i++ {
		currentOperation := []string{}

		splitLine := strings.Split(lines[i], " ")

		currentOperation = append(currentOperation, splitLine...)

		allOperations = append(allOperations, currentOperation)
	}

	// Now go through the allOperations var & start sorting it into columns for the Operations struct
	maxLength := len(allOperations[0])
	colNum := 0
	
	for colNum < maxLength {
		operation := Operation{}

		for i := 0; i < len(allOperations); i++ {
			currentElem := allOperations[i][colNum]

			if (currentElem != "*" && currentElem != "+") {
				convertedNum, _ := strconv.Atoi(currentElem)

				// add to struct .Numbers field
				operation.Numbers = append(operation.Numbers, convertedNum)

			} else {
				// add to struct .Op field
				operation.Op = currentElem
			}
		}

		// Add current built operation to wider ops list
		operationsList = append(operationsList, operation)

		colNum++
	}

	return operationsList
}

func GetTotalForCurrentOperation(operation Operation) int {
    if len(operation.Numbers) == 0 {
        return 0
    }

    switch operation.Op {
    case "+":
        total := 0
        for _, num := range operation.Numbers {
            total += num
        }
        return total
    case "*":
        total := 1
        for _, num := range operation.Numbers {
            total *= num
        }
        return total
    default:
        return 0 
    }
}

func ConvertInputRightToLeftToOperations(input string) []Operation {
    operationsList := []Operation{}

	lines := strings.Split(input, "\n")

	// Turn the `lines` above into seperate 2d array, each 2d array contains all numbers, no whitespace trimmed this time to try align each number up
	allOperations := [][]string{}

	for i := 0; i < len(lines); i++ {
		currentOperation := []string{}

		splitLine := strings.Split(lines[i], "")

		currentOperation = append(currentOperation, splitLine...)

		allOperations = append(allOperations, currentOperation)
	}

	maxCols := len(allOperations[0])
	numRows := len(allOperations)
	col := 0

	for col < maxCols {
		// Skip separator columns (all spaces)
		isSeparator := true

		for row := 0; row < numRows; row++ {
			if allOperations[row][col] != "" && allOperations[row][col] != " " {
				isSeparator = false
				break
			}
		}

		if isSeparator {
			col++

			continue
		}

		// Build up groups of non-separator columns
		groupCols := []int{}
		for col < maxCols {
			isSeparator := true
			for row := 0; row < numRows; row++ {
				if allOperations[row][col] != "" && allOperations[row][col] != " " {
					isSeparator = false

					break
				}
			}

			if isSeparator {
				break
			}

			groupCols = append(groupCols, col)
			col++
		}

		// Get the numbers and add to struct .Numbers field
		op := Operation{}

		for _, col := range groupCols {
			numStr := ""

			for row := 0; row < numRows - 1; row++ { // skip operator
				currentElem := allOperations[row][col]

				if currentElem != "" && currentElem != " " {
					numStr += currentElem
				}
			}

			if numStr != "" {
				// Convert to actual number/int
				num, _ := strconv.Atoi(numStr)

				op.Numbers = append(op.Numbers, num)

			}
		}

		// Get the Operator (at the bottom row of the last column in the group) and add to .Op field
		op.Op = allOperations[numRows - 1][groupCols[0]]

		operationsList = append(operationsList, op)
	}

    return operationsList
}