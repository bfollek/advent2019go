package day08

import (
	"fmt"
	"log"

	"github.com/bfollek/aoc19go/util"
)

const (
	black = iota // 0
	white
	transparent
)

// Row is a row of the image, an array of pixels.
type Row []int

// Layer is a layer of the image, an array of `Row`.
type Layer []Row

// Img is the full image - an array of `Layer`.
type Img []Layer

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
func Part2(height, width int, fileName string) Img {
	img := loadImg(height, width, fileName)
	decoded := newImg(1, height, width)
	onlyLayer := decoded[0]
	for h := range onlyLayer {
		for w := range onlyLayer[h] {
			pixel, err := findFirstNonTransparent(h, w, img)
			if err != nil {
				log.Fatal(err)
			}
			onlyLayer[h][w] = pixel
		}
	}
	decoded.readableDisplay(false)
	return decoded
}

func numPixelsInLayer(layer Layer, targetPixel int) int {
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

func findFirstNonTransparent(heightPos, widthPos int, img Img) (int, error) {
	// Check all layers for a non-transparent pixel.
	for _, layer := range img {
		row := layer[heightPos]
		pixel := row[widthPos]
		if pixel != transparent {
			return pixel, nil
		}
	}
	return transparent, fmt.Errorf("Could not find non-transparent pixel at heightPos %d, widthPost %d", heightPos, widthPos)
}

func newImg(numLayers, height, width int) Img {
	// Allocate the 3D slice
	img := make(Img, numLayers)
	for i := range img {
		img[i] = make(Layer, height)
		for j := range img[i] {
			img[i][j] = make(Row, width)
		}
	}
	return img
}

func loadImg(height, width int, fileName string) Img {
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

func (img Img) readableDisplay(labelLayer bool) {
	for i := range img {
		if labelLayer {
			fmt.Printf("\nLayer %d", i)
		}
		for j := range img[i] {
			fmt.Print("\n")
			for k := range img[i][j] {
				pixel := img[i][j][k]
				s := "*"
				if pixel == black {
					s = " "
				}
				fmt.Print(s)
			}
		}
	}
	fmt.Println()
}
