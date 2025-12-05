package day4

import (
	"bufio"
	"os"
)

func PartOne() int {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	output := 0
	currentLine := 0
	locations := [][]int{}
	columnCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		columnCount = len(line)

		for idx, char := range line {
			if char == '@' {
				locations = append(locations, []int{currentLine, idx})
			}
		}

		currentLine++
	}

	grid := make([][]bool, currentLine)
	for i := range grid {
		grid[i] = make([]bool, columnCount)
	}

	for _, location := range locations {
		grid[location[0]][location[1]] = true
	}

	for _, location := range locations {
		hitCount := 0
		for _, toCheck := range locationsToCheck(location[0], location[1], currentLine, columnCount) {
			if grid[toCheck[0]][toCheck[1]] {
				hitCount++
			}
		}

		if hitCount < 4 {
			output++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return output
}

func locationsToCheck(x int, y int, rowCount int, columnCount int) [][]int {
	output := [][]int{}

	candidateLocations := [][]int{}
	candidateLocations = append(candidateLocations, []int{x - 1, y - 1})
	candidateLocations = append(candidateLocations, []int{x, y - 1})
	candidateLocations = append(candidateLocations, []int{x + 1, y - 1})
	candidateLocations = append(candidateLocations, []int{x - 1, y})
	candidateLocations = append(candidateLocations, []int{x + 1, y})
	candidateLocations = append(candidateLocations, []int{x - 1, y + 1})
	candidateLocations = append(candidateLocations, []int{x, y + 1})
	candidateLocations = append(candidateLocations, []int{x + 1, y + 1})

	for _, location := range candidateLocations {
		if locationIsValid(location[0], location[1], rowCount, columnCount) {
			output = append(output, location)
		}
	}

	return output
}

func locationIsValid(x int, y int, rowCount int, columnCount int) bool {
	return (x >= 0) && (x < columnCount) && (y >= 0) && (y < rowCount)
}
