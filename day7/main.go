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

	fmt.Printf("Day 7, Part 1 Output: %d\n", p1Output)

	// Part 2
	//p2Output := Part2(rawInput)

	//fmt.Printf("Day 7, Part 2 Output: %d\n", p2Output)
}

func Part1(input string) int {
	total := 0

	gridArray := StringTo2DArray(input)

	splitCount := CountTotalSplits(gridArray)

	total += splitCount

	return total
}

func Part2(input string) int {
	total := 0

	return total
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

    var grid [][]string

    for _, lineRow := range lineRows {
        splitChars := strings.Split(lineRow, "")

        var finalChars []string

        for _, numStr := range splitChars {

            finalChars = append(finalChars, numStr)
        }

        grid = append(grid, finalChars)
    }

    return grid
}

func FindStartingBeam(grid [][]string) int {
	startingCol := -1

	for col := 0; col < len(grid[0]); col++ {
		if grid[0][col] == "S" {
			startingCol = col

			break
		}
	}

	return startingCol
}

func CountTotalSplits(grid [][]string) int {
	type Beam struct {
		row int
		col int
	}

	gridHeight := len(grid)
	gridWidth := len(grid[0])

	totalSplits := 0

	// Find where the tacyon beam enters the manifold ('S' on grid)
	startingBeamColPos := FindStartingBeam(grid)

	// Track visited positions for each beam
    visited := make(map[string]bool)

	// Begin the queue
	queue := NewQueue()
	queue.Enqueue(Beam{ row: 0, col: startingBeamColPos })

	for queue.Len() > 0 {
		item := queue.Dequeue()

		// Keep moving donwards until we either leave the grid or hit a splitter (need assert type first), checks bound of array height
		if beam, ok := item.(Beam); ok {
			for beam.row + 1 < gridHeight {
				beam.row++

               // Create a unique key for this beam position to track
                key := fmt.Sprintf("%d,%d", beam.row, beam.col)

				// Already visited this position, skip further traversal
                if visited[key] {
                    break
                }
				
				// Log visit to this location
                visited[key] = true

				// Found a splitter
				if grid[beam.row][beam.col] == "^" {
					totalSplits++

					// Sort the left beam (check it's within bounds)
					if beam.col - 1 >= 0 {
                        leftKey := fmt.Sprintf("%d,%d", beam.row, beam.col - 1)

                        if !visited[leftKey] {
                            queue.Enqueue(Beam{row: beam.row, col: beam.col - 1})
                        }
					}

					// Sort the right beam (check it's within bounds)
					if beam.col + 1 < gridWidth {
                        rightKey := fmt.Sprintf("%d,%d", beam.row, beam.col + 1)

                        if !visited[rightKey] {
                            queue.Enqueue(Beam{row: beam.row, col: beam.col + 1})
                        }
					}

					// Stop the beam after splitting
					break
				}
			}	
		}
	}

	return totalSplits
}