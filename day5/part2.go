package day5

import (
	"bufio"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func PartTwo() int {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	output := 0
	ranges := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()

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
