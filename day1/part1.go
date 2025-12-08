package day1

import (
	"strconv"

	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

func PartOne() int {
	lines := helpers.ReadInputFile("day1/input.txt")

	currentNumber := 50
	countOfZeros := 0

	for _, line := range lines {
		direction := string(line[0])
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if direction == "L" {
			currentNumber -= num
		} else {
			currentNumber += num
		}
		currentNumber = currentNumber % 100
		if currentNumber == 0 {
			countOfZeros++
		}
	}

	return countOfZeros
}
