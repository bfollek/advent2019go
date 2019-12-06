package day04

import (
	"strconv"

	"github.com/bfollek/advent2019go/util"
)

const passwordLen = 6

// Part1 "How many different passwords within the range
// given in your puzzle input meet these criteria?"
func Part1(fileName string) int {
	numValid := 0
	rng := loadRange(fileName)
	start := rng[0]
	end := rng[1]
	for i := start; i <= end; i++ {
		s := strconv.Itoa(i)
		if isValid(s) {
			numValid++
		}
	}
	return numValid
}

// Part2 "How many different passwords within the range
// given in your puzzle input meet all of the criteria?"
func Part2(fileName string) int {
	numValid := 0
	rng := loadRange(fileName)
	start := rng[0]
	end := rng[1]
	for i := start; i <= end; i++ {
		s := strconv.Itoa(i)
		if isValidWithSeqOf2(s) {
			numValid++
		}
	}
	return numValid
}

func isValid(password string) bool {
	if len(password) != passwordLen {
		return false
	}
	foundSeq := false
	// These passwords are digits, so we can safely work with bytes instead of runes.
	for i := 0; i < passwordLen; i++ {
		j := i + 1
		if j == passwordLen {
			break // Nothing left to test
		}
		nxt := password[i]
		nxtNxt := password[j]
		switch {
		case nxt > nxtNxt:
			return false // Decreasing pair
		case nxt == nxtNxt:
			foundSeq = true
		}
	}
	return foundSeq
}

func isValidWithSeqOf2(password string) bool {
	if !isValid(password) {
		return false
	}
	buf := []byte{password[0]}
	for i := 1; i < passwordLen; i++ {
		nxt := password[i]
		lb := len(buf)
		if nxt != buf[lb-1] { // Sequence ended
			if lb == 2 {
				return true // Found sequence of 2
			}
			buf = []byte{nxt} // Start new sequence
		} else {
			buf = append(buf, nxt) // Add to current sequence
		}
	}
	return len(buf) == 2 // Catch case where last 2 chars are a seq
}

func loadRange(fileName string) []int {
	rng := []int{}
	ss := util.MustLoadStringSlice(fileName, "-")
	for _, s := range ss {
		rng = append(rng, util.MustAtoi(s))
	}
	return rng
}
