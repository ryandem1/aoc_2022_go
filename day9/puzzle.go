package day9

import (
	"github.com/ryandem1/aoc_2022_go/common"
)

func Part1() (solution common.Solution) {
	solution.Prompt = `
Simulate your complete hypothetical series of motions. 
How many positions does the tail of the rope visit at 
least once?
`
	positionsVisited := 0
	var uniqueTailPosVisited []common.Coords2D

	rope := newBridgeRope()
	for motion := range readMotions() {
		rope.applyMotion(motion)

		if !common.Contains(uniqueTailPosVisited, rope.tailPos) {
			uniqueTailPosVisited = append(uniqueTailPosVisited, rope.tailPos)
		}
	}

	positionsVisited = len(uniqueTailPosVisited)
	solution.Answer = positionsVisited
	return solution
}

func Part2() (solution common.Solution) {
	solution.Prompt = ``
	return solution
}
