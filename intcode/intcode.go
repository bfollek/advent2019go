package intcode

import (
	"log"

	"github.com/bfollek/aoc19go/util"
	"github.com/golang-collections/collections/stack"
)

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

type opCodeAttribs struct {
	numParams int
	exec      func(opCodeAttribs, *VM)
}

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

var opCodes = map[int]opCodeAttribs{
	opAdd:         {3, add},
	opMultiply:    {3, multiply},
	opInput:       {1, in},
	opOutput:      {1, out},
	opJumpIfTrue:  {2, jumpIfTrue},
	opJumpIfFalse: {2, jumpIfFalse},
	opLessThan:    {3, lessThan},
	opEquals:      {3, equals},
}

const bufSize = 1000

// VM is the intcode virtual machine.
type VM struct {
	memory         []int
	iP             int // Instruction pointer
	parameterModes *stack.Stack
	In             chan int
	Out            chan int
	Mem            chan int // At end of program, return memory contents on this channel
}

// New returns an initialized VM.
func New() *VM {
	vm := new(VM)
	vm.In = make(chan int, bufSize)
	vm.Out = make(chan int, bufSize)
	vm.Mem = make(chan int, bufSize)
	return vm
}

// LoadFromFile loads an intcode program from a file.
func LoadFromFile(fileName string) []int {
	ss := util.MustLoadStringSlice(fileName, ",")
	program := []int{}
	for _, s := range ss {
		i := util.MustAtoi(s)
		program = append(program, i)
	}
	return program
}

// RunFromFile reads an intcode program from a file, then executes it.
func (vm *VM) RunFromFile(fileName string) {
	program := LoadFromFile(fileName)
	vm.Run(program)
}

// Run executes an intcode program.
// The `program` param is the program code.
//
// As the program runs, it reads input values from the In channel,
// and it writes output values to the Out channel. If the program
// terminates normally, it writes the contents of memory to the Mem
// channel.
//
// The program may modify the memory it's in as it runs.
// This means that the program may be self-modifying.
func (vm *VM) Run(program []int) {
	load(program, vm)
	for {
		// Parameter modes are stored in the same value as the instruction's opcode.
		// The opcode is a two-digit number based only on the ones and tens digit
		// of the value, that is, the opcode is the rightmost two digits of the
		// first value in an instruction.
		rawOpCode := vm.memory[vm.iP]
		setParameterModes(rawOpCode, vm)
		opCode := rawOpCode % 100
		// Opcode 99 means that the program is finished and should immediately halt.
		// The instruction 99 contains only an opcode and has no parameters.
		if opCode == opHalt {
			for _, m := range vm.memory {
				vm.Mem <- m
			}
			return
		}
		oca, ok := opCodes[opCode]
		if !ok {
			// Encountering an unknown opcode means something went wrong.
			log.Fatalf("Unexpected op code: %d", opCode)
		}
		oca.exec(oca, vm)
	}
}

// Parameter modes are single digits, one per parameter, read right-to-left
// from the opcode: the first parameter's mode is in the hundreds digit,
// the second parameter's mode is in the thousands digit, the third parameter's
// mode is in the ten-thousands digit, and so on. Any missing modes are 0.
func setParameterModes(remaining int, vm *VM) {
	// No harm setting more modes than the actual number of parameters.
	// Anything extra just won't get popped.
	var mode int
	for _, n := range []int{10000, 1000, 100} {
		if remaining > n {
			remaining -= n
			mode = 1
		} else {
			mode = 0
		}
		vm.parameterModes.Push(mode)
	}
}

// add (Opcode 1) - adds together numbers read from two positions and stores
// the result in a third position. The three integers immediately after the
// opcode tell you these three positions - the first two indicate the positions
// from which you should read the input values, and the third indicates the
// position at which the output should be stored.
func add(oca opCodeAttribs, vm *VM) {
	op1, op2 := next2Params(vm)
	store(op1+op2, vm.memory[vm.iP+3], vm)
	advanceInstructionPointer(oca.numParams+1, vm)
}

// multiply (Opcode 2) - works exactly like opcode 1, except it multiplies
// the two inputs instead of adding them. Again, the three integers after
// the opcode indicate where the inputs and outputs are, not their values.
func multiply(oca opCodeAttribs, vm *VM) {
	op1, op2 := next2Params(vm)
	store(op1*op2, vm.memory[vm.iP+3], vm)
	advanceInstructionPointer(oca.numParams+1, vm)
}

// in (Opcode 3) - takes a single integer as input and saves it to the position given
// by its only parameter. For example, the instruction 3,50 would take an input
// value and store it at address 50.
func in(oca opCodeAttribs, vm *VM) {
	i := <-vm.In
	store(i, vm.memory[vm.iP+1], vm)
	advanceInstructionPointer(oca.numParams+1, vm)
}

// out (Opcode 4) - outputs the value of its only parameter. For example,
// the instruction 4,50 would output the value at address 50.
func out(oca opCodeAttribs, vm *VM) {
	i := fetch(vm.iP+1, vm)
	vm.Out <- i
	advanceInstructionPointer(oca.numParams+1, vm)
}

// jumpIfTrue (opCode 5) - if the first parameter is non-zero, it sets the instruction
// pointer to the value from the second parameter. Otherwise, it does nothing.
func jumpIfTrue(oca opCodeAttribs, vm *VM) {
	p1, p2 := next2Params(vm)
	jump(p1 != 0, p2, oca, vm)
}

// jumpIfFalse (Opcode 6) - if the first parameter is zero, it sets the instruction
// pointer to the value from the second parameter. Otherwise, it does nothing.
func jumpIfFalse(oca opCodeAttribs, vm *VM) {
	p1, p2 := next2Params(vm)
	jump(p1 == 0, p2, oca, vm)
}

func jump(jump bool, jumpTo int, oca opCodeAttribs, vm *VM) {
	if jump {
		setInstructionPointer(jumpTo, vm)
	} else {
		advanceInstructionPointer(oca.numParams+1, vm)
	}
}

// lessThan (Opcode 7) - if the first parameter is less than the second parameter, it
// stores 1 in the position given by the third parameter. Otherwise, it stores 0.
func lessThan(oca opCodeAttribs, vm *VM) {
	op1, op2 := next2Params(vm)
	comparison(op1 < op2, oca, vm)
}

// equals (Opcode 8) - if the first parameter is equal to the second parameter, it
// stores 1 in the position given by the third parameter. Otherwise, it stores 0.
func equals(oca opCodeAttribs, vm *VM) {
	op1, op2 := next2Params(vm)
	comparison(op1 == op2, oca, vm)
}

func comparison(satisfied bool, oca opCodeAttribs, vm *VM) {
	var i int
	if satisfied {
		i = 1
	} else {
		i = 0
	}
	store(i, vm.memory[vm.iP+3], vm)
	advanceInstructionPointer(oca.numParams+1, vm)
}

func next2Params(vm *VM) (int, int) {
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
func fetch(address int, vm *VM) int {
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
func store(value int, address int, vm *VM) {
	// Make sure there's room. If there isn't, add a chunk.
	if address >= len(vm.memory) {
		more := address * 10
		vm.memory = append(vm.memory, make([]int, more)...)
	}
	vm.memory[address] = value
}

func advanceInstructionPointer(i int, vm *VM) {
	vm.iP += i
}

func setInstructionPointer(i int, vm *VM) {
	vm.iP = i
}

func load(program []int, vm *VM) {
	vm.memory = make([]int, len(program))
	// Copy so that we don't overwrite the program.
	copy(vm.memory, program)
	vm.iP = 0
	vm.parameterModes = stack.New()
}
