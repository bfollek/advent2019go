package day04

import (
	"strconv"

	"github.com/bfollek/advent2019go/util"
)

const passwordLen = 6

// type seqState struct {
// 	seqDigits []byte // Sequence of the same digit
// 	foundSeq  bool   // Have we found a sequence of the same digit two or more times?
// 	foundSeq2 bool   // Have we found a sequence of the same digit exactly two times?
// }

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
	seq := new(sequence)
	for i := 0; i < passwordLen; i++ {
		current := password[i]
		if j := i + 1; j < passwordLen && current > password[j] {
			return false // Decreasing digits are invalid
		}
		if lSeq := len(seq.digits); lSeq > 0 && current == seq.digits[lSeq-1] {
			seq.add(current)
			continue
		}
		sequenceEnded(seq)
		seq.digits = []byte{current} // Start a new sequence
	}
	sequenceEnded(seq) // Last digit ends a sequence
	if mustHaveSeq2 {
		return seq.found2
	}
	return seq.found
}

func sequenceEnded(seq *sequence) { //seq []byte, pFoundSeq *bool, pFoundSeq2 *bool) {
	lSeq := len(seq.digits)
	switch {
	case lSeq == 2:
		seq.found2 = true
		seq.found = true
	case lSeq > 1:
		seq.found = true
	}
}

func loadRange(fileName string) []int {
	rng := []int{}
	ss := util.MustLoadStringSlice(fileName, "-")
	for _, s := range ss {
		rng = append(rng, util.MustAtoi(s))
	}
	return rng
}
