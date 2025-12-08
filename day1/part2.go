package day1

import (
	"math"
	"strconv"

	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

func PartTwo() int {
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
			if currentNumber > 0 && (currentNumber-num < 0) {
				countOfZeros++
			}
			currentNumber -= num
		} else {
			if currentNumber < 0 && (currentNumber+num > 0) {
				countOfZeros++
			}
			currentNumber += num
		}

		if currentNumber == 0 {
			countOfZeros++
		} else {
			clicks := int(math.Abs(float64(currentNumber / 100)))
			countOfZeros += clicks
			currentNumber = currentNumber % 100
		}
	}

	return countOfZeros
}
