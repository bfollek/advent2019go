package intcode

import (
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
	iP             int          // Instruction pointer
	input          *stack.Stack // Input buffer
	output         []int        // Output buffer
	parameterModes *stack.Stack
}

// ------------------------------------------------------------------
// Modes
// ------------------------------------------------------------------

const (
	positionMode = iota
	immediateMode
)

// ------------------------------------------------------------------
// Opcodes
// ------------------------------------------------------------------

const (
	opAdd = iota + 1
	opMultiply
	opInput
	opOutput
	opJumpIfTrue
	opJumpIfFalse
	opLessThan
	opEquals
	opHalt = 99
)

// Opcodes are two-digit numbers. A single-digit opcode has an implied leading
// zero. This comes up when we're setting parameter modes.
const opCodeLen = 2

// opcode => number of params
var opCodeNumParams = map[int]int{opAdd: 3, opMultiply: 3, opInput: 1,
	opOutput: 1, opJumpIfTrue: 2, opJumpIfFalse: 2, opLessThan: 3, opEquals: 3,
	opHalt: 0}

// RunFromFile reads an intcode program from a file, then executes it.
func RunFromFile(fileName string, input []int) ([]int, []int) {
	ss := util.MustLoadStringSlice(fileName, ",")
	program := []int{}
	for _, s := range ss {
		i := util.MustAtoi(s)
		program = append(program, i)
	}
	return Run(program, input)
}

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
			add(numParams, vm)
		case opMultiply:
			multiply(numParams, vm)
		case opInput:
			in(numParams, vm)
		case opOutput:
			out(numParams, vm)
		case opLessThan:
			lessThan(numParams, vm)
		case opEquals:
			equals(numParams, vm)
		case opHalt:
			// Opcode 99 means that the program is finished and should immediately halt.
			// The instruction 99 contains only an opcode and has no parameters.
			return vm.memory, vm.output
		default:
			// Encountering an unknown opcode means something went wrong.
			log.Fatalf("Unexpected op code: %d", opCode)
		}
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
	// Leftmost chars are the parameter modes, if any
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
	modes = strings.Repeat("0", lenPrefix) + modes
	vm.parameterModes = stack.New()
	for _, r := range modes {
		vm.parameterModes.Push(int(r - '0')) // '0' => 0, e.g.
	}
}

// add (Opcode 1) - adds together numbers read from two positions and stores
// the result in a third position. The three integers immediately after the
// opcode tell you these three positions - the first two indicate the positions
// from which you should read the input values, and the third indicates the
// position at which the output should be stored.
func add(numParams int, vm *computer) {
	op1, op2 := next2Params(vm)
	store(op1+op2, vm.memory[vm.iP+3], vm)
	advanceInstructionPointer(numParams+1, vm)
}

// multiply (Opcode 2) - works exactly like opcode 1, except it multiplies
// the two inputs instead of adding them. Again, the three integers after
// the opcode indicate where the inputs and outputs are, not their values.
func multiply(numParams int, vm *computer) {
	op1, op2 := next2Params(vm)
	store(op1*op2, vm.memory[vm.iP+3], vm)
	advanceInstructionPointer(numParams+1, vm)
}

// in (Opcode 3) - takes a single integer as input and saves it to the position given
// by its only parameter. For example, the instruction 3,50 would take an input
// value and store it at address 50.
func in(numParams int, vm *computer) {
	i := vm.input.Pop()
	store(i.(int), vm.memory[vm.iP+1], vm)
	advanceInstructionPointer(numParams+1, vm)
}

// out (Opcode 4) - outputs the value of its only parameter. For example,
// the instruction 4,50 would output the value at address 50.
func out(numParams int, vm *computer) {
	i := fetch(vm.iP+1, vm)
	vm.output = append(vm.output, i)
	advanceInstructionPointer(numParams+1, vm)
}

// jumpIfTrue (Opcode 5) - if the first parameter is non-zero, it sets the instruction
// pointer to the value from the second parameter. Otherwise, it does nothing.

// jumpIfFalse (Opcode 6) - if the first parameter is zero, it sets the instruction
// pointer to the value from the second parameter. Otherwise, it does nothing.

// lessThan (Opcode 7) - if the first parameter is less than the second parameter, it
// stores 1 in the position given by the third parameter. Otherwise, it stores 0.
func lessThan(numParams int, vm *computer) {
	op1, op2 := next2Params(vm)
	var result int
	if op1 < op2 {
		result = 1
	} else {
		result = 0
	}
	store(result, vm.memory[vm.iP+3], vm)
	advanceInstructionPointer(numParams+1, vm)
}

// equals (Opcode 8) - if the first parameter is equal to the second parameter, it
// stores 1 in the position given by the third parameter. Otherwise, it stores 0.
func equals(numParams int, vm *computer) {
	op1, op2 := next2Params(vm)
	var result int
	if op1 == op2 {
		result = 1
	} else {
		result = 0
	}
	store(result, vm.memory[vm.iP+3], vm)
	advanceInstructionPointer(numParams+1, vm)
}

func next2Params(vm *computer) (int, int) {
	return fetch(vm.iP+1, vm), fetch(vm.iP+2, vm)
}

// fetch returns the value at the address, taking into consideration
// the parameter mode.
//
// Immediate mode interprets the `address` param as the address of the
// value to return. If the `address` param is 50, we return the value
// stored at address 50.
//
// Position mode adds a level of indirection. It interprets the `address`
// param as the address of an address. The second address is the address of
// the value to return. If the `address` param is 50, we get the value stored
// at address 50. Suppose that value is 100. We then get the value stored at
// address 100, and return it.
func fetch(address int, vm *computer) int {
	var i int
	switch mode := vm.parameterModes.Pop(); mode {
	case immediateMode:
		i = vm.memory[address]
	case positionMode:
		i = vm.memory[address]
		i = vm.memory[i]
	default:
		log.Fatalf("Unexpected parameter mode: %d", mode)
	}
	return i
}

// "Parameters that an instruction writes to will never be in immediate mode."
func store(value int, address int, vm *computer) {
	vm.memory[address] = value
}

func advanceInstructionPointer(i int, vm *computer) {
	vm.iP += i
}

// load creates the vm and loads the program into it.
func load(program []int, input []int) *computer {
	vm := new(computer)
	vm.memory = make([]int, len(program))
	copy(vm.memory, program)
	vm.iP = 0
	vm.input = stack.New()
	// Start at end so we can use a stack and simply pop as needed
	for i := len(input) - 1; i >= 0; i-- {
		vm.input.Push(input[i])
	}
	vm.output = []int{}
	return vm
}
