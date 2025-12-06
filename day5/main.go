package main

import (
	"fmt"
	"os"
	"sort"
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
	p2Output := Part2(rawInput)

	fmt.Printf("Day 5, Part 2 Output: %d\n", p2Output)
}

func Part1(input string) int {
	total := 0

	idRanges, availableIDs := GetIngredientsData(input)

	total = GetTotalFreshAvailableIngredients(idRanges, availableIDs)

	return total
}

func Part2(input string) int {
	total := 0

	idRanges, _ := GetIngredientsData(input)

	total = GetTotalUniqueIDsInRanges(idRanges)

	return total
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

func GetTotalUniqueIDsInRanges(idRanges []string) int {
    type interval struct { 
		start, end int 
	}

    intervals := make([]interval, 0, len(idRanges))

    // Parse ranges into the interval struct
    for _, idRange := range idRanges {
        parts := strings.Split(idRange, "-")

        start, _ := strconv.Atoi(parts[0])
        end, _ := strconv.Atoi(parts[1])

        intervals = append(intervals, interval{ start, end })
    }

    // Sort by each internal start value
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })

    // Merge the intervals
    merged := []interval{}
    for _, interval := range intervals {
        num := len(merged)

        if num == 0 || interval.start > merged[num - 1].end + 1 {
            merged = append(merged, interval)
        } else {
            if interval.end > merged[num - 1].end {
                merged[num - 1].end = interval.end
            }
        }
    }

    // Sum the total difference between each merge range
    total := 0

    for _, currentMerge := range merged {
        total += currentMerge.end - currentMerge.start + 1
    }

    return total
}