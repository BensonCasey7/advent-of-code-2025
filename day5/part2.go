package day5

import (
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

func PartTwo() int {
	lines := helpers.ReadInputFile("day5/input.txt")

	output := 0
	ranges := [][]int{}
	for _, line := range lines {

		if line == "" {
			break
		} else {
			endpoints := strings.Split(line, "-")
			startPoint, _ := strconv.Atoi(endpoints[0])
			endPoint, _ := strconv.Atoi(endpoints[1])
			ranges = append(ranges, []int{startPoint, endPoint})
		}
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	combinedRanges := combineRanges(ranges)
	for _, endpoints := range combinedRanges {
		output += (1 + endpoints[1] - endpoints[0])
	}
	return output
}

func combineRanges(ranges [][]int) [][]int {
	output := ranges

	keepCombining := true
	for keepCombining {
		keepCombining = false

		for idx, endpoints := range output {
			for i := idx + 1; i < len(output); i++ {
				candidate := output[i]
				if (endpoints[0] <= candidate[0] && endpoints[1] >= candidate[0]) || (endpoints[0] <= candidate[1] && endpoints[1] >= candidate[1]) {
					startPoint := slices.Min([]int{endpoints[0], candidate[0]})
					endPoint := slices.Max([]int{endpoints[1], candidate[1]})

					newOut := [][]int{}
					for j, rangeToKeep := range output {
						if j != idx && j != i {
							newOut = append(newOut, rangeToKeep)
						}
					}
					newOut = append(newOut, []int{startPoint, endPoint})
					sort.Slice(newOut, func(i, j int) bool {
						return newOut[i][0] < newOut[j][0]
					})
					output = newOut

					keepCombining = true
					break
				}
			}

			if keepCombining {
				break
			}
		}
	}

	return output
}
