package day6

import (
	"strconv"
	"strings"

	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

func PartTwo() int {
	lines := helpers.ReadInputFile("day6/input.txt")

	output := 0
	operands, operations := splitInputByColumns(lines)

	for i, operation := range operations {
		output += computeProblem(operands[i], operation)
	}

	return output
}

func splitInputByColumns(lines []string) ([][]int, []string) {
	outputNumbers := [][]int{}
	outputOperations := []string{}

	operands := lines[:(len(lines) - 1)]
	operationString := lines[(len(lines) - 1)]

	operationIndices := []int{}
	for idx, char := range operationString {
		if char != ' ' {
			outputOperations = append(outputOperations, string(char))
			operationIndices = append(operationIndices, idx)
		}
	}

	for i, opIdx := range operationIndices {
		numbersForThisProblem := []int{}

		var nextOpIdx int
		if i == (len(operationIndices) - 1) {
			nextOpIdx = len(operands[0])
		} else {
			nextOpIdx = operationIndices[i+1] - 1
		}

		operandStrings := []string{}
		for _, operandRow := range operands {
			operandStrings = append(operandStrings, operandRow[opIdx:(nextOpIdx)])
		}

		for j := len(operandStrings[0]) - 1; j >= 0; j-- {
			var builder strings.Builder
			for _, operandString := range operandStrings {
				char := rune(operandString[j])
				if char != ' ' {
					builder.WriteString(string(char))
				}
			}
			num, _ := strconv.Atoi(builder.String())
			numbersForThisProblem = append(numbersForThisProblem, num)
		}

		outputNumbers = append(outputNumbers, numbersForThisProblem)
	}

	return outputNumbers, outputOperations
}
