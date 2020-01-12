package main

import (
	"fmt"

	"github.com/bfollek/advent2019go/day01"
	"github.com/bfollek/advent2019go/day02"
	"github.com/bfollek/advent2019go/day03"
	"github.com/bfollek/advent2019go/day04"
	"github.com/bfollek/advent2019go/day05"
	"github.com/bfollek/advent2019go/day06"
	"github.com/bfollek/advent2019go/day07"
)

func main() {
	// Anything I want to hit first in delve goes here...
	// fmt.Printf("day05.Part1: %d\n",
	// 	day05.Part1("../../day05/testdata/immediate_add_v2.dat"))

	fmt.Printf("day01.Part1: %d\n", day01.Part1("../../day01/testdata/day01.dat"))
	fmt.Printf("day01.Part2: %d\n", day01.Part2("../../day01/testdata/day01.dat"))

	fmt.Printf("day02.Part1: %d\n", day02.Part1("../../day02/testdata/day02.dat"))
	if i, err := day02.Part2("../../day02/testdata/day02.dat"); err != nil {
		fmt.Printf("day02.Part2 error: %s\n", err)
	} else {
		fmt.Printf("day02.Part2: %d\n", i)
	}

	fmt.Printf("day03.Part1: %d\n", day03.Part1("../../day03/testdata/day03.dat"))
	fmt.Printf("day03.Part2: %d\n", day03.Part2("../../day03/testdata/day03.dat"))

	fmt.Printf("day04.Part1: %d\n", day04.Part1("../../day04/testdata/day04.dat"))
	fmt.Printf("day04.Part2: %d\n", day04.Part2("../../day04/testdata/day04.dat"))

	fmt.Printf("day05.Part1: %d\n", day05.Part1("../../day05/testdata/day05.dat"))
	fmt.Printf("day05.Part2: %d\n", day05.Part2("../../day05/testdata/day05.dat"))

	fmt.Printf("day06.Part1: %d\n", day06.Part1("../../day06/testdata/day06_part1_short.dat"))
	fmt.Printf("day06.Part1: %d\n", day06.Part1("../../day06/testdata/day06.dat"))
	fmt.Printf("day06.Part2: %d\n", day06.Part2("../../day06/testdata/day06_part2_short.dat"))
	fmt.Printf("day06.Part2: %d\n", day06.Part2("../../day06/testdata/day06.dat"))

	fmt.Printf("day07.Part1: %d\n", day07.Part1("../../day07/testdata/day07.dat"))
}
