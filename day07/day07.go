package day07

import (
	"math"

	ic "github.com/bfollek/aoc19go/intcode"
	"github.com/gitchander/permutation"
)

// Part1 tries every combination of phase settings on the amplifiers.
// What is the highest signal that can be sent to the thrusters?
func Part1(fileName string) int {
	program := ic.LoadFromFile(fileName)
	maxSoFar := math.MinInt32
	combos := phaseSettings([]int{0, 1, 2, 3, 4})
	for _, combo := range combos {
		opSig := outputSignal(combo, program)
		if opSig > maxSoFar {
			maxSoFar = opSig
		}
	}
	return maxSoFar
}

// Part2: Try every combination of the new phase settings
// on the amplifier feedback loop. What is the highest
// signal that can be sent to the thrusters?
// func Part2(fileName string) int {
// 	program := intcode.LoadFromFile(fileName)
// 	maxSoFar := math.MinInt32
// 	combos := phaseSettings([]int{5, 6, 7, 8, 9})
// 	for _, combo := range combos {
// 		opSig := outputSignalLoop(combo, program)
// 		if opSig > maxSoFar {
// 			maxSoFar = opSig
// 		}
// 	}
// 	return maxSoFar
// }

func phaseSettings(sl []int) [][]int {
	combos := [][]int{}
	p := permutation.New(permutation.IntSlice(sl))
	for p.Next() {
		// If I don't make a copy, all the []int slices in combos
		// have the same value - whatever the last combo generated is.
		// I think this is because permutation reuses the same slice -
		// it permutes in place.
		tmp := make([]int, len(sl))
		copy(tmp, sl)
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

// func outputSignalLoop(combo []int, program []int) int {
// 	vms := []intcode.VM{}
// 	var wg sync.WaitGroup
// 	for _, phaseSetting := range combo {
// 		vm := intcode.New()
// 		vms = append(vms, vm)
// 		vm.In <- phaseSetting
// 		wg.Add(1)
// 		go vm.Run(program)
// 	}
// }

// var wg sync.WaitGroup

// 	for i := 0; i < 5; i++ {
// 		fmt.Println("Main: Starting worker", i)
// 		wg.Add(1)
