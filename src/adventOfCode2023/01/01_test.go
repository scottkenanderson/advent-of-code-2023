package aoc01_test

import (
	aoc01 "adventOfCode2023/01"
	"testing"
)

func TestParseLine(t *testing.T) {
	got := aoc01.ParseLine([]string{"1", "2"})
	want := 12
	if got != want {
		t.Errorf("ParseLine(['1', '2']) = %d; want %d", got, want)
	}
}
