package day01

import (
	"testing"
)

func TestFuelRequired(t *testing.T) {
	type fuelRequiredTest struct {
		mass      int
		expecting int
	}
	var fuelRequiredTests = []fuelRequiredTest{
		{
			12,
			2,
		},
		{
			14,
			2,
		},
		{
			1969,
			654,
		},
		{
			100756,
			33583,
		},
	}
	for _, tst := range fuelRequiredTests {
		if result := fuelRequired(tst.mass); tst.expecting != result {
			t.Errorf("fuelRequired: expecting [%v], got [%v]", tst.expecting, result)
		}
	}
}

func TestFuelRequiredMeta(t *testing.T) {
	type fuelRequiredMetaTest struct {
		mass      int
		expecting int
	}
	var fuelRequiredMetaTests = []fuelRequiredMetaTest{
		{
			14,
			2,
		},
		{
			1969,
			966,
		},
		{
			100756,
			50346,
		},
	}
	for _, tst := range fuelRequiredMetaTests {
		if result := fuelRequiredMeta(tst.mass); tst.expecting != result {
			t.Errorf("fuelRequiredMeta: expecting [%v], got [%v]", tst.expecting, result)
		}
	}
}

func TestPart1(t *testing.T) {
	expecting := 3246455
	result := Part1("testdata/day01.dat")
	if expecting != result {
		t.Errorf("Expecting %d, got %d.", expecting, result)
	}
}

func TestPart2(t *testing.T) {
	expecting := 4866824
	result := Part2("testdata/day01.dat")
	if expecting != result {
		t.Errorf("Expecting %d, got %d.", expecting, result)
	}
}
