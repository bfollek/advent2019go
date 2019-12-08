package day02

import "testing"

func TestDay2Part1(t *testing.T) {
	expecting := 2692315
	result := Part1("testdata/day02.dat")
	if expecting != result {
		t.Errorf("Expecting %d, got %d.", expecting, result)
	}
}

func TestDay2Part2(t *testing.T) {
	expecting := 9507
	result, err := Part2("testdata/day02.dat")
	if err != nil {
		t.Errorf("Expecting no error, got %s.", err)
	}
	if expecting != result {
		t.Errorf("Expecting %d, got %d.", expecting, result)
	}
}
