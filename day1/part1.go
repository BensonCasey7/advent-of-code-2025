package day1

import (
	"bufio"
	"os"
	"strconv"
)

func PartOne() int {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	currentNumber := 50
	countOfZeros := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

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

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return countOfZeros
}
