package aoc03

import (
	"adventOfCode2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type number struct {
	value    int
	position utils.CoordinateRange
}

func Part1(filename string) {
	input := utils.ReadFile(filename)

	grid := buildGrid(input)
	numbers := findNumbers(grid)
	parts := findParts(numbers, grid)
	var output int
	for _, part := range parts {
		output += part
	}
	fmt.Println(output)
}

func Part2(filename string) {
	input := utils.ReadFile(filename)

	grid := buildGrid(input)
	numbers := findNumbers(grid)
	gears := findGears(grid)
	gearRatios := calculateGearsRatios(gears, numbers)

	var output int
	for _, ratio := range gearRatios {
		output += ratio
	}
	fmt.Println(output)
}

func buildGrid(input string) [][]string {
	var grid [][]string
	for _, rowString := range strings.Split(input, "\n") {
		var row []string
		for _, col := range strings.Split(rowString, "") {
			row = append(row, col)
		}
		grid = append(grid, row)
	}
	return grid
}

func findNumbers(grid [][]string) []number {
	var numbers []number
	numberRegex := regexp.MustCompile(`\d`)
	for j, row := range grid {
		var buffer string
		startPos := utils.Coordinate{Y: j}
		endPos := utils.Coordinate{Y: j}
		for i, col := range row {
			if numberRegex.MatchString(col) {
				if buffer == "" {
					startPos.X = i
				}

				buffer += col
			}
			if !numberRegex.MatchString(col) || i == len(grid[0])-1 {
				if buffer != "" {
					endPos.X = i - 1
					value, _ := strconv.ParseInt(buffer, 10, 0)
					position := utils.CoordinateRange{
						Start: startPos,
						End:   endPos,
					}
					numbers = append(numbers, number{
						value:    int(value),
						position: makeSearchRange(position, len(grid[0])-1, len(grid)-1),
					})
					buffer = ""
				}
			}
		}
	}
	return numbers
}

func findGears(grid [][]string) []utils.Coordinate {
	var gears []utils.Coordinate
	for j, row := range grid {
		for i, col := range row {
			if col == "*" {
				gears = append(gears, utils.Coordinate{X: i, Y: j})
			}
		}
	}
	return gears
}

func findParts(numbers []number, grid [][]string) []int {
	var parts []int
	for _, number := range numbers {
		if isPartNumber(number, grid) {
			parts = append(parts, number.value)
		}
	}
	return parts
}

func makeSearchRange(r utils.CoordinateRange, maxX, maxY int) utils.CoordinateRange {
	return utils.CoordinateRange{
		Start: utils.Coordinate{
			X: max(r.Start.X-1, 0),
			Y: max(r.Start.Y-1, 0),
		},
		End: utils.Coordinate{
			X: min(r.End.X+1, maxX),
			Y: min(r.End.Y+1, maxY),
		},
	}
}

func isPartNumber(number number, grid [][]string) bool {
	symbolRegex := regexp.MustCompile(`(\d)`)
	searchRange := number.position
	for j := searchRange.Start.Y; j <= searchRange.End.Y; j++ {
		for i := searchRange.Start.X; i <= searchRange.End.X; i++ {
			char := grid[j][i]
			if char != "." && !symbolRegex.MatchString(char) {
				return true
			}
		}
	}
	// for j := searchRange.Start.Y; j <= searchRange.End.Y; j++ {
	// 	for i := searchRange.Start.X; i <= searchRange.End.X; i++ {
	// 		char := grid[j][i]
	// 		fmt.Print(char)
	// 	}
	// 	fmt.Println()
	// }
	return false
}

func calculateGearsRatios(gears []utils.Coordinate, numbers []number) []int {
	var gearRatios []int
	for _, gear := range gears {
		var matchingNumbers []int
		for _, number := range numbers {
			if utils.IsInRange(gear, number.position) {
				matchingNumbers = append(matchingNumbers, number.value)
			}
		}
		if len(matchingNumbers) == 2 {
			gearRatios = append(gearRatios, matchingNumbers[0]*matchingNumbers[1])
		}
	}
	return gearRatios
}
