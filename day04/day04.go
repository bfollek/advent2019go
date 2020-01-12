package day04

import (
	"strconv"

	"github.com/bfollek/aoc19go/sequence"
	"github.com/bfollek/aoc19go/util"
)

const passwordLen = 6

// Part1 "How many different passwords within the range
// given in your puzzle input meet these criteria?"
func Part1(fileName string) int {
	return countValid(fileName, false)
}

// Part2 "How many different passwords within the range
// given in your puzzle input meet all of the criteria?"
func Part2(fileName string) int {
	return countValid(fileName, true)
}

func countValid(fileName string, mustHaveSeq2 bool) int {
	numValid := 0
	rng := loadRange(fileName)
	start := rng[0]
	end := rng[1]
	for i := start; i <= end; i++ {
		s := strconv.Itoa(i)
		if isValid(s, mustHaveSeq2) {
			numValid++
		}
	}
	return numValid
}

// isValid returns true if the password passes criteria, else false.
// The passwords are digits, so we can safely work with bytes instead of runes.
func isValid(password string, mustHaveSeq2 bool) bool {
	if len(password) != passwordLen {
		return false
	}
	seq := new(sequence.Sequence)
	for i := 0; i < passwordLen; i++ {
		current := password[i]
		if j := i + 1; j < passwordLen && current > password[j] {
			return false // Decreasing digits are invalid
		}
		if current == seq.Last() {
			seq.Add(current)
			continue
		}
		seq.Ended()
		seq.Reset(current)
	}
	seq.Ended() // Last digit ends a sequence
	if mustHaveSeq2 {
		return seq.Found2
	}
	return seq.Found
}

func loadRange(fileName string) []int {
	rng := []int{}
	ss := util.MustLoadStringSlice(fileName, "-")
	for _, s := range ss {
		rng = append(rng, util.MustAtoi(s))
	}
	return rng
}
