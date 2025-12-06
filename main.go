package main

import (
	"fmt"
	"time"

	"github.com/bensoncasey7/advent-of-code-2025/day1"
	"github.com/bensoncasey7/advent-of-code-2025/day2"
	"github.com/bensoncasey7/advent-of-code-2025/day3"
	"github.com/bensoncasey7/advent-of-code-2025/day4"
	"github.com/bensoncasey7/advent-of-code-2025/day5"
	"github.com/bensoncasey7/advent-of-code-2025/day6"
	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

func main() {
	start := time.Now()

	fmt.Println("=== Day 1 ===")
	helpers.TimeSolution("Part one", day1.PartOne)
	helpers.TimeSolution("Part two", day1.PartTwo)
	fmt.Println()

	fmt.Println("=== Day 2 ===")
	helpers.TimeSolution("Part one", day2.PartOne)
	helpers.TimeSolution("Part two", day2.PartTwo)
	fmt.Println()

	fmt.Println("=== Day 3 ===")
	helpers.TimeSolution("Part one", day3.PartOne)
	helpers.TimeSolution("Part two", day3.PartTwo)
	fmt.Println()

	fmt.Println("=== Day 4 ===")
	helpers.TimeSolution("Part one", day4.PartOne)
	helpers.TimeSolution("Part two", day4.PartTwo)
	fmt.Println()

	fmt.Println("=== Day 5 ===")
	helpers.TimeSolution("Part one", day5.PartOne)
	helpers.TimeSolution("Part two", day5.PartTwo)
	fmt.Println()

	fmt.Println("=== Day 6 ===")
	helpers.TimeSolution("Part one", day6.PartOne)
	helpers.TimeSolution("Part two", day6.PartTwo)

	fmt.Printf("\n=== Total time: %v ===\n", time.Since(start))
}
