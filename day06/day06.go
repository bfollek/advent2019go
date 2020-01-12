package day06

import (
	"strings"

	"github.com/bfollek/aoc19go/util"
	"github.com/yourbasic/graph"
)

const centerOfMass = "COM"
const santa = "SAN"
const you = "YOU"

// Part1 "What is the total number of direct and indirect orbits in your map data?"
func Part1(fileName string) int {
	orbits := loadOrbits(fileName)
	memoCache := map[string]int{}
	total := 0
	for k := range orbits {
		total += countOrbits(k, orbits, memoCache)
	}
	return total
}

// Part2 "What is the minimum number of orbital transfers required to move from the object YOU are orbiting to the object SAN is orbiting?"
func Part2(fileName string) int64 {
	orbits := loadOrbits(fileName)
	// For the graph, give each object an int ID
	idMap := mapObjectsToIDs(orbits)
	g := graph.New(len(idMap))
	var cost int64
	for k, v := range orbits {
		if k == you || k == santa {
			cost = 0
		} else {
			cost = 1
		}
		g.AddBothCost(idMap[k], idMap[v], cost)
	}
	_, distance := graph.ShortestPath(g, idMap[you], idMap[santa])
	return distance
}

func countOrbits(obj string, orbits map[string]string, memoCache map[string]int) int {
	if cnt, ok := memoCache[obj]; ok {
		//fmt.Printf("cache hit: %s %d\n", obj, cnt)
		return cnt
	}
	var rv int
	whatObjOrbits := orbits[obj]
	if whatObjOrbits == centerOfMass {
		rv = 1
	} else {
		rv = 1 + countOrbits(whatObjOrbits, orbits, memoCache)
	}
	memoCache[obj] = rv
	return rv
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

// Give each object a unique int ID for the graph
func mapObjectsToIDs(orbits map[string]string) map[string]int {
	m := map[string]int{}
	id := 0
	for k, v := range orbits {
		ensureObjectInIDMap(k, m, &id)
		ensureObjectInIDMap(v, m, &id)
	}
	return m
}

func ensureObjectInIDMap(obj string, m map[string]int, id *int) {
	if _, ok := m[obj]; !ok {
		m[obj] = *id
		*id++
	}
}
