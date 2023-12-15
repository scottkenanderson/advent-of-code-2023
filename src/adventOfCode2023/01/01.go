package aoc01

import (
	"adventOfCode2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func Part1(filename string) {
	input := utils.SplitLinesFromFile(filename)

	var output int
	for _, line := range input {
		lineSplit := strings.Split(line, "")

		output += ParseLine(lineSplit)
	}
	fmt.Println(output)
}

func Part2(filename string) {
	input := utils.SplitLinesFromFile(filename)

	var output int
	for _, line := range input {
		lineSplit := strings.Split(line, "")
		numbers := getNumbers(lineSplit)

		output += ParseLine(numbers)
	}
	fmt.Println(output)
}

func ParseLine(nums []string) int {
	var numbers []string
	for _, c := range nums {
		_, err := strconv.ParseInt(c, 10, 0)
		if err == nil {
			numbers = append(numbers, c)
		}
	}

	first := numbers[0]
	last := numbers[len(numbers)-1]
	number, _ := strconv.ParseInt(first+last, 10, 0)

	return int(number)
}

var numberMapping = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func getNumbers(line []string) []string {
	var numbers []string

	for i, c := range line {
		_, err := strconv.ParseInt(c, 10, 0)
		if err == nil {
			numbers = append(numbers, c)
			continue
		}

		for j := i + 3; j <= len(line) && j <= i+5; j++ {
			word := strings.Join(line[i:j], "")
			if num, found := numberMapping[word]; found {
				numbers = append(numbers, num)
				break
			}
		}
	}
	return numbers
}
