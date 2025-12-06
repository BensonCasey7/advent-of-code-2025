package day5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func PartOne() int {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	output := 0
	isReadingRanges := true
	ranges := [][]int{}
	ingredientIds := []int{}
	for scanner.Scan() {
		line := scanner.Text()

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
