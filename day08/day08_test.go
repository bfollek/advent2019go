package day08

import (
	"testing"
)

func TestDay8Part1(t *testing.T) {
	var part1Tests = []struct {
		desc      string
		width     int
		height    int
		fileName  string
		expecting int
	}{
		{"short data file", 3, 2, "testdata/day08_example.dat", 1},
		//{"full data file", 25, 6, "testdata/day08.dat", 1596},
	}
	for _, tt := range part1Tests {
		result := Part1(tt.width, tt.height, tt.fileName)
		if tt.expecting != result {
			t.Errorf("%s: expecting %d, got %d.", tt.desc, tt.expecting, result)
		}
	}
}
