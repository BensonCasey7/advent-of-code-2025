package day8

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

type JunctionBox struct {
	x int
	y int
	z int
}

type Wire struct {
	box1   int
	box2   int
	length float64
}

const connectionsToMake = 1000
const circuitsToCount = 3

func PartOne() int {
	lines := helpers.ReadInputFile("day8/input.txt")

	boxes := []JunctionBox{}
	for _, line := range lines {
		boxes = append(boxes, createBoxFromCoordinateString(line))
	}

	wires := []Wire{}
	for i, box1 := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			distance := calculateDistance(box1, boxes[j])
			wires = append(wires, Wire{box1: i, box2: j, length: distance})
		}
	}
	sort.Slice(wires, func(i, j int) bool {
		return wires[i].length < wires[j].length
	})
	// One circuit is a map where each key is the index of a box
	circuits := []map[int]bool{}
	for _, wire := range wires[:connectionsToMake] {
		circuitFound := false
		for i, circuit := range circuits {
			if circuit[wire.box1] || circuit[wire.box2] {
				circuit[wire.box1] = true
				circuit[wire.box2] = true
				circuitFound = true

				// Now we loop through the rest to check if this circuit should be merged with another
				for j := i + 1; j < len(circuits); j++ {
					if circuits[j][wire.box1] || circuits[j][wire.box2] {
						// Merge this circuit, then delete it
						for key, _ := range circuits[j] {
							circuit[key] = true
						}
						circuits = append(circuits[:j], circuits[j+1:]...)
						break
					}
				}
				break
			}
		}

		if !circuitFound {
			circuits = append(circuits, map[int]bool{wire.box1: true, wire.box2: true})
		}
	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	output := 1
	for _, circuit := range circuits[:circuitsToCount] {
		output *= len(circuit)
	}

	return output
}

func createBoxFromCoordinateString(coordinateString string) JunctionBox {
	coordinates := strings.Split(coordinateString, ",")
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])
	z, _ := strconv.Atoi(coordinates[2])

	return JunctionBox{x: x, y: y, z: z}
}

func calculateDistance(b1, b2 JunctionBox) float64 {
	dx := float64(b2.x - b1.x)
	dy := float64(b2.y - b1.y)
	dz := float64(b2.z - b1.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
