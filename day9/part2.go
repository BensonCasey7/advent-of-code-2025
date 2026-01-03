package day9

import "github.com/bensoncasey7/advent-of-code-2025/helpers"

func PartTwo() int {
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
