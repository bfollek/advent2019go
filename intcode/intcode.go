package intcode

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bfollek/advent2019go/util"
	"github.com/golang-collections/collections/stack"
)

// NoInput is a convenience for clients.
var NoInput = []int{}

type computer struct {
	memory         []int
	iP             int   // Instruction pointer
	input          []int // Input buffer
	inP            int   // Input pointer
	output         []int // Output buffer
	parameterModes *stack.Stack
}

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

const (
	positionMode = iota
	immediateMode
)

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

const (
	opAdd = iota + 1
	opMultiply
	opInput
	opOutput
	opHalt = 99
)

// Opcodes are two-digit numbers. A single-digit opcode has an implied leading
// zero. This comes up when we're setting parameter modes.
const opCodeLen = 2

// opcode => number of params
var opCodeNumParams = map[int]int{opAdd: 3, opMultiply: 3, opInput: 1,
	opOutput: 1, opHalt: 0}

// Run executes an intcode program.
// The first param, `program`, is the program code.
// The second param, `input`, is any input the program needs.
// The first return value is memory after the program runs.
// The second return value is the program's output.
//
// The program may modify the memory it's in as it runs.
// This means that the program may be self-modifying.
func Run(program []int, input []int) ([]int, []int) {
	vm := load(program, input)
	for {
		opCode, numParams := nextOpCode(vm)
		switch opCode {
		case opAdd:
			add(vm)
		case opMultiply:
			multiply(vm)
		case opInput:
			in(vm)
		case opOutput:
			out(vm)
		case opHalt:
			return vm.memory, vm.output
		default:
			log.Fatalf("Unexpected op code: %d", opCode)
		}
		vm.iP += (numParams + 1)
	}
}

// Parameter modes are stored in the same value as the instruction's opcode.
// The opcode is a two-digit number based only on the ones and tens digit
// of the value, that is, the opcode is the rightmost two digits of the
// first value in an instruction.
func nextOpCode(vm *computer) (int, int) {
	// Add a leading zero to the opCode, if necessary.
	s := strconv.Itoa(vm.memory[vm.iP])
	if len(s) < opCodeLen {
		s = "0" + s
	}
	// Rightmost chars are the opCode
	opCode := util.MustAtoi(s[len(s)-opCodeLen:])
	numParams := opCodeNumParams[opCode]
	// Leftmost chars are the parameter modes
	modes := s[0 : len(s)-opCodeLen]
	setParameterModes(modes, numParams, vm)
	return opCode, numParams
}

// Parameter modes are single digits, one per
// parameter, read right-to-left from the opcode: the first parameter's mode
// is in the hundreds digit, the second parameter's mode is in the thousands
// digit, the third parameter's mode is in the ten-thousands digit, and so on.
// Any missing modes are 0.
func setParameterModes(modes string, numParams int, vm *computer) {
	// Add any missing leading zeros (the default) for the parameter modes.
	lenPrefix := numParams - len(modes)
	modes += strings.Repeat("0", lenPrefix)
	vm.parameterModes = stack.New()
	for i := len(modes) - 1; i >= 0; i-- {
		vm.parameterModes.Push(util.CharToIntValue(modes[i]))
	}
	fmt.Println(modes)
}

func add(vm *computer) {
	op1, op2 := next2Params(vm)
	store(op1+op2, vm.memory[vm.iP+3], vm)
}

func multiply(vm *computer) {
	op1, op2 := next2Params(vm)
	store(op1*op2, vm.memory[vm.iP+3], vm)
}

func in(vm *computer) {
	i := vm.input[vm.inP]
	vm.inP++
	store(i, vm.memory[vm.iP+1], vm)
}

func out(vm *computer) {
	i := fetchPosition(vm.iP+1, vm)
	vm.output = append(vm.output, i)
}

func next2Params(vm *computer) (int, int) {
	return fetchPosition(vm.iP+1, vm), fetchPosition(vm.iP+2, vm)
}

// fetchImmediate interprets the `address` param as the address of the
// value to return. If the `address` param is 50, we return the value
// stored at address 50.
func fetchImmediate(address int, vm *computer) int {
	return vm.memory[address]
}

// fetchPosition adds a level of indirection. It interprets the `address`
// param as the address of an address. The second address is the address of
// the value to return. If the `address` param is 50, we get the value stored
// at address 50. Suppose that value is 100. We then get the value stored at
// address 100, and return it.
func fetchPosition(address int, vm *computer) int {
	i := vm.memory[address]
	return fetchImmediate(i, vm)
}

func store(value int, address int, vm *computer) {
	vm.memory[address] = value
}

// load creates the vm and loads the program into it.
func load(program []int, input []int) *computer {
	vm := new(computer)
	vm.memory = make([]int, len(program))
	copy(vm.memory, program)
	vm.iP = 0
	vm.input = input
	vm.inP = 0
	vm.output = []int{}
	return vm
}
