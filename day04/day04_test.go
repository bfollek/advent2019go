package day04

import "testing"

func TestDay04IsValid(t *testing.T) {
	var isValidTests = []struct {
		password         string
		seqLenMustEqual2 bool
		expecting        bool
	}{
		{"111111", false, true},
		{"123455", false, true},
		{"223450", false, false},  // decreasing pair of digits: `50`
		{"123789", false, false},  // no seq of 2 or more
		{"11237", false, false},   // too short
		{"1123789", false, false}, // too long
		{"112233", true, true},
		{"123455", true, true},
		{"556789", true, true},
		{"125589", true, true},
		{"111122", true, true},
		{"123444", true, false},  // no seq of 2
		{"111111", true, false},  // no seq of 2
		{"223450", true, false},  // decreasing pair of digits 50
		{"123789", true, false},  // no seq of 2
		{"11237", true, false},   // too short
		{"1123789", true, false}, // too long
	}
	for _, tt := range isValidTests {
		result := isValid(tt.password, tt.seqLenMustEqual2)
		if tt.expecting != result {
			t.Errorf("Expecting %t, got %t for %s.", tt.expecting, result, tt.password)
		}
	}
}

func TestDay4Part1(t *testing.T) {
	var part1Tests = []struct {
		fileName  string
		expecting int
	}{
		{"testdata/day04.dat", 1873},
	}
	for _, tt := range part1Tests {
		result := Part1(tt.fileName)
		if tt.expecting != result {
			t.Errorf("Expecting %d, got %d.", tt.expecting, result)
		}
	}
}

func TestDay4Part2(t *testing.T) {
	var part1Tests = []struct {
		fileName  string
		expecting int
	}{
		{"testdata/day04.dat", 1264},
	}
	for _, tt := range part1Tests {
		result := Part2(tt.fileName)
		if tt.expecting != result {
			t.Errorf("Expecting %d, got %d.", tt.expecting, result)
		}
	}
}
