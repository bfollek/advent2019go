package day07

import (
	"math"
	"sync"

	"github.com/bfollek/aoc19go/intcode"
	ic "github.com/bfollek/aoc19go/intcode"
	"github.com/gitchander/permutation"
)

type opSigFunc func([]int, []int) int

// Part1 tries every combination of phase settings on the amplifiers.
// What is the highest signal that can be sent to the thrusters?
func Part1(fileName string) int {
	return calcOpSig(fileName, []int{0, 1, 2, 3, 4}, outputSignal)
}

// Part2: Try every combination of the new phase settings
// on the amplifier feedback loop. What is the highest
// signal that can be sent to the thrusters?
func Part2(fileName string) int {
	return calcOpSig(fileName, []int{5, 6, 7, 8, 9}, loopedOutputSignal)
}

func calcOpSig(fileName string, phases []int, f opSigFunc) int {
	program := intcode.LoadFromFile(fileName)
	maxSoFar := math.MinInt32
	combos := phaseSettings(phases)
	for _, combo := range combos {
		opSig := f(combo, program)
		if opSig > maxSoFar {
			maxSoFar = opSig
		}
	}
	return maxSoFar
}

func phaseSettings(phases []int) [][]int {
	combos := [][]int{}
	p := permutation.New(permutation.IntSlice(phases))
	for p.Next() {
		// If I don't make a copy, all the []int slices in combos
		// have the same value - whatever the last combo generated is.
		// I think this is because permutation reuses the same slice -
		// it permutes in place.
		tmp := make([]int, len(phases))
		copy(tmp, phases)
		combos = append(combos, tmp)
	}
	return combos
}

func outputSignal(combo []int, program []int) int {
	opSig := 0
	for _, phaseSetting := range combo {
		vm := ic.New(ic.MakeAllChannels())
		go vm.Run(program)
		vm.In <- phaseSetting
		vm.In <- opSig
		opSig = <-vm.Out
	}
	return opSig
}

func loopedOutputSignal(combo []int, program []int) int {
	numVms := len(combo)
	vms := wireUpLoop(numVms)
	for i, phaseSetting := range combo {
		vm := vms[i]
		vm.In <- phaseSetting
		// "To start the process, a 0 signal is sent to amplifier A's input exactly once."
		if i == 0 {
			vm.In <- 0
		}
	}
	// var wg sync.WaitGroup
	// for _, vm := range vms {
	// 	wg.Add(1)
	// 	go vm.RunInWaitGroup(program, &wg)
	// }
	// wg.Wait()
	runVms(program, vms)
	return <-vms[numVms-1].Out
}

func wireUpLoop(numVms int) []*ic.VM {
	vms := []*ic.VM{}
	var vm *ic.VM
	in, out, mem := ic.MakeAllChannels()
	for i := 0; i < numVms; i++ {
		vm = intcode.New(in, out, mem)
		vms = append(vms, vm)
		in = vm.Out
		out = ic.MakeChannel()
		mem = ic.MakeChannel()
	}
	// The first vm gets input from the last vm's output
	vms[0].In = vms[numVms-1].Out
	return vms
}

func runVms(program []int, vms []*ic.VM) {
	var wg sync.WaitGroup
	for _, vm := range vms {
		wg.Add(1)
		go vm.RunInWaitGroup(program, &wg)
	}
	wg.Wait()
}
