package day04

import "testing"

func TestDay04IsValid(t *testing.T) {
	var isValidTests = []struct {
		password  string
		expecting bool
	}{
		{"111111", true},
		{"123455", true},
		{"223450", false},  // decreasing pair of digits 50
		{"123789", false},  // no double
		{"11237", false},   // too short
		{"1123789", false}, // too long
	}
	for _, tt := range isValidTests {
		result := isValid(tt.password)
		if tt.expecting != result {
			t.Errorf("Expecting %t, got %t.", tt.expecting, result)
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
