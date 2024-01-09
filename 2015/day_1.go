package main

import (
	"fmt"

	"github.com/tisp/aoc-2015/utils"
)

const (
	FLOOR_UP   = "("
	FLOOR_DOWN = ")"
	BASEMENT   = -1
)

func main() {
	SolutionPartOne()
	SolutionPartTwo()
}

func SolutionPartOne() {
	input := utils.ReadSampleByDay(1, 1)

	current_floor := 0
	for _, dir := range input {
		switch string(dir) {
		case FLOOR_UP:
			current_floor += 1
		case FLOOR_DOWN:
			current_floor -= 1
		}
	}
	fmt.Println(current_floor)
}

func SolutionPartTwo() {
	input := utils.ReadSampleByDay(2, 1)
	current_floor := 0
	for pos, dir := range input {
		switch string(dir) {
		case FLOOR_UP:
			current_floor += 1
		case FLOOR_DOWN:
			current_floor -= 1
		}
		if current_floor == BASEMENT {
			fmt.Println(pos + 1)
		}
	}
}
