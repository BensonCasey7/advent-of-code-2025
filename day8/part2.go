package day8

import (
	"sort"

	"github.com/bensoncasey7/advent-of-code-2025/helpers"
)

func PartTwo() int {
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
	for _, wire := range wires {
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

		// Check if we've condensed down to one circuit with the size of the initial input
		if len(circuits[0]) == len(lines) {
			return boxes[wire.box1].x * boxes[wire.box2].x
		}
	}

	// We should never get here
	return 0
}
