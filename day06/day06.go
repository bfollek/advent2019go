package day06

import (
	"strings"

	"github.com/bfollek/advent2019go/util"
)

const centerOfMass = "COM"

// Part1 "What is the total number of direct and indirect orbits in your map data?"
func Part1(fileName string) int {
	orbits := loadOrbits(fileName)
	total := 0
	for k := range orbits {
		total += countOrbits(k, orbits)
	}
	return total
}

func countOrbits(obj string, orbits map[string]string) int {
	whatObjOrbits := orbits[obj]
	if whatObjOrbits == centerOfMass {
		return 1
	}
	return 1 + countOrbits(whatObjOrbits, orbits)
}

func loadOrbits(fileName string) map[string]string {
	lines := util.MustReadLines(fileName)
	m := map[string]string{}
	// obj1 is orbited by obj2. Build the map the other way around:
	// obj2 orbits obj1. obj2 is the key, obj1 is the value.
	for _, line := range lines {
		sl := strings.Split(line, ")")
		obj1, obj2 := sl[0], sl[1]
		m[obj2] = obj1
	}
	return m
}
