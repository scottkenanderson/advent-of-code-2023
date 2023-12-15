package aoc02_test

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	got := aoc02.ParseLine([]string{"1", "2"})
	want := 12
	if got != want {
		t.Errorf("ParseLine(['1', '2']) = %d; want %d", got, want)
	}
}
