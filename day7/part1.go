package day7

import (
	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

func PartOne() int {
	lines := helpers.ReadInputFile("day7/input.txt")
	beams := make(map[int]bool)
	for idx, char := range lines[0] {
		if char == 'S' {
			beams[idx] = true
		}
	}

	output := 0

	for _, line := range lines[1:] {
		nextBeams := make(map[int]bool)

		for idx, char := range line {
			if !beams[idx] {
				continue
			}

			if char == '.' {
				nextBeams[idx] = true
			} else if char == '^' {
				output += 1
				nextBeams[idx-1] = true
				nextBeams[idx+1] = true
			}
		}

		beams = nextBeams
	}

	return output
}
