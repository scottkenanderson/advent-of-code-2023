package aoc02

import (
	"adventOfCode2023/utils"
	"fmt"
	"regexp"
	"strings"
)

type cubes struct {
	red, green, blue int
}

var limits = cubes{
	red:   12,
	green: 13,
	blue:  14,
}

func Part1(filename string) {
	input := utils.SplitLinesFromFile(filename)
	gameRecordRegExp := regexp.MustCompile(`^Game (\d+): (.*)`)
	var output int

	for _, line := range input {
		match := gameRecordRegExp.FindStringSubmatch(line)
		_, underLimit := getCubes(match[2])
		gameID := utils.StringToInt(match[1])

		if underLimit {
			output += gameID
		}
	}
	fmt.Println(output)
}

func Part2(filename string) {
	input := utils.SplitLinesFromFile(filename)
	gameRecordRegExp := regexp.MustCompile(`^Game (\d+): (.*)`)
	var output int

	for _, line := range input {
		match := gameRecordRegExp.FindStringSubmatch(line)
		gameCubes, _ := getCubes(match[2])
		var maxCubes cubes
		for _, cube := range gameCubes {
			maxCubes = calculateMaxCubes(maxCubes, cube)
		}

		output += cubePower(maxCubes)
	}
	fmt.Println(output)
}

func getCubes(input string) ([]cubes, bool) {
	var cubeRecords []cubes
	ok := true
	for _, cubeRecordString := range strings.Split(input, "; ") {
		newGameCubes := getCube(cubeRecordString)
		if !cubeRecordIsUnderLimit(newGameCubes) {
			ok = false
		}
		cubeRecords = append(cubeRecords, newGameCubes)
	}

	return cubeRecords, ok
}

func calculateMaxCubes(gameCubes, newCubeRecord cubes) cubes {
	gameCubes.green = max(gameCubes.green, newCubeRecord.green)
	gameCubes.red = max(gameCubes.red, newCubeRecord.red)
	gameCubes.blue = max(gameCubes.blue, newCubeRecord.blue)

	return gameCubes
}

func getCube(cubeRecord string) cubes {
	var gameCubes cubes

	gameCubes.green += matchColour("green", cubeRecord)
	gameCubes.red += matchColour("red", cubeRecord)
	gameCubes.blue += matchColour("blue", cubeRecord)

	return gameCubes
}

func cubePower(cubes cubes) int {
	return cubes.green * cubes.red * cubes.blue
}

func matchColour(colour, cubeRecord string) int {
	colourCubesRegExp := regexp.MustCompile(`(\d+) ` + colour)
	match := colourCubesRegExp.FindStringSubmatch(cubeRecord)
	if len(match) > 0 {
		return utils.StringToInt(match[1])
	}
	return 0
}

func cubeRecordIsUnderLimit(cubeRecord cubes) bool {
	return cubeRecord.red <= limits.red && cubeRecord.green <= limits.green && cubeRecord.blue <= limits.blue
}
