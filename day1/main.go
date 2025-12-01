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

	// Vars for both parts
	minNum, maxNum := 0, 99
	startingPosition := 50

	// Part 1
	p1Output := Part1(rawInput, minNum, maxNum, startingPosition)

	fmt.Printf("Day 1, Part 1 Output: %d\n", p1Output)

	// Part 2
	p2Output := Part2(rawInput, minNum, maxNum, startingPosition)

	fmt.Printf("Day 1, Part 2 Output: %d\n", p2Output)
}

func Part1(input string, minNum int, maxNum int, startingPosition int) int {
	dialInstructions := ConvertInstructionsToList(input)

	finalPassword := 0

	// Staring Position for the challenge is 50
	position := startingPosition

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
		newLocation, _ := MakeMove(directionToGo, movesToMake, position, minNum, maxNum)

		// If returned location is 0, increment 'finalPassword' +1
		if newLocation == 0 {
			finalPassword++
		}

		// Update position to where latest position now is
		position = newLocation
	}

	return finalPassword
}

func Part2(input string, minNum int, maxNum int, startingPosition int) int {
	dialInstructions := ConvertInstructionsToList(input)

	totalZeroClicks := 0

	// Staring Position for the challenge is 50
	position := startingPosition

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
		newLocation, numOfZeroClicksThisInstruction := MakeMove(directionToGo, movesToMake, position, minNum, maxNum)

		// Add number of 'zero clicks' hit in this iteration to current running total
		totalZeroClicks += numOfZeroClicksThisInstruction

		// Update position to where latest position now is
		position = newLocation
	}

	return totalZeroClicks
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

func MakeMove(direction string, numOfMoves int, currentPosition int, minNum int, maxNum int) (int, int) {
	location := currentPosition
	totalNumOfZeroClicks := 0

	for i := 1; i <= numOfMoves; i++ {
		if direction == "L" {
			location -= 1

			if location < minNum {
				location = maxNum
			}

			if location == 0 {
				totalNumOfZeroClicks++
			}
		}

		if direction == "R" {
			location += 1

			if location > maxNum {
				location = minNum
			}

			if location == 0 {
				totalNumOfZeroClicks++
			}
		}
	}

	return location, totalNumOfZeroClicks
}