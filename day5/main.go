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

	fmt.Printf("Day 5, Part 1 Output: %d\n", p1Output)

	// Part 2
	//p2Output := Part2(rawInput)

	//fmt.Printf("Day 5, Part 2 Output: %d\n", p2Output)
}

func Part1(input string) int {
	total := 0

	idRanges, availableIDs := GetIngredientsData(input)

	total = GetTotalFreshAvailableIngredients(idRanges, availableIDs)

	return total
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

func GetIngredientsData(input string) ([]string, []int) {
	rangeIDs := []string{}
	availableIDs := []int{}
	
    parts := strings.Split(input, "\n\n")

    if len(parts) != 2 {
        return rangeIDs, availableIDs
    }

    // First part: ranges ("3-5")
    for line := range strings.SplitSeq(parts[0], "\n") {
        line = strings.TrimSpace(line)

        if line != "" {
            rangeIDs = append(rangeIDs, line)
        }
    }

    // Second part: single numbers (convert to actual int's)
    for line := range strings.SplitSeq(parts[1], "\n") {
        line = strings.TrimSpace(line)

        if line != "" {
            if num, err := strconv.Atoi(line); err == nil {
                availableIDs = append(availableIDs, num)
            }
        }
    }

	return rangeIDs, availableIDs
}

func IsNumberInRange(numToCheck int, rangeToCheck string) bool {
	isNumInRange := false

	// seperate our number range into 2 numbers
	numRangeStr := strings.Split(rangeToCheck, "-")

    firstNum, err := strconv.Atoi(numRangeStr[0])

    if err != nil {
        fmt.Printf("Invalid number: %s\n", numRangeStr[0])
    }

    secondNum, err := strconv.Atoi(numRangeStr[1])

    if err != nil {
        fmt.Printf("Invalid number: %s\n", numRangeStr[1])
    }

	if numToCheck >= firstNum && numToCheck <= secondNum {
		isNumInRange = true
	}

	return isNumInRange
}

func GetTotalFreshAvailableIngredients(idRanges []string, availableIDs []int) int {
	total := 0

	for i := 0; i < len(availableIDs); i++ {
		for j := 0; j < len(idRanges); j++ {
			isNumInRange := IsNumberInRange(availableIDs[i], idRanges[j])

			if isNumInRange {
				total++

				break
			}
		}
	}

	return total
}