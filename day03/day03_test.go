package day03

import "testing"

func TestDay3Part1(t *testing.T) {
	var part1Tests = []struct {
		fileName  string
		expecting int64
	}{
		{"testdata/day03_example1.dat", 6},
		{"testdata/day03_example2.dat", 159},
		{"testdata/day03_example3.dat", 135},
		{"testdata/day03.dat", 627},
	}
	for _, tt := range part1Tests {
		result := Part1(tt.fileName)
		if tt.expecting != result {
			t.Errorf("Expecting %d, got %d.", tt.expecting, result)
		}
	}
}

func TestDay3Part2(t *testing.T) {
	var part1Tests = []struct {
		fileName  string
		expecting int64
	}{
		{"testdata/day03_example1.dat", 30},
		{"testdata/day03_example2.dat", 610},
		{"testdata/day03_example3.dat", 410},
		{"testdata/day03.dat", 13190},
	}
	for _, tt := range part1Tests {
		result := Part2(tt.fileName)
		if tt.expecting != result {
			t.Errorf("Expecting %d, got %d.", tt.expecting, result)
		}
	}
}
