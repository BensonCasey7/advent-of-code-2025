package day9

import (
	"math"
	"strconv"
	"strings"

	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

func PartOne() int {
	lines := helpers.ReadInputFile("day9/input.txt")
	coordinates := [][]int{}

	for _, line := range lines {
		x, y := getCoordinatesFromString(line)
		coordinates = append(coordinates, []int{x, y})
	}
	output := 0

	for i, firstPair := range coordinates {
		for j := i + 1; j < len(coordinates); j++ {
			area := calculateArea(firstPair, coordinates[j])
			if area > output {
				output = area
			}
		}
	}

	return output
}

func getCoordinatesFromString(coordinateString string) (x, y int) {
	coordinates := strings.Split(coordinateString, ",")
	x, _ = strconv.Atoi(coordinates[0])
	y, _ = strconv.Atoi(coordinates[1])

	return x, y
}

func calculateArea(firstPair, secondPair []int) int {
	length := math.Abs(float64(secondPair[0]-firstPair[0])) + 1
	width := math.Abs(float64(secondPair[1]-firstPair[1])) + 1
	return int(length * width)
}
