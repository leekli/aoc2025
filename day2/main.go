package main

import (
	"fmt"
	"os"
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

	fmt.Printf("Day 2, Part 1 Output: %d\n", p1Output)

	// Part 2
	//p2Output := Part2(rawInput)

	//fmt.Printf("Day 2, Part 2 Output: %d\n", p2Output)
}

func Part1(input string) int {
	ranges := ConvertRangesToList(input)

	totalSumOfInvalidIDs := 0

	for i := 0; i < len(ranges); i++ {
		invalidIDsFound := CheckRangeForInValidIDs(ranges[i])

		if (len(invalidIDsFound) > 0) {
			currentSum := AddInvalidIDs(invalidIDsFound)

			totalSumOfInvalidIDs += currentSum
		}
	}

	return totalSumOfInvalidIDs
}

func Part2(input string) int {
	return 0
}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}
	
	return string(content), nil
}

func ConvertRangesToList(input string) []string {
	return strings.Split(input, ",")
}

func CheckRangeForInValidIDs(numRange string) []int {
	// Split string by - character to get starting and ending number(s) of the given range. Convert to Go numbers
	numRangeSplit := strings.Split(numRange, "-")
	
    startingNum, err := strconv.Atoi(numRangeSplit[0])

    if err != nil {
        fmt.Printf("Invalid number: %s\n", numRangeSplit[0])
    }

	endingNum, err := strconv.Atoi(numRangeSplit[1])

    if err != nil {
        fmt.Printf("Invalid number: %s\n", numRangeSplit[1])
    }

	invalidIDList := []int{}

	for i := startingNum; i <= endingNum; i++ {
		isInvalidID := IsNumberAnInvalidID(i)

		if (isInvalidID) {
			invalidIDList = append(invalidIDList, i)
		}
	}

	return invalidIDList
}

func IsNumberAnInvalidID(numToCheck int) bool {
	numString := strconv.Itoa(numToCheck)

	numLength := len(numString)

    if numLength % 2 != 0 {
        return false
    }

    strHalf := numLength / 2

	isIdentical := numString[:strHalf] == numString[strHalf:]

	if isIdentical {
		firstNum := numString[:1]
		if firstNum == "0" {
			isIdentical = false
		}
	}

    return isIdentical
}

func AddInvalidIDs(nums []int) int {
	total := 0

	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	return total 
}