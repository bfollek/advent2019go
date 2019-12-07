// Package sequence is for day04
package sequence

// Sequence struct manages sequences of the same digit.
type Sequence struct {
	digits []byte // Sequence of the same digit
	Found  bool   // Have we found a sequence of the same digit two or more times?
	Found2 bool   // Have we found a sequence of the same digit exactly two times?
}

// Len returns the length of a sequence.
func (seq *Sequence) Len() int {
	return len(seq.digits)
}

// Add adds a digit to a sequence.
func (seq *Sequence) Add(b byte) {
	seq.digits = append(seq.digits, b)
}

// Ended handles the end of a sequence.
func (seq *Sequence) Ended() {
	switch seq.Len() {
	case 0, 1:
		break // No effect
	case 2:
		seq.Found2 = true
		seq.Found = true
	default: // Len > 2
		seq.Found = true
	}
}

// Last returns the last byte of the sequence or 0 if the sequence is empty.
func (seq *Sequence) Last() byte {
	if seq.Len() > 0 {
		return seq.digits[seq.Len()-1]
	}
	return 0
}

// Reset resets the digits slice in the sequence. It DOES NOT reset the
// `found` and `found2` flags. We want to preserve their state across multiple
// digit sequences so that the client will know if an earlier sequence set them
//  to true.
func (seq *Sequence) Reset(b byte) {
	seq.digits = []byte{b}
}
