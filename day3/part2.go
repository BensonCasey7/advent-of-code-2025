package day3

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func PartTwo() int {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	output := 0
	for scanner.Scan() {
		line := scanner.Text()

		digits := []int{}

		for _, char := range line {
			num := int(char - '0')
			digits = append(digits, num)
		}

		output += maxDigits(digits, 12)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
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
