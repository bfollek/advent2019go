package day06

import "testing"

func TestDay6Part1(t *testing.T) {
	var part1Tests = []struct {
		desc      string
		fileName  string
		expecting int
	}{
		{"short data file", "testdata/day06_part1_short.dat", 42},
		{"full data file", "testdata/day06.dat", 158090},
	}
	for _, tt := range part1Tests {
		result := Part1(tt.fileName)
		if tt.expecting != result {
			t.Errorf("%s: expecting %d, got %d.", tt.desc, tt.expecting, result)
		}
	}
}

func TestDay6Part2(t *testing.T) {
	var part2Tests = []struct {
		desc      string
		fileName  string
		expecting int64
	}{
		{"short data file", "testdata/day06_part2_short.dat", 4},
		{"full data file", "testdata/day06.dat", 241},
	}
	for _, tt := range part2Tests {
		result := Part2(tt.fileName)
		if tt.expecting != result {
			t.Errorf("%s: expecting %d, got %d.", tt.desc, tt.expecting, result)
		}
	}
}
