package main

import (
	aoc01 "adventOfCode2023/01"
	aoc02 "adventOfCode2023/02"
	aoc03 "adventOfCode2023/03"
	aoc04 "adventOfCode2023/04"
	aoc05 "adventOfCode2023/05"
	"os"
)

func main() {
	arg := os.Args[1]
	filename := os.Args[2]
	switch arg {
	case "01":
		day01("01/" + filename)
	case "02":
		day02("02/" + filename)
	case "03":
		day03("03/" + filename)
	case "04":
		day04("04/" + filename)
	case "05":
		day05("05/" + filename)
	}
}

func day01(filename string) {
	aoc01.Part1(filename)
	aoc01.Part2(filename)
}

func day02(filename string) {
	aoc02.Part1(filename)
	aoc02.Part2(filename)
}

func day03(filename string) {
	aoc03.Part1(filename)
	aoc03.Part2(filename)
}

func day04(filename string) {
	aoc04.Part1(filename)
	aoc04.Part2(filename)
}

func day05(filename string) {
	aoc05.Part1(filename)
	aoc05.Part2(filename)
}
