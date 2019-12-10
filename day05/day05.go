package day05

import (
	"fmt"

	"github.com/bfollek/advent2019go/intcode"
)

// Part1 "After providing 1 to the only input instruction and passing all the tests, what diagnostic code does the program produce?"
func Part1(fileName string) int {
	_, output := intcode.RunFromFile(fileName, []int{1})
	fmt.Printf("output == %v\n", output)
	return output[len(output)-1]
}
