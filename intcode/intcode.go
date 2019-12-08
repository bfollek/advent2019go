package intcode

import "log"

const add = 1
const multiply = 2
const halt = 99

// RunProgram - "The three integers immediately after the opcode tell you these
// three positions - the first two indicate the positions from which you should read
// the input values, and the third indicates the position at which the output should
// be stored."
func RunProgram(program []int) []int {
	opCodeIndex := 0
	for {
		switch opCode := program[opCodeIndex]; opCode {
		case add, multiply:
			op1 := program[program[opCodeIndex+1]]
			op2 := program[program[opCodeIndex+2]]
			var value int
			if opCode == add {
				value = op1 + op2
			} else {
				value = op1 * op2
			}
			program[program[opCodeIndex+3]] = value
		case halt:
			return program
		default:
			log.Fatalf("Unexpected op code: %d", opCode)
		}
		// "Once you're done processing an opcode, move to the next one
		// by stepping forward 4 positions."
		opCodeIndex += 4
	}
}
