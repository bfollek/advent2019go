package day03

import "testing"

func TestDay2Part1(t *testing.T) {
	expecting := int64(627)
	result := Part1("testdata/day03.dat")
	if expecting != result {
		t.Errorf("Expecting %d, got %d.", expecting, result)
	}
}
