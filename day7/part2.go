package day7

import (
	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

func PartTwo() int {
	lines := helpers.ReadInputFile("./day7/input.txt")

	pathCounts := make(map[int]int)

	for idx, char := range lines[0] {
		if char == 'S' {
			pathCounts[idx] = 1
		}
	}

	for rowNum := 1; rowNum < len(lines); rowNum++ {
		line := lines[rowNum]
		newPathCounts := make(map[int]int)

		for colNum, char := range line {
			parentCount, ok := pathCounts[colNum]
			if !ok {
				continue
			}

			if char == '.' {
				newPathCounts[colNum] += parentCount
			} else if char == '^' {
				newPathCounts[colNum-1] += parentCount
				newPathCounts[colNum+1] += parentCount
			}
		}

		pathCounts = newPathCounts
	}

	output := 0
	for _, count := range pathCounts {
		output += count
	}

	return output
}
