package aoc04

import (
	"adventOfCode2023/utils"
	"regexp"
	"strings"
)

var cardRegexp = regexp.MustCompile(`^Card +(\d+): (.*)$`)
var numberRegexp = regexp.MustCompile(`\d+`)

func Part1(filename string) {
	input := utils.SplitLinesFromFile(filename)
	var output int

	for _, cardInput := range input {
		matchingNumbers := findMatchingNumbers(cardInput)
		score := 0
		for i := 0; i < len(matchingNumbers); i++ {
			if score == 0 {
				score = 1
				continue
			}
			score *= 2
		}
		output += score

	}
	utils.Print(output)
}

func Part2(filename string) {
	input := utils.SplitLinesFromFile(filename)
	var output int
	cardCounts := make([]int, len(input))
	for i := range cardCounts {
		cardCounts[i] = 1
	}

	for i, cardInput := range input {
		matchingNumbers := findMatchingNumbers(cardInput)
		for j := 1; j <= len(matchingNumbers); j++ {
			if i+j < len(cardCounts) {
				cardCounts[i+j] += cardCounts[i]
			}
		}
	}

	for _, count := range cardCounts {
		output += count
	}
	utils.Print(output)
}

func findMatchingNumbers(cardInput string) []string {
	match := cardRegexp.FindStringSubmatch(cardInput)
	// cardID, _ := strconv.ParseInt(match[1], 10, 0)
	numberString := match[2]
	numbersSplit := strings.Split(numberString, "|")

	winningNumbersStr, cardNumbersStr := numbersSplit[0], numbersSplit[1]

	winningNumbersMatches := numberRegexp.FindAllString(winningNumbersStr, -1)
	cardNumbersMatches := numberRegexp.FindAllString(cardNumbersStr, -1)

	winningNumbers := make(map[string]struct{})
	for _, number := range winningNumbersMatches {
		winningNumbers[number] = struct{}{}
	}

	var matchingNumbers []string

	for _, number := range cardNumbersMatches {
		if _, ok := winningNumbers[number]; ok {
			matchingNumbers = append(matchingNumbers, number)
		}
	}
	return matchingNumbers
}
