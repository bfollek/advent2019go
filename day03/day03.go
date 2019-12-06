package day03

import (
	"math"
	"strings"

	"github.com/bfollek/advent2019go/util"
)

const up = 'U'
const down = 'D'
const right = 'R'
const left = 'L'

type closeFunc func(point, map[point]int64, map[point]int64) int64

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
	return calcClosest(fileName, closestByManhattanDistance)
}

// Part2 "What is the fewest combined steps the wires
// must take to reach an intersection?"
func Part2(fileName string) int64 {
	return calcClosest(fileName, closestBySteps)
}

func calcClosest(fileName string, f closeFunc) int64 {
	wire1Moves, wire2Moves := loadMoves(fileName)
	wire1Path := getPath(wire1Moves)
	wire2Path := getPath(wire2Moves)
	crossPoints := intersection(wire1Path, wire2Path)
	closest := int64(math.MaxInt64)
	for _, p := range crossPoints {
		if howClose := f(p, wire1Path, wire2Path); howClose < closest {
			closest = howClose
		}
	}
	return closest
}

// manhattanDistance formula: |x1 - x2| + |y1 - y2|
// https://xlinux.nist.gov/dads/HTML/manhattanDistance.html
func closestByManhattanDistance(p point, _, _ map[point]int64) int64 {
	// The second point is always centralPort
	return util.AbsInt64(p.x-centralPort.x) + util.AbsInt64(p.y-centralPort.y)
}

func closestBySteps(p point, wire1Path, wire2Path map[point]int64) int64 {
	return wire1Path[p] + wire2Path[p]
}

// getPath creates a map of the wire's path.
// Each key is a  point the wire moves to.
// Each value is the number of steps the wire takes to get there.
func getPath(moves []move) map[point]int64 {
	path := map[point]int64{}
	steps := int64(0)
	currentX := int64(0)
	currentY := int64(0)
	for _, m := range moves {
		for ; m.distance > 0; m.distance-- {
			steps++
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
			// The first time we get to a point, save the number of steps it took to get here.
			if _, ok := path[p]; !ok {
				path[p] = steps
			}
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
		d := util.MustAtoi(s[1:])
		m.distance = d
		moves = append(moves, m)
	}
	return moves
}

func intersection(m1, m2 map[point]int64) []point {
	inter := []point{}
	// Drive the loop off the shorter map
	var short, long map[point]int64
	if len(m1) < len(m2) {
		short = m1
		long = m2
	} else {
		short = m2
		long = m1
	}
	for key := range short {
		if _, ok := long[key]; ok {
			inter = append(inter, key)
		}
	}
	return inter
}
