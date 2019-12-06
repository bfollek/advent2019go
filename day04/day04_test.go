package day04

import "testing"

func TestDay04IsValid(t *testing.T) {
	var isValidTests = []struct {
		password  string
		expecting bool
	}{
		{"111111", true},
		{"123455", true},
		{"223450", false},  // decreasing pair of digits: `50`
		{"123789", false},  // no seq of 2
		{"11237", false},   // too short
		{"1123789", false}, // too long
	}
	for _, tt := range isValidTests {
		result := isValid(tt.password)
		if tt.expecting != result {
			t.Errorf("Expecting %t, got %t for %s.", tt.expecting, result, tt.password)
		}
	}
}

func TestDay04IsValidWithSeqOf2(t *testing.T) {
	var isValidWithSeqOf2Tests = []struct {
		password  string
		expecting bool
	}{
		{"112233", true},
		{"123455", true},
		{"556789", true},
		{"125589", true},
		{"111122", true},
		{"123444", false},  // no seq of 2
		{"111111", false},  // no seq of 2
		{"223450", false},  // decreasing pair of digits 50
		{"123789", false},  // no seq of 2
		{"11237", false},   // too short
		{"1123789", false}, // too long
	}
	for _, tt := range isValidWithSeqOf2Tests {
		result := isValidWithSeqOf2(tt.password)
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
