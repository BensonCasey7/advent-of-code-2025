package day1

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func PartTwo() int {
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

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return countOfZeros
}
