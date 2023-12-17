package aoc05

import (
	"adventOfCode2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(filename string) {
	input := utils.ReadFile(filename)

	seeds := parseSeeds(input)

	findLocation(seeds, input)
}

func Part2(filename string) {
	input := utils.ReadFile(filename)

	seeds := parseSeedRanges(input)

	findLocation(seeds, input)
}

func findLocation(seeds []int, input string) {
	seedToSoilMap := parseCategoryMap(input, "seed-to-soil")
	utils.Print(seedToSoilMap)
	soilToFertilizerMap := parseCategoryMap(input, "soil-to-fertilizer")
	utils.Print(soilToFertilizerMap)
	fertilizerToWaterMap := parseCategoryMap(input, "fertilizer-to-water")
	utils.Print(fertilizerToWaterMap)
	waterToLightMap := parseCategoryMap(input, "water-to-light")
	utils.Print(waterToLightMap)
	lightToTemperatureMap := parseCategoryMap(input, "light-to-temperature")
	utils.Print(lightToTemperatureMap)
	temperatureToHumidityMap := parseCategoryMap(input, "temperature-to-humidity")
	utils.Print(temperatureToHumidityMap)
	humidityToLocationMap := parseCategoryMap(input, "humidity-to-location")
	utils.Print(humidityToLocationMap)

	locations := make(map[int]struct{})

	for _, seed := range seeds {
		soil := QueryAlmanac(seedToSoilMap, seed)
		fertilizer := QueryAlmanac(soilToFertilizerMap, soil)
		water := QueryAlmanac(fertilizerToWaterMap, fertilizer)
		light := QueryAlmanac(waterToLightMap, water)
		temperature := QueryAlmanac(lightToTemperatureMap, light)
		humidity := QueryAlmanac(temperatureToHumidityMap, temperature)
		location := QueryAlmanac(humidityToLocationMap, humidity)
		utils.Print(location)
		locations[location] = struct{}{}
	}

	minLocation := int(^uint(0) >> 1)
	for location := range locations {
		minLocation = min(minLocation, location)
	}
	utils.Print(minLocation)

}

func parseSeeds(input string) []int {
	seedsRegex := regexp.MustCompile(`^seeds: ([\d ]+)`)
	matches := seedsRegex.FindStringSubmatch(input)
	var seeds []int
	for _, match := range strings.Split(matches[1], " ") {
		seed, _ := strconv.ParseInt(match, 10, 32)
		seeds = append(seeds, int(seed))
	}
	return seeds
}

func parseSeedRanges(input string) []int {
	seedsRegex := regexp.MustCompile(`^seeds: ([\d ]+)`)
	matches := seedsRegex.FindStringSubmatch(input)
	var seeds []int
	numbers := strings.Split(matches[1], " ")
	for i := range numbers {
		if (i % 2) == 1 {
			continue
		}
		seedStart, _ := strconv.ParseInt(numbers[i], 10, 32)
		seedRange, _ := strconv.ParseInt(numbers[i+1], 10, 32)

		for i := 0; i < int(seedRange); i++ {
			utils.Print(int(seedStart) + i)
			seeds = append(seeds, int(seedStart)+i)
		}
	}
	return seeds
}

type Almanac struct {
	Source, Destination, Length int
}

func parseCategoryMap(input, mapName string) []Almanac {
	almanacRegex := regexp.MustCompile(mapName + ` map:\n(\d+ \d+ \d+(?:\n\d+ \d+ \d+)+)`)
	matches := almanacRegex.FindStringSubmatch(input)
	return makeCategoryMap(matches[1])
}

func makeCategoryMap(match string) []Almanac {
	// categoryMap := make(map[int]int)
	var categoryMap []Almanac

	categoryMapRegex := regexp.MustCompile(`(\d+) (\d+) (\d+)`)
	for _, line := range strings.Split(match, "\n") {
		categoryMatches := categoryMapRegex.FindStringSubmatch(line)
		destinationStr, sourceStr, lengthStr := categoryMatches[1], categoryMatches[2], categoryMatches[3]
		length, _ := strconv.ParseInt(lengthStr, 10, 32)
		destination, _ := strconv.ParseInt(destinationStr, 10, 32)
		source, _ := strconv.ParseInt(sourceStr, 10, 32)
		entry := Almanac{
			Source:      int(source),
			Destination: int(destination),
			Length:      int(length),
		}
		categoryMap = append(categoryMap, entry)
	}

	return categoryMap
}

func QueryAlmanac(almanac []Almanac, target int) int {
	for _, entry := range almanac {
		if target >= entry.Source && target <= entry.Source+entry.Length {
			offset := target - entry.Source
			result := entry.Destination + offset
			fmt.Printf("%d -> %d ", target, result)
			return result
		}
	}

	fmt.Printf("%d -> %d ", target, target)
	return target
}
