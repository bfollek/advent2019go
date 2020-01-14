package day05

import (
	"log"

	"github.com/bfollek/aoc19go/intcode"
)

// Part1 "After providing 1 to the only input instruction and passing all the tests, what diagnostic code does the program produce?"
func Part1(fileName string) int {
	vm := intcode.New()
	go vm.RunFromFile(fileName)
	vm.In <- 1
	// "For each test, it will run an output instruction indicating how far
	// the result of the test was from the expected value, where 0 means the
	// test was successful. Non-zero outputs mean that a function is not working
	// correctly; check the instructions that were run before the output instruction
	//  to see which one failed.
	//
	// 	Finally, the program will output a diagnostic code and immediately halt."
	for {
		output := <-vm.Out
		if output == 0 {
			continue
		}
		// The first non-zero output should be the last output
		if _, ok := <-vm.Out; ok {
			log.Fatalf("Unexpected non-zero output: %d", output)
		}
		return output
	}
}

// Part2 "This time, when the TEST diagnostic program runs its input instruction to get the ID of the system to test, provide it 5, the ID for the ship's thermal radiator controller. This diagnostic test suite only outputs one number, the diagnostic code. What is the diagnostic code for system ID 5?"
func Part2(fileName string) int {
	vm := intcode.New()
	go vm.RunFromFile(fileName)
	vm.In <- 5
	return <-vm.Out
}
