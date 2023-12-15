package main

import (
	aoc01 "adventOfCode2023/01"
	aoc02 "adventOfCode2023/02"
	"os"
)

func main() {
	arg := os.Args[1]
	switch arg {
	case "01":
		day01()
	case "02":
		day02()
	}
}

func day01() {
	filename := os.Args[2]
	aoc01.Part1(filename)
	aoc01.Part2(filename)
}

func day02() {
	filename := os.Args[2]
	aoc02.Part1(filename)
	aoc02.Part2(filename)
}
