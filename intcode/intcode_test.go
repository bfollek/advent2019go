package intcode

import (
	"testing"
)

func TestIntcodeRun(t *testing.T) {
	type runTest struct {
		program         []int
		input           []int
		expectingMemory []int
		expectingOutput []int
	}
	var runTests = []runTest{
		{
			[]int{1, 0, 0, 0, 99},
			[]int{},
			[]int{2, 0, 0, 0, 99},
			[]int{},
		},
		{
			[]int{2, 3, 0, 3, 99},
			[]int{},
			[]int{2, 3, 0, 6, 99},
			[]int{},
		},
		{
			[]int{2, 4, 4, 5, 99, 0},
			[]int{},
			[]int{2, 4, 4, 5, 99, 9801},
			[]int{},
		},
		{
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int{},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
			[]int{},
		},
		{
			[]int{3, 0, 4, 0, 99}, // Input 78
			[]int{78},
			[]int{78, 0, 4, 0, 99}, // Output 78
			[]int{78},
		},
		{
			[]int{1001, 5, -3000, 5, 99, 4000}, // Add -3000 to 4000
			[]int{},
			[]int{1001, 5, -3000, 5, 99, 1000},
			[]int{},
		},
		{
			[]int{101, -123, 5, 5, 99, 246}, // Add -123 to 246
			[]int{},
			[]int{101, -123, 5, 5, 99, 123},
			[]int{},
		},
		{
			// Using position mode, consider whether the input is less than 8;
			// output 1 (if it is) or 0 (if it is not). This is the less than test.
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{2},
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, 1, 8},
			[]int{1},
		},
		{
			// Using position mode, consider whether the input is less than 8;
			// output 1 (if it is) or 0 (if it is not). This is the not less than test.
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{22},
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, 0, 8},
			[]int{0},
		},
		{
			// Using position mode, consider whether the input is equal to 8;
			// output 1 (if it is) or 0 (if it is not). This is the equal test.
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{8},
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, 1, 8},
			[]int{1},
		},
		{
			// Using position mode, consider whether the input is equal to 8;
			// output 1 (if it is) or 0 (if it is not). This is the not equal test.
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{22},
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, 0, 8},
			[]int{0},
		},
		{
			// Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero:
			// (using position mode)
			//    0, 1,  2, 3,  4,  5, 6,  7,  8,  9, 10, 11, 12,13,14,15
			[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			[]int{22},
			[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 22, 1, 1, 9},
			[]int{1},
		},
	}

	for _, test := range runTests {
		vm := New()
		for _, i := range test.input {
			vm.In <- i
		}
		go vm.Run(test.program) // This can be above or below the for loop
		for idx, i := range test.expectingOutput {
			j := <-vm.Out
			if i != j {
				t.Errorf("Run output: expecting [%d] at index [%d], got [%d]", i, idx, j)
			}
		}
		for idx, i := range test.expectingMemory {
			j := <-vm.Mem // Won't happen till the goroutine exits
			if i != j {
				t.Errorf("Run memory: expecting [%d] at index [%d], got [%d]", i, idx, j)
			}
		}
	}
}
