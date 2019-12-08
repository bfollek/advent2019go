package intcode

import (
	"reflect"
	"testing"
)

func TestIntcodeRunProgram(t *testing.T) {
	type runProgramTest struct {
		program   []int
		expecting []int
	}
	var runProgramTests = []runProgramTest{
		{
			[]int{1, 0, 0, 0, 99},
			[]int{2, 0, 0, 0, 99},
		},
		{
			[]int{2, 3, 0, 3, 99},
			[]int{2, 3, 0, 6, 99},
		},
		{
			[]int{2, 4, 4, 5, 99, 0},
			[]int{2, 4, 4, 5, 99, 9801},
		},
		{
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}
	for _, rpt := range runProgramTests {
		if result := RunProgram(rpt.program); !reflect.DeepEqual(rpt.expecting, result) {
			t.Errorf("runProgram: expecting [%v], got [%v]", rpt.expecting, result)
		}
	}
}
