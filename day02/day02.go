package day02

import (
	"errors"

	"github.com/bfollek/advent2019go/intcode"
	"github.com/bfollek/advent2019go/util"
)

const moonLanding = 19690720

// Part1 "What value is left at position 0 after the program halts?"
func Part1(fileName string) int {
	program := loadProgram(fileName)
	// "...before running the program, replace position 1 with the value 12 and replace
	// position 2 with the value 2.
	program[1] = 12
	program[2] = 2
	program = intcode.RunProgram(program)
	return program[0]
}

// Part2 "...you need to determine what pair of inputs produces the output 19690720."
func Part2(fileName string) (int, error) {
	cleanMemory := loadProgram(fileName)
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			// Slice assignments overlap and would clobber cleanMemory, so...
			program := make([]int, len(cleanMemory))
			copy(program, cleanMemory)
			program[1] = i
			program[2] = j
			program = intcode.RunProgram(program)
			if program[0] == moonLanding {
				return 100*i + j, nil
			}
		}
	}
	return -1, errors.New("No solution found")
}

func loadProgram(fileName string) []int {
	ss := util.MustLoadStringSlice(fileName, ",")
	program := []int{}
	for _, s := range ss {
		i := util.MustAtoi(s)
		program = append(program, i)
	}
	return program
}
