package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"math"
)

func main() {
	lines := parse.Lines("cmd/day8/input.txt")
	fmt.Printf("Part 1: %d\n", verifyImage(lines[0], 25, 6))
	buildImage(lines[0], 25, 6)
}

func buildImage(image string, width, height int) {
	layers := parseLayers(image, width, height)

	type point struct{ row, column int }

	imageMap := map[point]int{}

	for _, layer := range layers {
		for row, vals := range layer {
			for col, val := range vals {
				currentPoint := point{row, col}
				if _, exists := imageMap[currentPoint]; !exists && val != 2 {
					imageMap[currentPoint] = val
				}
			}
		}
	}

	black := '▓'
	white := '░'
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if imageMap[point{row, col}] == 0 {
				fmt.Printf(string(black))
			} else {
				fmt.Printf(string(white))
			}
		}
		fmt.Println()
	}
}

func verifyImage(image string, width, height int) int {
	layers := parseLayers(image, width, height)

	var maxLayer [][]int
	minZeroes := math.MaxInt64
	for _, l := range layers {
		c := countItemsInLayer(0, l)
		if c < minZeroes {
			maxLayer = l
			minZeroes = c
		}
	}

	return countItemsInLayer(1, maxLayer) * countItemsInLayer(2, maxLayer)
}

func countItemsInLayer(item int, layer [][]int) int {
	count := 0
	for _, row := range layer {
		for _, val := range row {
			if item == val {
				count++
			}
		}
	}
	return count
}

func parseLayers(image string, width, height int) [][][]int {
	var layers [][][]int
	var currentLayer [][]int
	var currentRow []int

	for _, r := range image {
		v := int(r - '0')
		currentRow = append(currentRow, v)
		if len(currentRow) > (width - 1) {
			currentLayer = append(currentLayer, currentRow)
			currentRow = []int{}
		}
		if len(currentLayer) > (height - 1) {
			layers = append(layers, currentLayer)
			currentLayer = [][]int{}
		}
	}
	return layers
}
