package day4

import (
	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

func PartTwo() int {
	lines := helpers.ReadInputFile("day4/input.txt")

	output := 0
	locations := [][]int{}
	columnCount := 0
	for currentLine, line := range lines {
		columnCount = len(line)

		for idx, char := range line {
			if char == '@' {
				locations = append(locations, []int{currentLine, idx})
			}
		}
	}
	currentLine := len(lines)

	grid := make([][]bool, currentLine)
	for i := range grid {
		grid[i] = make([]bool, columnCount)
	}

	for _, location := range locations {
		grid[location[0]][location[1]] = true
	}

	keepIterating := true
	for keepIterating {
		locationsToRemove := [][]int{}
		nextIterationLocations := [][]int{}

		for _, location := range locations {
			hitCount := 0
			for _, toCheck := range locationsToCheck(location[0], location[1], currentLine, columnCount) {
				if grid[toCheck[0]][toCheck[1]] {
					hitCount++
				}
			}

			if hitCount < 4 {
				locationsToRemove = append(locationsToRemove, location)
			} else {
				nextIterationLocations = append(nextIterationLocations, location)
			}
		}

		if len(locationsToRemove) > 0 {
			output += len(locationsToRemove)
			for _, toRemove := range locationsToRemove {
				grid[toRemove[0]][toRemove[1]] = false
			}
		} else {
			keepIterating = false
		}

		locations = nextIterationLocations
	}

	return output
}
