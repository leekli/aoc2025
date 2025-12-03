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

	fmt.Printf("Day 3, Part 1 Output: %d\n", p1Output)

	// Part 2
	p2Output := Part2(rawInput)

	fmt.Printf("Day 3, Part 2 Output: %d\n", p2Output)
}

func Part1(input string) int {
	total := 0

	batteryBanks := ConvertRangesToList(input)

	for i := 0; i < len(batteryBanks); i++ {
		currentBankAsInt := ConvertBankSliceToInts(batteryBanks[i])

		highestNum, highestNumIndex := FindHighestNumber(currentBankAsInt)

		joltageFound := FindSecondBattery(batteryBanks[i], highestNum, highestNumIndex)

		total += joltageFound
	}

	return total
}

func Part2(input string) int {
	// Approach to Part 1 isn't going work here! - “remove k digits to form the largest possible number”, greedy stack based algo, could go back and refactor part 1 now but won't

	total := 0

	batteryBanks := ConvertRangesToList(input)

	for i := 0; i < len(batteryBanks); i++ {
		batteriesFound := FindTwelveJoltages(batteryBanks[i])

		batteriesFoundInt, err := strconv.Atoi(batteriesFound)

		if err != nil {
			panic(err)
		}

		total += batteriesFoundInt
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

func ConvertRangesToList(input string) []string {
	return strings.Split(input, "\n")
}

func ConvertBankSliceToInts(bank string) []int {
    ints := make([]int, len(bank))

    for i, ch := range bank {
        ints[i] = int(ch - '0')
    }

    return ints
}

func ConvertBankSliceToStrs(bank []int) []string {
	var bankStrs = []string{}

    for _, i := range bank {
        num := strconv.Itoa(i)

        bankStrs = append(bankStrs, num)
    }

    return bankStrs
}

func FindHighestNumber(bank []int) (int, int) {
	highestNum, indexOfNum := -1, -1

	if len(bank) <= 0 {
		return highestNum, indexOfNum
	}

	for i := 0; i < len(bank); i++ {
		if bank[i] > highestNum {
			highestNum = bank[i]
			indexOfNum = i
		}
	}

	return highestNum, indexOfNum
}

func FindSecondBattery(bank string, highestNum int, highestNumIndex int) int {
	secondHighestNum := -1

	highestNumStr := strconv.Itoa(highestNum)

	bankInts := ConvertBankSliceToInts(bank)

	for i := 0; i < len(bankInts); i++ {
		if i == highestNumIndex {
			continue
		} else {
			currentNumStr := strconv.Itoa(bankInts[i])

			if i < highestNumIndex {
				// Concatenate before

				conStrLeft := currentNumStr + highestNumStr

				convertedNum, err := strconv.Atoi(conStrLeft)

				if err != nil {
					panic(err)
				}

				if convertedNum > secondHighestNum {
					secondHighestNum = convertedNum
				}
			}

			if i > highestNumIndex {
				// Concatenate after

				conStrRight := highestNumStr + currentNumStr

				convertedNum, err := strconv.Atoi(conStrRight)

				if err != nil {
					panic(err)
				}

				if convertedNum > secondHighestNum {
					secondHighestNum = convertedNum
				}
			}
		}
	}

	return secondHighestNum
}

func FindTwelveJoltages(bank string) string {
	joltagesFound := []int{}

	lengthOfBank := len(bank)

	batteriesRequired := 12
	batteriesToRemove := lengthOfBank - batteriesRequired

	bankInts := ConvertBankSliceToInts(bank)

	for i := 0; i < len(bankInts); i++ {
		for len(joltagesFound) != 0 && joltagesFound[len(joltagesFound) - 1] < bankInts[i] && batteriesToRemove > 0 {
			joltagesFound = joltagesFound[:len(joltagesFound) - 1]
			batteriesToRemove--
		}

		joltagesFound = append(joltagesFound, bankInts[i])
	}

	if batteriesToRemove > 0 {
		joltagesFound = joltagesFound[:len(joltagesFound) - batteriesToRemove]
	}

	joltagesAsString := ConvertBankSliceToStrs(joltagesFound)

	return strings.Join(joltagesAsString, "")
}