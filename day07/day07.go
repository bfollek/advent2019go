package day07

import (
	"math"

	"github.com/bfollek/aoc19go/intcode"
	"github.com/gitchander/permutation"
)

// Part1 tries every combination of phase settings on the amplifiers.
// What is the highest signal that can be sent to the thrusters?
func Part1(fileName string) int {
	program := intcode.LoadFromFile(fileName)
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
		_, output := intcode.Run(program, []int{phaseSetting, opSig})
		opSig = output[0]
	}
	return opSig
}
