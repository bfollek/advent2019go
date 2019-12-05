package day03

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/bfollek/advent2019go/util"
	mapset "github.com/deckarep/golang-set"
)

type point struct {
	x int64
	y int64
}

// Part1 "What is the Manhattan distance from the central port
// to the closest intersection?"
func Part1(fileName string) int64 {
	wire1Moves, wire2Moves := loadMoves(fileName)
	wire1Path := getPath(wire1Moves)
	wire2Path := getPath(wire2Moves)
	crossPoints := wire1Path.Intersect(wire2Path)

	closest := int64(math.MaxInt64)
	for _, p := range crossPoints.ToSlice() {
		if md := manhattanDistanceFromCentralPort(p.(point)); md < closest {
			closest = md
		}
	}
	return closest
}

// |x1 - x2| + |y1 - y2|
// https://xlinux.nist.gov/dads/HTML/manhattanDistance.html
func manhattanDistanceFromCentralPort(p point) int64 {
	// The second point is always the central port, {0 0}
	return util.AbsInt64(p.x-0) + util.AbsInt64(p.y-0)
}

func getPath(moves []string) mapset.Set {
	path := mapset.NewSet()
	currentX := int64(0)
	currentY := int64(0)
	for _, move := range moves {
		runes := []rune(move)
		direction := runes[0]
		distance, err := strconv.Atoi(string(runes[1:]))
		if err != nil {
			log.Fatal(err)
		}
		for ; distance > 0; distance-- {
			switch direction {
			case 'U':
				currentY++
			case 'D':
				currentY--
			case 'R':
				currentX++
			case 'L':
				currentX--
			}
			p := point{currentX, currentY}
			path.Add(p)
		}
	}
	return path
}

func loadMoves(fileName string) ([]string, []string) {
	lines := util.MustReadLines(fileName)
	return strings.Split(lines[0], ","), strings.Split(lines[1], ",")
}
