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

	fmt.Printf("Day 1, Part 1 Output: %d\n", p1Output)

	// Part 2
	//p2Output := Part2(rawInput)

	//fmt.Printf("Day 1, Part 2 Output: %d\n", p2Output)
}

func Part1(input string) int {
	dialInstructions := ConvertInstructionsToList(input)

	finalPassword := 0

	minNum, maxNum := 0, 99

	// Staring Position for the challenge is 50
	position := 50

	for i := 0; i < len(dialInstructions); i++ {
		instruction := dialInstructions[i]

		// On current instruction, seperate out the L or R and leave just the number
		if len(instruction) < 2 {
            continue // skip invalid lines
        }

		directionToGo := instruction[:1]

		movesStr := instruction[1:]
        movesToMake, err := strconv.Atoi(movesStr)

        if err != nil {
            fmt.Printf("Invalid number in instruction: %s\n", instruction)
            continue
        }

		// Submit direction and number to MakeMove()
		newLocation := MakeMove(directionToGo, movesToMake, position, minNum, maxNum)

		// If returned location is 0, increment 'finalPassword' +1
		if newLocation == 0 {
			finalPassword++
		}

		// Update position to where latest position now is
		position = newLocation
	}

	return finalPassword
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

func ConvertInstructionsToList(input string) []string {
	return strings.Split(input, "\n")
}

func MakeMove(direction string, numOfMoves int, currentPosition int, minNum int, maxNum int) int {
	location := currentPosition

	for i := 1; i <= numOfMoves; i++ {
		if direction == "L" {
			location -= 1

			if location < minNum {
				location = maxNum
			}
		}

		if direction == "R" {
			location += 1

			if location > maxNum {
				location = minNum
			}
		}
	}

	return location
}