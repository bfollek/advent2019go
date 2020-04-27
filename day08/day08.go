package day08

import (
	"github.com/bfollek/aoc19go/util"
)

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
func Part1(width, height int, fileName string) int {
	img := loadImg(width, height, fileName)
	minSoFar := width*height + 1
	minIndex := -1
	for i, layer := range img.layers {
		if nd := numDigitsInLayer(layer, 0); nd < minSoFar {
			minSoFar = nd
			minIndex = i
		}
	}
	minLayer := img.layers[minIndex]
	return numDigitsInLayer(minLayer, 1) * numDigitsInLayer(minLayer, 2)
}

func loadImg(width, height int, fileName string) img {
	img := img{width, height, []layer{}}
	pixels := util.MustLoadIntSlice(fileName, "")
	rows := []imgRow{}
	for i := 0; i < len(pixels); i += width {
		row := pixels[i : i+width]
		rows = append(rows, row)
	}
	for i := 0; i < len(rows); i += height {
		layer := rows[i : i+height]
		img.layers = append(img.layers, layer)
	}
	return img
}

func numDigitsInLayer(layer layer, target int) int {
	total := 0
	for _, row := range layer {
		for _, i := range row {
			if i == target {
				total++
			}
		}
	}
	return total
}
