package day01

import "github.com/bfollek/advent2019go/util"

// Part1 calculates the fuel requirements for all of the modules on my spacecraft.
func Part1(fileName string) int {
	sum := 0
	masses := loadMasses(fileName)
	for _, mass := range masses {
		sum += fuelRequired(mass)
	}
	return sum
}

// Part2 calculates the fuel requirements for all of the modules on my spacecraft,
// plus the fuel costs of the fuel.
func Part2(fileName string) int {
	sum := 0
	masses := loadMasses(fileName)
	for _, mass := range masses {
		sum += fuelRequiredMeta(mass, 0)
	}
	return sum
}

// "...to find the fuel required for a module, take its mass,
// divide by three, round down, and subtract 2."
func fuelRequired(mass int) int {
	return mass/3 - 2
}

// "So, for each module mass, calculate its fuel and add it to the total. Then,
// treat the fuel amount you just calculated as the input mass and repeat the
// process, continuing until a fuel requirement is zero or negative."
func fuelRequiredMeta(mass, total int) int {
	f := fuelRequired(mass)
	if f < 1 {
		return total
	}
	return fuelRequiredMeta(f, total+f)
}

func loadMasses(fileName string) []int {
	ss := util.MustLoadStringSlice(fileName, "\n")
	masses := []int{}
	for _, s := range ss {
		i := util.MustAtoi(s)
		masses = append(masses, i)
	}
	return masses
}
