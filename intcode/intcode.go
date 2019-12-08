package intcode

import "log"

// NoInput is a convenience for clients.
var NoInput = []int{}

// ------------------------------------------------------------------
// Modes
// ------------------------------------------------------------------

// Your ship computer already understands parameter mode 0, position mode, which
// causes the parameter to be interpreted as a position - if the parameter is 50,
// its value is the value stored at address 50 in memory. Until now, all parameters
//  have been in position mode.

// Now, your ship computer will also need to handle parameters in mode 1,
// immediate mode. In immediate mode, a parameter is interpreted as a value -
// if the parameter is 50, its value is simply 50.

const positionMode = 0
const immediateMode = 1

// ------------------------------------------------------------------
// Opcodes
// ------------------------------------------------------------------

// Opcode 1 adds together numbers read from two positions and stores the result
// in a third position. The three integers immediately after the opcode tell you
// these three positions - the first two indicate the positions from which you
// should read the input values, and the third indicates the position at which
// the output should be stored.

// Opcode 2 works exactly like opcode 1, except it multiplies the two inputs
// instead of adding them. Again, the three integers after the opcode indicate
// where the inputs and outputs are, not their values.

// Opcode 3 takes a single integer as input and saves it to the position given
// by its only parameter. For example, the instruction 3,50 would take an input
// value and store it at address 50.

// Opcode 4 outputs the value of its only parameter. For example, the instruction
// 4,50 would output the value at address 50.

// Opcode 99 means that the program is finished and should immediately halt.
// The instruction 99 contains only an opcode and has no parameters.

// Encountering an unknown opcode means something went wrong.

const add = 1
const multiply = 2
const input = 3
const output = 4
const halt = 99

// opcode => number of params
var opCodeNumParams = map[int]int{add: 3, multiply: 3, input: 1, output: 1, halt: 0}

// Run executes an intcode program. The first param, `program`, is the program code.
// The second param, `input`, is a slice of the input values the program needs.
// The first return value is memory after the program runs.
// The second return value is a slice of the output the program creates.
//
// The `program` param is also the initial state of machine memory.
// The program may modify memory as it runs. This means that the program may be
// self-modifying.
func Run(program []int, input []int) ([]int, []int) {
	output := []int{}
	var opCode int
	instructionPointer := 0
	for {
		switch opCode = program[instructionPointer]; opCode {
		case add, multiply:
			op1 := program[program[instructionPointer+1]]
			op2 := program[program[instructionPointer+2]]
			var value int
			if opCode == add {
				value = op1 + op2
			} else {
				value = op1 * op2
			}
			program[program[instructionPointer+3]] = value
		case halt:
			return program, output
		default:
			log.Fatalf("Unexpected op code: %d", opCode)
		}
		instructionPointer += (opCodeNumParams[opCode] + 1)
	}
}
