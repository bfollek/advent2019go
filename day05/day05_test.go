package day05

import "testing"

func TestDay5Part1(t *testing.T) {
	var part1Tests = []struct {
		fileName  string
		expecting int
	}{
		{"testdata/day05.dat", 4511442},
	}
	for _, tt := range part1Tests {
		result := Part1(tt.fileName)
		if tt.expecting != result {
			t.Errorf("Expecting %d, got %d.", tt.expecting, result)
		}
	}
}
