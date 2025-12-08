package day5

import (
	"strconv"
	"strings"

	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

func PartOne() int {
	lines := helpers.ReadInputFile("day5/input.txt")

	output := 0
	isReadingRanges := true
	ranges := [][]int{}
	ingredientIds := []int{}
	for _, line := range lines {

		if line == "" {
			isReadingRanges = false
		} else {
			if isReadingRanges {
				endpoints := strings.Split(line, "-")
				startPoint, _ := strconv.Atoi(endpoints[0])
				endPoint, _ := strconv.Atoi(endpoints[1])
				ranges = append(ranges, []int{startPoint, endPoint})
			} else {
				id, _ := strconv.Atoi(line)
				ingredientIds = append(ingredientIds, id)
			}
		}
	}

	for _, id := range ingredientIds {
		for _, endpoints := range ranges {
			if id >= endpoints[0] && id <= endpoints[1] {
				output++
				break
			}
		}
	}

	return output
}
