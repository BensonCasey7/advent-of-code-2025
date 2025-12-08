package day3

import (
	"strconv"
	"strings"

	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

func PartTwo() int {
	lines := helpers.ReadInputFile("day3/input.txt")

	output := 0
	for _, line := range lines {

		digits := []int{}

		for _, char := range line {
			num := int(char - '0')
			digits = append(digits, num)
		}

		output += maxDigits(digits, 12)
	}

	return output
}

func maxDigits(digits []int, desiredLength int) int {
	maxSlice := digits[:desiredLength]
	maxValue := squashIntSlice(maxSlice)

	for i := desiredLength; i < len(digits); i++ {
		candidateDigit := digits[i]
		thisWindowSlice := maxSlice
		thisWindowSlice = append(thisWindowSlice, candidateDigit)

		for j := 0; j < len(thisWindowSlice); j++ {
			candidateSlice := make([]int, 0, len(thisWindowSlice)-1)
			candidateSlice = append(candidateSlice, thisWindowSlice[:j]...)
			candidateSlice = append(candidateSlice, thisWindowSlice[j+1:]...)
			candidateValue := squashIntSlice(candidateSlice)

			if candidateValue > maxValue {
				maxSlice = candidateSlice
				maxValue = candidateValue
			}
		}
	}

	return maxValue
}

func squashIntSlice(numbers []int) int {
	var builder strings.Builder
	for _, num := range numbers {
		builder.WriteString(strconv.Itoa(num))
	}
	result, err := strconv.Atoi(builder.String())
	if err != nil {
		return 0
	}

	return result
}
