package day6

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func PartOne() int {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	output := 0
	lines := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, splitLines(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	operands := lines[:(len(lines) - 1)]
	operations := lines[(len(lines) - 1)]

	for i := 0; i < len(operations); i++ {
		nums := []int{}
		for _, numSlice := range operands {
			num, _ := strconv.Atoi(numSlice[i])
			nums = append(nums, num)
		}

		output += computeProblem(nums, operations[i])
	}

	return output
}

func splitLines(line string) []string {
	output := []string{}

	var builder strings.Builder
	for _, char := range line {
		if char == ' ' {
			if builder.Len() > 0 {
				output = append(output, builder.String())
			}
			builder.Reset()
		} else {
			builder.WriteString(string(char))
		}
	}
	if builder.Len() > 0 {
		output = append(output, builder.String())
	}

	return output
}

func computeProblem(operands []int, operation string) int {
	output := 0

	if operation == "*" {
		output = 1
		for _, operand := range operands {
			output *= operand
		}
	} else {
		for _, operand := range operands {
			output += operand
		}
	}

	return output
}
