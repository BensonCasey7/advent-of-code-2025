package helpers

import (
	"fmt"
	"time"
)

// In helpers/timer.go
func TimeSolution[T any](label string, solutionFunc func() T) {
	start := time.Now()
	solution := solutionFunc()
	fmt.Printf("%s: %v (Time taken: %v)\n", label, solution, time.Since(start))
}
