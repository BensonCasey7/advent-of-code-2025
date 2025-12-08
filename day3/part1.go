package day3

import (
	"strconv"

	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

func PartOne() int {
	lines := helpers.ReadInputFile("day3/input.txt")

	output := 0
	for _, line := range lines {

		digits := []int{}

		for _, char := range line {
			num := int(char - '0')
			digits = append(digits, num)
		}

		largestDigit := -1
		largestDigitIdx := 0
		for i, num := range digits {
			if num > largestDigit {
				largestDigit = num
				largestDigitIdx = i
			}
		}

		largestToTheLeft := -1
		for j := 0; j < largestDigitIdx; j++ {
			if digits[j] > largestToTheLeft {
				largestToTheLeft = digits[j]
			}
		}

		largestToTheRight := -1
		for j := (largestDigitIdx + 1); j < len(digits); j++ {
			if digits[j] > largestToTheRight {
				largestToTheRight = digits[j]
			}
		}

		if largestToTheRight != -1 {
			output += combineDigits(largestDigit, largestToTheRight)
		} else {
			output += combineDigits(largestToTheLeft, largestDigit)
		}
	}

	return output
}

func combineDigits(int1 int, int2 int) int {
	combined := strconv.Itoa(int1) + strconv.Itoa(int2)
	result, err := strconv.Atoi(combined)
	if err != nil {
		return 0
	}

	return result
}
