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

// Part2 "This time, when the TEST diagnostic program runs its input instruction to get the ID of the system to test, provide it 5, the ID for the ship's thermal radiator controller. This diagnostic test suite only outputs one number, the diagnostic code. What is the diagnostic code for system ID 5?"
func Part2(fileName string) int {
	_, output := intcode.RunFromFile(fileName, []int{5})
	fmt.Printf("output == %v\n", output)
	return output[len(output)-1]
}
