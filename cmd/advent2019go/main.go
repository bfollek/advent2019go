package main

import (
	"fmt"

	"github.com/bfollek/advent2019go/day01"
	"github.com/bfollek/advent2019go/day02"
)

func main() {
	fmt.Printf("day01.Part1: %d\n", day01.Part1("../../day01/testdata/day01.dat"))
	fmt.Printf("day01.Part2: %d\n", day01.Part2("../../day01/testdata/day01.dat"))

	fmt.Printf("day02.Part1: %d\n", day02.Part1("../../day02/testdata/day02.dat"))
	if i, err := day02.Part2("../../day02/testdata/day02.dat"); err != nil {
		fmt.Printf("day02.Part2 error: %s\n", err)
	} else {
		fmt.Printf("day02.Part2: %d\n", i)
	}
}
