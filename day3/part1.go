package day3

import (
	"bufio"
	"os"
	"strconv"
)

func PartOne() int {
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

	if err := scanner.Err(); err != nil {
		panic(err)
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
