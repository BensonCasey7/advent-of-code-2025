package helpers

import (
	"bufio"
	"os"
)

func ReadInputFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	output := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return output
}
