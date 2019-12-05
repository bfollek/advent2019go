package day03

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/bfollek/advent2019go/util"
	mapset "github.com/deckarep/golang-set"
)

const up = 'U'
const down = 'D'
const right = 'R'
const left = 'L'

type move struct {
	direction byte
	distance  int
}

type point struct {
	x int64
	y int64
}

// centralPort is de facto a constant.
var centralPort = point{0, 0}

// Part1 "What is the Manhattan distance from the central port
// to the closest intersection?"
func Part1(fileName string) int64 {
	wire1Moves, wire2Moves := loadMoves(fileName)
	wire1Path := getPath(wire1Moves)
	wire2Path := getPath(wire2Moves)
	crossPoints := wire1Path.Intersect(wire2Path)

	closest := int64(math.MaxInt64)
	for _, p := range crossPoints.ToSlice() {
		if md := manhattanDistance(p.(point)); md < closest {
			closest = md
		}
	}
	return closest
}

// manhattanDistance formula: |x1 - x2| + |y1 - y2|
// https://xlinux.nist.gov/dads/HTML/manhattanDistance.html
func manhattanDistance(p point) int64 {
	// The second point is always centralPort
	return util.AbsInt64(p.x-centralPort.x) + util.AbsInt64(p.y-centralPort.y)
}

// getPath creates a set of the points the wire moves across.
// This is the wire's path.
func getPath(moves []move) mapset.Set {
	path := mapset.NewSet()
	currentX := int64(0)
	currentY := int64(0)
	for _, m := range moves {
		for ; m.distance > 0; m.distance-- {
			switch m.direction {
			case up:
				currentY++
			case down:
				currentY--
			case right:
				currentX++
			case left:
				currentX--
			}
			p := point{currentX, currentY}
			path.Add(p)
		}
	}
	return path
}

// loadMoves uses just the first two lines in the file, one for each wire.
func loadMoves(fileName string) ([]move, []move) {
	lines := util.MustReadLines(fileName)
	first := lineToMoves(lines[0])
	second := lineToMoves(lines[1])
	return first, second
}

// lineToMoves turns a line like "R8,U5,L5,D3"
// into a slice of move structs.
func lineToMoves(line string) []move {
	moves := []move{}
	ss := strings.Split(line, ",")
	for _, s := range ss {
		m := move{}
		m.direction = s[0]
		d, err := strconv.Atoi(s[1:])
		if err != nil {
			log.Fatal(err)
		}
		m.distance = d
		moves = append(moves, m)
	}
	return moves
}
