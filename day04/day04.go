package day04

import (
	"fmt"
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
		if isValid(s, false) {
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

func isValid(password string, mustHaveSeq2 bool) bool {
	if len(password) != passwordLen {
		return false
	}
	seq := []byte{}
	foundSeq := false
	foundSeq2 := false
	// The passwords are digits, so we can safely work with bytes instead of runes.
	for i := 0; i < passwordLen; i++ {
		lSeq := len(seq)
		current := password[i]
		if lSeq > 0 && current == seq[lSeq-1] {
			seq = append(seq, current) // Add to sequence
			continue
		}
		// If we get here, sequence ended
		j := i + 1
		if j < passwordLen && current > password[j] {
			return false // Decreasing pair
		}
		switch {
		case lSeq == 0:
			break
		case lSeq == 2:
			foundSeq2 = true
		case lSeq > 1:
			foundSeq = true
		}
		fmt.Printf("inside loop %s %d %t %t\n", seq, lSeq, foundSeq, foundSeq2)
		seq = []byte{current} // Start new sequence
	}
	lSeq := len(seq)
	fmt.Printf("outside loop %s %d %t %t\n", seq, lSeq, foundSeq, foundSeq2)
	switch {
	case lSeq == 2:
		foundSeq2 = true
	case lSeq > 1:
		foundSeq = true
	}
	if mustHaveSeq2 {
		return foundSeq2
	}
	return foundSeq || foundSeq2
}

func isValidWithSeqOf2(password string) bool {
	if !isValid(password, false) {
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
