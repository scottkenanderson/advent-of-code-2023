package utils

import (
	"os"
	"strings"
)

func SplitLinesFromFile(filename string) []string {
	raw_input := ReadFile(filename)
	return strings.Split(raw_input, "\n")
}

func ReadFile(filename string) string {
	dat, err := os.ReadFile(filename)
	check(err)
	return string(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
