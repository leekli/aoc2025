package main

import (
	"fmt"
	"os"
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

	fmt.Printf("Day 4, Part 1 Output: %d\n", p1Output)

	// Part 2
	//p2Output := Part2(rawInput)

	//fmt.Printf("Day 4, Part 2 Output: %d\n", p2Output)
}

func Part1(input string) int {
	totalRollsAccessible := 0

	grid := StringTo2DArray(input)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			currentGridLocation := grid[i][j]

			if currentGridLocation != "@" {
            	continue
        	}

			if currentGridLocation == "@" {
				totalAdjacentRolls := 0

				// Look up
				if i - 1 >= 0 {
					itemFound := grid[i - 1][j]

					if IsPaperRollFound(itemFound) {
						totalAdjacentRolls++
					}
				}

				// Look down
				if i + 1 < len(grid) {
					itemFound := grid[i + 1][j]

					if IsPaperRollFound(itemFound) {
						totalAdjacentRolls++
					}
				}

				// Look left
				if j - 1 >= 0 {
					itemFound := grid[i][j - 1]

					if IsPaperRollFound(itemFound) {
						totalAdjacentRolls++
					}
				}

				// Look right
				if j + 1 < len(grid[i]) {
					itemFound := grid[i][j + 1]

					if IsPaperRollFound(itemFound) {
						totalAdjacentRolls++
					}
				}

				// Look diagonal top left
				if i - 1 >= 0 && j - 1 >= 0 {
					itemFound := grid[i - 1][j - 1]

					if IsPaperRollFound(itemFound) {
						totalAdjacentRolls++
					}
				}

				// Look diagonal top right
				if i - 1 >= 0 && j + 1 < len(grid[i - 1]) {
					itemFound := grid[i - 1][j + 1]

					if IsPaperRollFound(itemFound) {
						totalAdjacentRolls++
					}
				}

				// Look diagonal bottom left
				if i + 1 < len(grid) && j - 1 >= 0 {
					itemFound := grid[i + 1][j - 1]

					if IsPaperRollFound(itemFound) {
						totalAdjacentRolls++
					}
				}

				// Look diagonal bottom right
				if i + 1 < len(grid) && j + 1 < len(grid[i + 1]) {
					itemFound := grid[i + 1][j + 1]

					if IsPaperRollFound(itemFound) {
						totalAdjacentRolls++
					}
				}

				if currentGridLocation == "@" && totalAdjacentRolls < 4 {
					totalRollsAccessible++
				}
			}
		}
	}

	return totalRollsAccessible
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

func StringTo2DArray(input string) [][]string {
    lineRows := strings.Split(input, "\n")

    var wordSearchArray [][]string

    for _, lineRow := range lineRows {
        splitChars := strings.Split(lineRow, "")

        var finalChars []string

        for _, numStr := range splitChars {

            finalChars = append(finalChars, numStr)
        }

        wordSearchArray = append(wordSearchArray, finalChars)
    }

    return wordSearchArray
}


func IsPaperRollFound(item string) bool {
	rollFound := false

	if item == "@" {
		rollFound = true
	}

	return rollFound
}