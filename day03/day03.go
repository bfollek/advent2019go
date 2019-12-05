package day03

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/bfollek/advent2019go/util"
	mapset "github.com/deckarep/golang-set"
)

const UP = 'U'
const DOWN = 'D'
const RIGHT = 'R'
const LEFT = 'L'

type move struct {
	direction byte
	distance  int
}

type point struct {
	x int64
	y int64
}

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

// |x1 - x2| + |y1 - y2|
// https://xlinux.nist.gov/dads/HTML/manhattanDistance.html
func manhattanDistance(p point) int64 {
	// The second point is always centralPort
	return util.AbsInt64(p.x-centralPort.x) + util.AbsInt64(p.y-centralPort.y)
}

func getPath(moves []move) mapset.Set {
	path := mapset.NewSet()
	currentX := int64(0)
	currentY := int64(0)
	for _, m := range moves {
		for ; m.distance > 0; m.distance-- {
			switch m.direction {
			case UP:
				currentY++
			case DOWN:
				currentY--
			case RIGHT:
				currentX++
			case LEFT:
				currentX--
			}
			p := point{currentX, currentY}
			path.Add(p)
		}
	}
	return path
}

func loadMoves(fileName string) ([]move, []move) {
	lines := util.MustReadLines(fileName)
	first := lineToMoves(lines[0])
	second := lineToMoves(lines[1])
	return first, second
}

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
