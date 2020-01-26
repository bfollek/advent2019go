package util

import (
	"reflect"
	"testing"
)

func TestMustLoadIntSlice(t *testing.T) {
	type runTest struct {
		fileName        string
		sep             string
		expectingResult []int
	}
	var runTests = []runTest{
		{
			"testdata/ints_with_commas.dat",
			",",
			[]int{1, 2, 3},
		},
		{
			"testdata/ints_without_seps.dat",
			"",
			[]int{4, 5, 6, 7, 8},
		},
	}

	for _, test := range runTests {
		result := MustLoadIntSlice(test.fileName, test.sep)
		if !reflect.DeepEqual(result, test.expectingResult) {
			t.Errorf("Run output: expecting [%v], got [%v]", test.expectingResult, result)
		}
	}
}
