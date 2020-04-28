package day08

import (
	"testing"
)

func TestDay8Part1(t *testing.T) {
	var part1Tests = []struct {
		desc      string
		height    int
		width     int
		fileName  string
		expecting int
	}{
		{"short data file", 2, 3, "testdata/day08_example.dat", 1},
		{"full data file", 6, 25, "testdata/day08.dat", 1596},
	}
	for _, tt := range part1Tests {
		result := Part1(tt.height, tt.width, tt.fileName)
		if tt.expecting != result {
			t.Errorf("%s: expecting %d, got %d.", tt.desc, tt.expecting, result)
		}
	}
}
