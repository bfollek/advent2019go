package day08

type imgRow []int

type layer []imgRow

type img struct {
	width  int
	height int
	layers []layer
}

const width = 25
const height = 6

// Part1 - find the layer that contains the fewest 0 digits. On that layer,
// what is the number of 1 digits multiplied by the number of 2 digits?
func Part1(fileName string) int {
	return 0
}

func loadImg(fileName string) {
}
