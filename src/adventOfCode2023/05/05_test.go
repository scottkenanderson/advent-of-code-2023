package aoc05_test

import (
	aoc05 "adventOfCode2023/05"
	"testing"
)

func tutu(a int) int {
	return a + 1
}

func Test_queryAlmanac(t *testing.T) {
	tests := []struct {
		almanac []aoc05.Almanac
		input   int
		want    int
	}{
		{almanac: []aoc05.Almanac{{Source: 25, Destination: 18, Length: 70}}, input: 81, want: 74},
		{almanac: []aoc05.Almanac{{Source: 64, Destination: 68, Length: 13}}, input: 74, want: 78},
		// {input: 2, want: 3},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := aoc05.QueryAlmanac(tt.almanac, tt.input)
			if got != tt.want {
				t.Errorf("QueryAlmanac = %d; want %d", got, tt.want)
			}
		})
	}
}
