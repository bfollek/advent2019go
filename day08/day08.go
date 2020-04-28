package day08

import (
	"fmt"

	"github.com/bfollek/aoc19go/util"
)

const (
	black = iota // 0
	white
	transparent
)

type row []int
type layer []row
type img []layer

// Part1 - find the layer that contains the fewest 0 digits. On that layer,
// what is the number of 1 digits multiplied by the number of 2 digits?
func Part1(height, width int, fileName string) int {
	img := loadImg(height, width, fileName)
	//img.display(true)
	minSoFar := height*width + 1
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
func Part2(height, width int, fileName string) int {
	// img := loadImg(width, height, fileName)
	// decoded := layer{}
	// for i := 0; i < width; i++ {
	// 	for j := 0; j < height; j++ {
	// 		layer = append(layer, findFirstNonTransparent(i, j, img))
	// 	}
	// }
	return 99
}

func numPixelsInLayer(layer layer, targetPixel int) int {
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

func findFirstNonTransparent(heightPos, widthPos int, img img) int {
	return white
}

func newImg(numLayers, height, width int) img {
	// Allocate the 3D slice
	img := make(img, numLayers)
	for i := range img {
		img[i] = make(layer, height)
		for j := range img[i] {
			img[i][j] = make(row, width)
		}
	}
	return img
}

func loadImg(height, width int, fileName string) img {
	pixels := util.MustLoadIntSlice(fileName, "")
	numLayers := len(pixels) / (height * width)
	img := newImg(numLayers, height, width)
	// Load the pixels into the 3D img slice
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

func (img img) display(labelLayer bool) {
	for i := range img {
		if labelLayer {
			fmt.Printf("\nLayer %d", i)
		}
		for j := range img[i] {
			fmt.Print("\n\t")
			for k := range img[i][j] {
				fmt.Print(img[i][j][k])
			}
		}
	}
}
