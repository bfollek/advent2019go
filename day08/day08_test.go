package day08

import (
	"fmt"
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
		{"short data file", 2, 3, "testdata/day08_part1_example.dat", 1},
		{"full data file", 6, 25, "testdata/day08.dat", 1596},
	}
	for _, tt := range part1Tests {
		result := Part1(tt.height, tt.width, tt.fileName)
		if tt.expecting != result {
			t.Errorf("%s: expecting %d, got %d.", tt.desc, tt.expecting, result)
		}
	}
}

func TestDay8Part2(t *testing.T) {
	var part2Tests = []struct {
		desc      string
		height    int
		width     int
		fileName  string
		expecting string
	}{
		{"short data file", 2, 2, "testdata/day08_part2_example.dat",
			fmt.Sprintf("%v", Img{{{0, 1}, {1, 0}}})},
	}
	for _, tt := range part2Tests {
		result := fmt.Sprintf("%v", Part2(tt.height, tt.width, tt.fileName))
		if tt.expecting != result {
			t.Errorf("%s: expecting %s, got %s.", tt.desc, tt.expecting, result)
		}
	}
	// No test here. Just run and look at the output.
	Part2(6, 25, "testdata/day08.dat")
	// Answer is LBRCE
	/*
	 *    ***  ***   **  ****
	 *    *  * *  * *  * *
	 *    ***  *  * *    ***
	 *    *  * ***  *    *
	 *    *  * * *  *  * *
	 **** ***  *  *  **  ****
	 */
}
