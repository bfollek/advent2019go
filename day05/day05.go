package day05

import (
	"github.com/bfollek/advent2019go/intcode"
	"github.com/bfollek/advent2019go/util"
)

// Part1 "After providing 1 to the only input instruction and passing all the tests, what diagnostic code does the program produce?"
func Part1(fileName string) int {
	program := loadProgram(fileName)
	_, output := intcode.Run(program, []int{1})
	//fmt.Printf("output == %v\n", output)
	return output[len(output)-1]
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
