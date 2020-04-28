package day08

import (
	"github.com/bfollek/aoc19go/util"
)

const (
	black = iota // 0
	white
	transparent
)

type layer [][]int // layer[width][height]
type img [][][]int // img[layers][width][height]

// type imgRow []int

// type layer []imgRow

// type img struct {
// 	width  int
// 	height int
// 	layers []layer
// }

// Part1 - find the layer that contains the fewest 0 digits. On that layer,
// what is the number of 1 digits multiplied by the number of 2 digits?
func Part1(width, height int, fileName string) int {
	img := loadImg(width, height, fileName)
	minSoFar := width*height + 1
	minIndex := -1
	for i, layer := range img {
		if nd := numPixelsInLayer(layer, black); nd < minSoFar {
			minSoFar = nd
			minIndex = i
		}
	}
	minLayer := img[minIndex]
	return numPixelsInLayer(minLayer, white) * numPixelsInLayer(minLayer, transparent)
}

// Part2 - What message is produced after decoding your image?
func Part2(width, height int, fileName string) int {
	// img := loadImg(width, height, fileName)
	// decoded := layer{}
	// for i := 0; i < width; i++ {
	// 	for j := 0; j < height; j++ {
	// 		layer = append(layer, findFirstNonTransparent(i, j, img))
	// 	}
	// }
	return 99
}

func loadImg(width, height int, fileName string) img {
	pixels := util.MustLoadIntSlice(fileName, "")
	numLayers := len(pixels) / (width * height)
	// Allocate the 3D slice
	img := make(img, numLayers)
	for i := range img {
		img[i] = make([][]int, width)
		for j := range img[i] {
			img[i][j] = make([]int, height)
		}
	}
	// Load the pixels into it
	p := 0
	for i := range img {
		for j := range img[i] {
			for k := range img[i][j] {
				img[i][j][k] = pixels[p]
				p++
			}
		}
	}
	return img
}

func numPixelsInLayer(layer [][]int, targetPixel int) int {
	total := 0
	for _, row := range layer {
		for _, pixel := range row {
			if pixel == targetPixel {
				total++
			}
		}
	}
	return total
}

func findFirstNonTransparent(widthPos, heightPos int, img img) int {
	return white
}
