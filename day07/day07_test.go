package day07

import "testing"

func TestDay7Part1(t *testing.T) {
	var part1Tests = []struct {
		desc      string
		fileName  string
		expecting int
	}{
		{"short data file", "testdata/day07_example1.dat", 43210},
		{"full data file", "testdata/day07.dat", 18812},
	}
	for _, tt := range part1Tests {
		result := Part1(tt.fileName)
		if tt.expecting != result {
			t.Errorf("%s: expecting %d, got %d.", tt.desc, tt.expecting, result)
		}
	}
}
